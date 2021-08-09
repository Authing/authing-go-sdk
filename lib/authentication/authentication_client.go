package authentication

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"github.com/Authing/authing-go-sdk/lib/util"
	"github.com/Authing/authing-go-sdk/lib/util/cacheutil"
	simplejson "github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Client struct {
	HttpClient              *http.Client
	AppId                   string
	Protocol                constant.ProtocolEnum
	Secret                  string
	Host                    string
	RedirectUri             string
	userPoolId              string
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

func (c *Client) BuildAuthorizeUrlByOidc(params model.OidcParams) (string, error) {
	if c.AppId == "" {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 appId")
	}
	if c.Protocol != constant.OIDC {
		return constant.StringEmpty, errors.New("初始化 AuthenticationClient 传入的 protocol 应为 ProtocolEnum.OIDC")
	}
	if params.RedirectUri == "" {
		return constant.StringEmpty, errors.New("redirectUri 不能为空")
	}
	var scope = ""
	if strings.Contains(params.Scope, "offline_access") {
		scope = "consent"
	}
	dataMap := map[string]string{
		"client_id":     util.GetValidValue(params.AppId, c.AppId),
		"scope":         util.GetValidValue(params.Scope, "openid profile email phone address"),
		"state":         util.GetValidValue(params.State, util.RandomString(12)),
		"nonce":         util.GetValidValue(params.Nonce, util.RandomString(12)),
		"response_mode": util.GetValidValue(params.ResponseMode, constant.StringEmpty),
		"response_type": util.GetValidValue(params.ResponseType, "code"),
		"redirect_uri":  util.GetValidValue(params.RedirectUri, c.RedirectUri),
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
		"client_id":     c.AppId,
		"client_secret": c.Secret,
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  c.RedirectUri,
	}

	switch c.TokenEndPointAuthMethod {
	case constant.ClientSecretPost:
		body["client_id"] = c.AppId
		body["client_secret"] = c.Secret
	case constant.ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.AppId, c.Secret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = c.AppId
	}
	resp, err := c.SendHttpRequest(url, constant.HttpMethodPost, header, body)
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

func (c *Client) GetNewAccessTokenByRefreshToken(refreshToken string) (string, error) {
	if c.Protocol != constant.OIDC && c.Protocol != constant.OAUTH {
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
		"client_id":     c.AppId,
		"client_secret": c.Secret,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}

	switch c.TokenEndPointAuthMethod {
	case constant.ClientSecretPost:
		body["client_id"] = c.AppId
		body["client_secret"] = c.Secret
	case constant.ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.AppId, c.Secret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = c.AppId
	}
	resp, err := c.SendHttpRequest(url, constant.HttpMethodPost, header, body)
	return string(resp), err
}

func (c *Client) IntrospectToken(token string) (string, error) {
	url := c.Host + fmt.Sprintf("/%s/token/introspection", c.Protocol)

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"token": token,
	}

	switch c.TokenEndPointAuthMethod {
	case constant.ClientSecretPost:
		body["client_id"] = c.AppId
		body["client_secret"] = c.Secret
	case constant.ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.AppId, c.Secret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = c.AppId
	}
	resp, err := c.SendHttpRequest(url, constant.HttpMethodPost, header, body)
	return string(resp), err
}

func (c *Client) ValidateToken(req model.ValidateTokenRequest) (string, error) {
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

	resp, err := c.SendHttpRequest(url, constant.HttpMethodGet, nil, nil)
	return string(resp), err
}

func (c *Client) RevokeToken(token string) (string, error) {
	if c.Protocol != constant.OIDC && c.Protocol != constant.OAUTH {
		return constant.StringEmpty, errors.New("初始化 AuthenticationClient 时传入的 protocol 参数必须为 ProtocolEnum.OAUTH 或 ProtocolEnum.OIDC，请检查参数")
	}
	if c.Secret == "" && c.TokenEndPointAuthMethod != constant.None {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 Secret")
	}

	url := c.Host + fmt.Sprintf("/%s/token/revocation", c.Protocol)

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	body := map[string]string{
		"client_id": c.AppId,
		"token":     token,
	}

	switch c.TokenEndPointAuthMethod {
	case constant.ClientSecretPost:
		body["client_id"] = c.AppId
		body["client_secret"] = c.Secret
	case constant.ClientSecretBasic:
		base64String := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.AppId, c.Secret)))
		header["Authorization"] = base64String
	default:
		body["client_id"] = c.AppId
	}
	resp, err := c.SendHttpRequest(url, constant.HttpMethodPost, header, body)
	return string(resp), err
}

func (c *Client) GetAccessTokenByClientCredentials(req model.GetAccessTokenByClientCredentialsRequest) (string, error) {
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
		"client_id":     req.ClientCredentialInput.AccessKey,
		"client_secret": req.ClientCredentialInput.SecretKey,
		"grant_type":    "client_credentials",
		"scope":         req.Scope,
	}

	resp, err := c.SendHttpRequest(url, constant.HttpMethodPost, header, body)
	return string(resp), err
}

func (c *Client) LoginByUserName(request model.LoginByUsernameInput) (*model.User, error) {
	request.Password = util.RsaEncrypt(request.Password)
	reqParam := make(map[string]interface{})
	reqParam["input"] = request
	data, _ := json.Marshal(&reqParam)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.LoginByUsernameDocument, variables)
	if err != nil {
		return nil, err
	}
	return loginGetUserInfo(b, "loginByUsername")
}

