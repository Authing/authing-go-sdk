package authentication

import (
	"errors"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/util"
	"golang.org/x/oauth2"
	"log"
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
		accessToken, err := GetAccessToken(c)
		if err != nil {
			log.Println(err)
			return nil
		}
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken},
		)
		c.HttpClient = oauth2.NewClient(context.Background(), src)
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
	switch c.TokenEndPointAuthMethod {
	case constant.TokenEndPointAuthMethod:

	}
}
