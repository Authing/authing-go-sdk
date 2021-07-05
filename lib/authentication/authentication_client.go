package authentication

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"github.com/Authing/authing-go-sdk/lib/util"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	HttpClient              *http.Client
	AppId                   string
	Protocol                constant.ProtocolEnum
	Secret                  string
	Host                    string
	RedirectUri             string
	TokenEndPointAuthMethod constant.AuthMethodEnum

	Log func(s string)
}

func NewClient(appId string, secret string, host ...string) *Client {
	var clientHost string
	if len(host) == 0 {
		clientHost = constant.CoreAuthingDefaultUrl
	} else {
		clientHost = host[0]
	}
	c := &Client{
		HttpClient:  nil,
		AppId:       appId,
		Protocol:    "",
		Secret:      secret,
		Host:        clientHost,
		RedirectUri: "",
		Log:         nil,
	}
	if c.HttpClient == nil {
		c.HttpClient = &http.Client{}
	}
	return c
}

type OidcParams struct {
	appId               string
	redirectUri         string
	responseType        string
	responseMode        string
	state               string
	nonce               string
	scope               string
	codeChallengeMethod string
	codeChallenge       string
}

func (c *Client) BuildAuthorizeUrlByOidc(params OidcParams) (string, error) {
	if c.AppId == "" {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 appId")
	}
	if c.Protocol == constant.OIDC {
		return constant.StringEmpty, errors.New("初始化 AuthenticationClient 传入的 protocol 应为 ProtocolEnum.OIDC")
	}
	if params.redirectUri == "" {
		return constant.StringEmpty, errors.New("redirectUri 不能为空")
	}
	var scope = ""
	if strings.Contains(params.scope, "offline_access") {
		scope = "consent"
	}
	dataMap := map[string]string{
		"client_id":     util.GetValidValue(params.appId, c.AppId),
		"scope":         util.GetValidValue(params.scope, "openid profile email phone address"),
		"state":         util.GetValidValue(params.state, util.RandomString(12)),
		"nonce":         util.GetValidValue(params.nonce, util.RandomString(12)),
		"response_mode": util.GetValidValue(params.responseMode, constant.StringEmpty),
		"response_type": util.GetValidValue(params.responseType, "code"),
		"redirect_uri":  util.GetValidValue(params.redirectUri, c.RedirectUri),
		"prompt":        util.GetValidValue(scope),
	}
	return c.Host + "/oidc/auth?" + util.GetQueryString(dataMap), nil
}

func (c *Client) GetAccessTokenByCode(code string) (string, error) {
	if c.AppId == "" {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 appId")
	}
	if c.Secret == "" && c.TokenEndPointAuthMethod != constant.None {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 Secret")
	}
	url := c.Host + "/oidc/token"

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"client_id": c.AppId,
		"client_secret": c.Secret,
		"grant_type": "authorization_code",
		"code": code,
		"redirect_uri": c.RedirectUri,
	}

	switch c.TokenEndPointAuthMethod {
	case constant.ClientSecretPost:
		body["client_id"] = c.AppId
		body["client_secret"] = c.Secret
	case constant.ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",c.AppId, c.Secret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = c.AppId
	}
	resp, err := c.SendHttpRequest(url,constant.HttpMethodPost,header,body)
	return string(resp), err
}

func (c *Client) GetUserInfoByAccessToken(accessToken string) (string, error) {
	if accessToken == constant.StringEmpty {
		return constant.StringEmpty, errors.New("accessToken 不能为空")
	}
	url := c.Host + "/oidc/me?access_token=" + accessToken
	resp, err := c.SendHttpRequest(url, constant.HttpMethodGet, nil, nil)
	return string(resp), err
}

func (c *Client) GetNewAccessTokenByRefreshToken (refreshToken string) (string, error) {
	if c.Protocol != constant.OIDC || c.Protocol != constant.OAUTH {
		return constant.StringEmpty, errors.New("初始化 AuthenticationClient 时传入的 protocol 参数必须为 ProtocolEnum.OAUTH 或 ProtocolEnum.OIDC，请检查参数")
	}
	if c.Secret == "" && c.TokenEndPointAuthMethod != constant.None {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 Secret")
	}

	url := c.Host + fmt.Sprintf("/%s/token", c.Protocol)

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"client_id": c.AppId,
		"client_secret": c.Secret,
		"grant_type": "refresh_token",
		"refresh_token": refreshToken,
	}

	switch c.TokenEndPointAuthMethod {
	case constant.ClientSecretPost:
		body["client_id"] = c.AppId
		body["client_secret"] = c.Secret
	case constant.ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",c.AppId, c.Secret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = c.AppId
	}
	resp, err := c.SendHttpRequest(url,constant.HttpMethodPost,header,body)
	return string(resp), err
}

func (c *Client) IntrospectToken (req model.ValidateTokenRequest) (string, error) {
	if req.IdToken == constant.StringEmpty && req.AccessToken == constant.StringEmpty {
		return constant.StringEmpty, errors.New("请传入 AccessToken 或 IdToken")
	}
	if req.IdToken != constant.StringEmpty && req.AccessToken != constant.StringEmpty {
		return constant.StringEmpty, errors.New("AccessToken 和 IdToken 不能同时传入")
	}

	url := c.Host + "/api/v2/oidc/validate_token?"
	if req.IdToken != constant.StringEmpty {
		url += "id_token=" + req.IdToken
	} else if req.AccessToken != constant.StringEmpty {
		url += "access_token=" + req.AccessToken
	}

	resp, err := c.SendHttpRequest(url,constant.HttpMethodGet,nil,nil)
	return string(resp), err
}

func (c *Client) GetAccessTokenByClientCredentials (req model.GetAccessTokenByClientCredentialsRequest) (string, error) {
	if req.Scope == constant.StringEmpty {
		return constant.StringEmpty, errors.New("请传入 scope 参数，请看文档：https://docs.authing.cn/v2/guides/authorization/m2m-authz.html")
	}
	if req.ClientCredentialInput == nil {
		return constant.StringEmpty, errors.New("请在调用本方法时传入 ClientCredentialInput 参数，请看文档：https://docs.authing.cn/v2/guides/authorization/m2m-authz.html")
	}

	url := c.Host + "/oidc/token"

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"client_id": req.ClientCredentialInput.AccessKey,
		"client_secret": req.ClientCredentialInput.SecretKey,
		"grant_type": "client_credentials",
		"scope": req.Scope,
	}

	resp, err := c.SendHttpRequest(url,constant.HttpMethodPost,header,body)
	return string(resp), err
}

func (c *Client) SendHttpRequest(url string, method string, header map[string]string, body map[string]string) ([]byte, error) {
	var form http.Request
	form.ParseForm()
	if body != nil && len(body) != 0 {
		for key, value := range body {
			form.Form.Add(key, value)
		}
	}
	reqBody := strings.TrimSpace(form.Form.Encode())
	req, err := http.NewRequest(method, url, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//增加header选项
	if header != nil && len(header) != 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}
	res, err := c.HttpClient.Do(req)
	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	return respBody, nil
}