func (c *Client) LoginByEmail(request model.LoginByEmailInput) (*model.User, error) {
	request.Password = util.RsaEncrypt(request.Password)
	reqParam := make(map[string]interface{})
	reqParam["input"] = request
	data, _ := json.Marshal(&reqParam)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.LoginByEmailDocument, variables)
	if err != nil {
		return nil, err
	}
	return loginGetUserInfo(b, "loginByEmail")
}

func (c *Client) LoginByPhonePassword(request model.LoginByPhonePasswordInput) (*model.User, error) {
	request.Password = util.RsaEncrypt(request.Password)
	reqParam := make(map[string]interface{})
	reqParam["input"] = request
	data, _ := json.Marshal(&reqParam)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.LoginByPhonePasswordDocument, variables)
	if err != nil {
		return nil, err
	}
	return loginGetUserInfo(b, "loginByPhonePassword")
}

/*func (c *Client) LoginByPhoneCode(request model.LoginByPhoneCodeInput) (*model.User,error) {
	reqParam := make(map[string]interface{})
	reqParam["input"] = request
	data, _ := json.Marshal(&reqParam)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.LoginByPhoneCodeDocument, variables)
	if err != nil {
		return nil, err
	}
	return loginGetUserInfo(b,"loginByPhoneCode")
}

func (c *Client) SendSmsCode(phone string) (*model.CommonMessage, error) {
	var result *model.CommonMessage
	variables := map[string]interface{}{
		"phone": phone,
	}
	b, err := c.SendHttpRequestManage(c.Host+"/api/v2/sms/send", constant.HttpMethodPost, constant.StringEmpty, variables)
	if err != nil {
		return result, err
	}
	log.Println(string(b))
	jsoniter.Unmarshal(b, result)
	return result, nil
}*/

func loginGetUserInfo(b []byte, userKey string) (*model.User, error) {
	var result *simplejson.Json
	result, err := simplejson.NewJson(b)
	if _, r := result.CheckGet("errors"); r {
		msg, err := result.Get("errors").GetIndex(0).Get("message").Get("message").String()
		if err != nil {
			return nil, err
		}
		return nil, errors.New(msg)
	}
	byteUser, err := result.Get("data").Get(userKey).MarshalJSON()
	if err != nil {
		return nil, err
	}
	resultUser := model.User{}
	err = json.Unmarshal(byteUser, &resultUser)
	if err != nil {
		return nil, err
	}
	return &resultUser, nil
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

func (c *Client) SendHttpRequestManage(url string, method string, query string, variables map[string]interface{}) ([]byte, error) {
	var req *http.Request
	if method == constant.HttpMethodGet {
		req, _ = http.NewRequest(http.MethodGet, url, nil)
		if variables != nil && len(variables) > 0 {
			q := req.URL.Query()
			for key, value := range variables {
				q.Add(key, fmt.Sprintf("%v", value))
			}
			req.URL.RawQuery = q.Encode()
		}

	} else {
		in := struct {
			Query     string                 `json:"query"`
			Variables map[string]interface{} `json:"variables,omitempty"`
		}{
			Query:     query,
			Variables: variables,
		}
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(in)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, &buf)
		req.Header.Add("Content-Type", "application/json")
	}

	//增加header选项
	if !strings.HasPrefix(query, "query accessToken") {
		token, _ := GetAccessToken(c)
		req.Header.Add("Authorization", "Bearer "+token)
	}
	req.Header.Add("x-authing-userpool-id", ""+c.userPoolId)
	req.Header.Add("x-authing-request-from", constant.SdkType)
	req.Header.Add("x-authing-sdk-version", constant.SdkVersion)
	req.Header.Add("x-authing-app-id", ""+constant.AppId)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, nil
}

func QueryAccessToken(client *Client) (*model.AccessTokenRes, error) {
	type Data struct {
		AccessToken model.AccessTokenRes `json:"accessToken"`
	}
	type Result struct {
		Data Data `json:"data"`
	}

	variables := map[string]interface{}{
		"userPoolId": client.userPoolId,
		"secret":     client.Secret,
	}

	b, err := client.SendHttpRequestManage(client.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.AccessTokenDocument, variables)
	if err != nil {
		return nil, err
	}
	var r Result
	if b != nil {
		json.Unmarshal(b, &r)
	}
	return &r.Data.AccessToken, nil
}

func GetAccessToken(client *Client) (string, error) {
	// 从缓存获取token
	cacheToken, b := cacheutil.GetCache(constant.TokenCacheKeyPrefix + client.userPoolId)
	if b && cacheToken != nil {
		return cacheToken.(string), nil
	}
	// 从服务获取token，加锁
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	cacheToken, b = cacheutil.GetCache(constant.TokenCacheKeyPrefix + client.userPoolId)
	if b && cacheToken != nil {
		return cacheToken.(string), nil
	}
	token, err := QueryAccessToken(client)
	if err != nil {
		return "", err
	}
	var expire = *(token.Exp) - time.Now().Unix() - 43200
	cacheutil.SetCache(constant.TokenCacheKeyPrefix+client.userPoolId, *token.AccessToken, time.Duration(expire*int64(time.Second)))
	return *token.AccessToken, nil
}
