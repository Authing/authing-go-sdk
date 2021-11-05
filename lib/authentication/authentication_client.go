package authentication

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"github.com/Authing/authing-go-sdk/lib/util"
	"github.com/Authing/authing-go-sdk/lib/util/cacheutil"
	simplejson "github.com/bitly/go-simplejson"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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

	ClientToken *string
	ClientUser  *model.User

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

	//c.AuthingRequest = util.NewAuthingRequest(appId,secret,clientHost)
	return c
}

// TODO
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

// GetAccessTokenByCode
//  code 换取 accessToken
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

	//resp, err := c.AuthingRequest.SendRequest(url, constant.HttpMethodPost, header, body)
	return string(resp), err
}

// GetUserInfoByAccessToken
// accessToken 换取用户信息
func (c *Client) GetUserInfoByAccessToken(accessToken string) (string, error) {
	if accessToken == constant.StringEmpty {
		return constant.StringEmpty, errors.New("accessToken 不能为空")
	}
	url := c.Host + "/oidc/me?access_token=" + accessToken
	resp, err := c.SendHttpRequest(url, constant.HttpMethodGet, nil, nil)
	return string(resp), err
}

// GetNewAccessTokenByRefreshToken
//   使用 Refresh token 获取新的 Access token
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

// IntrospectToken
// 检查 Access token 或 Refresh token 的状态
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

// ValidateToken
// 效验Token合法性
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

// RevokeToken
// 撤回 Access token 或 Refresh token
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

// GetAccessTokenByClientCredentials
// Client Credentials 模式获取 Access Token
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

// LoginByUserName
// 使用用户名登录
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
	return c.loginGetUserInfo(b, "loginByUsername")
}

// LoginByEmail
// 使用邮箱登录
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
	return c.loginGetUserInfo(b, "loginByEmail")
}

// LoginByPhonePassword
// 使用手机号密码登录
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
	return c.loginGetUserInfo(b, "loginByPhonePassword")
}

//TODO
func (c *Client) loginGetUserInfo(b []byte, userKey string) (*model.User, error) {
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
	c.SetCurrentUser(&resultUser)
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
	if !strings.HasPrefix(query, "query accessToken") && c.ClientToken != nil {
		token := c.ClientToken
		req.Header.Add("Authorization", "Bearer "+*token)
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

//TODO
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

// GetAccessToken
// 获取访问Token
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

func (c *Client) SendHttpRestRequest(url string, method string, token *string, variables map[string]interface{}) ([]byte, error) {
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
		var buf bytes.Buffer
		var err error
		if variables != nil {
			err = json.NewEncoder(&buf).Encode(variables)
		}
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, &buf)
		req.Header.Add("Content-Type", "application/json")
	}

	if token == nil {
		selfToken, _ := GetAccessToken(c)
		token = &selfToken
	}
	req.Header.Add("Authorization", "Bearer "+*token)

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

func (c *Client) SendHttpRestRequestNotToken(url string, method string, variables map[string]interface{}) ([]byte, error) {
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
		var buf bytes.Buffer
		var err error
		if variables != nil {
			err = json.NewEncoder(&buf).Encode(variables)
		}
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, &buf)
		req.Header.Add("Content-Type", "application/json")
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

// GetCurrentUser
// 获取资源列表
func (c *Client) GetCurrentUser(token *string) (*model.User, error) {

	url := fmt.Sprintf("%s/api/v2/users/me", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, token, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string     `json:"message"`
		Code    int64      `json:"code"`
		Data    model.User `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

func (c *Client) getCurrentUser() (*model.User, error) {
	k, e := cacheutil.GetCache(constant.UserCacheKeyPrefix + c.userPoolId)
	if !e {
		return nil, errors.New("未登录")
	}
	return k.(*model.User), nil
}

// SetCurrentUser
// 设置当前用户
func (c *Client) SetCurrentUser(user *model.User) (*model.User, error) {
	c.ClientUser = user
	c.ClientToken = user.Token
	//cacheutil.SetDefaultCache(constant.UserCacheKeyPrefix+c.userPoolId, user)
	//c.SetToken(*user.Token)

	return user, nil
}

// SetToken
// 设置 Token
func (c *Client) SetToken(token string) {
	c.ClientToken = &token
	//cacheutil.SetDefaultCache(constant.TokenCacheKeyPrefix+c.userPoolId, token)
}

// RegisterByEmail
// 使用邮箱注册
func (c *Client) RegisterByEmail(request *model.RegisterByEmailInput) (*model.User, error) {
	request.Password = util.RsaEncrypt(request.Password)
	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RegisterByEmailDocument,
		map[string]interface{}{"input": variables})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			RegisterByEmail model.User `json:"registerByEmail"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.RegisterByEmail)
	return &response.Data.RegisterByEmail, nil
}

// RegisterByUsername
// 使用用户名注册
func (c *Client) RegisterByUsername(request *model.RegisterByUsernameInput) (*model.User, error) {
	request.Password = util.RsaEncrypt(request.Password)
	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RegisterByUsernameDocument,
		map[string]interface{}{"input": variables})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			RegisterByUsername model.User `json:"registerByUsername"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.RegisterByUsername)
	return &response.Data.RegisterByUsername, nil
}

// RegisterByPhoneCode
// 使用手机号及验证码注册
func (c *Client) RegisterByPhoneCode(request *model.RegisterByPhoneCodeInput) (*model.User, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RegisterByPhoneCodeDocument,
		map[string]interface{}{"input": variables})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			RegisterByPhoneCode model.User `json:"registerByPhoneCode"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.RegisterByPhoneCode)
	return &response.Data.RegisterByPhoneCode, nil
}

// CheckPasswordStrength
// 检查密码强度
func (c *Client) CheckPasswordStrength(password string) (*struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}, error) {

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.CheckPasswordStrengthDocument,
		map[string]interface{}{"password": password})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			CheckPasswordStrength struct {
				Valid   bool   `json:"valid"`
				Message string `json:"message"`
			} `json:"checkPasswordStrength"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}

	return &response.Data.CheckPasswordStrength, nil
}

// SendSmsCode
// 发送短信验证码
func (c *Client) SendSmsCode(phone string) (*struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}, error) {

	url := fmt.Sprintf("%s/api/v2/sms/send", c.Host)
	b, err := c.SendHttpRestRequestNotToken(url, http.MethodPost, map[string]interface{}{
		"phone": phone,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// LoginByPhoneCode
// 使用手机号验证码登录
func (c *Client) LoginByPhoneCode(req *model.LoginByPhoneCodeInput) (*model.User, error) {
	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.LoginByPhoneCodeDocument, map[string]interface{}{
		"input": vars,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			LoginByPhoneCode model.User `json:"loginByPhoneCode"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}

	return &response.Data.LoginByPhoneCode, nil
}

// CheckLoginStatus
// 检测 Token 登录状态
func (c *Client) CheckLoginStatus(token string) (*model.CheckLoginStatusResponse, error) {

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.CheckLoginStatusDocument,
		map[string]interface{}{"token": token})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			CheckLoginStatus model.CheckLoginStatusResponse `json:"checkLoginStatus"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.CheckLoginStatus, nil
}

// SendEmail
// 发送邮件
func (c *Client) SendEmail(email string, scene model.EnumEmailScene) (*model.CommonMessageAndCode, error) {

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.SendMailDocument,
		map[string]interface{}{"email": email, "scene": scene})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			SendMail model.CommonMessageAndCode `json:"sendEmail"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.SendMail, nil
}

// ResetPasswordByPhoneCode
// 通过短信验证码重置密码
func (c *Client) ResetPasswordByPhoneCode(phone, code, newPassword string) (*model.CommonMessageAndCode, error) {

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ResetPasswordDocument,
		map[string]interface{}{"phone": phone, "code": code, "newPassword": util.RsaEncrypt(newPassword)})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			ResetPassword model.CommonMessageAndCode `json:"resetPassword"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.ResetPassword, nil
}

// ResetPasswordByEmailCode
// 通过邮件验证码重置密码
func (c *Client) ResetPasswordByEmailCode(email, code, newPassword string) (*model.CommonMessageAndCode, error) {

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ResetPasswordDocument,
		map[string]interface{}{"email": email, "code": code, "newPassword": util.RsaEncrypt(newPassword)})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			ResetPassword model.CommonMessageAndCode `json:"resetPassword"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.ResetPassword, nil
}

// UpdateProfile
// 修改用户资料
func (c *Client) UpdateProfile(req *model.UpdateUserInput) (*model.User, error) {
	vars := make(map[string]interface{})
	currentUser, e := c.getCurrentUser()
	if e != nil {
		return nil, e
	}
	vars["id"] = currentUser.Id
	vars["input"] = req

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UpdateProfileDocument,
		vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			UpdateUser model.User `json:"updateUser"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.UpdateUser)
	return &response.Data.UpdateUser, nil
}

// UpdatePassword
// 更新用户密码
func (c *Client) UpdatePassword(oldPassword *string, newPassword string) (*model.User, error) {

	vars := make(map[string]interface{})
	vars["newPassword"] = util.RsaEncrypt(newPassword)
	if oldPassword != nil {
		vars["oldPassword"] = util.RsaEncrypt(*oldPassword)
	}

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UpdatePasswordDocument, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			UpdatePassword model.User `json:"updatePassword"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.UpdatePassword)
	return &response.Data.UpdatePassword, nil
}

// UpdatePhone
// 更新用户手机号
func (c *Client) UpdatePhone(phone, code string, oldPhone, oldPhoneCode *string) (*model.User, error) {

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UpdatePhoneDocument,
		map[string]interface{}{
			"phone":        phone,
			"phoneCode":    code,
			"oldPhone":     oldPhone,
			"oldPhoneCode": oldPhoneCode,
		})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			UpdatePhone model.User `json:"updatePhone"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.UpdatePhone)
	return &response.Data.UpdatePhone, nil
}

// UpdateEmail
// 更新用户邮箱
func (c *Client) UpdateEmail(email, code string, oldEmail, oldEmailCode *string) (*model.User, error) {

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UpdateEmailDocument,
		map[string]interface{}{
			"email":        email,
			"emailCode":    code,
			"oldEmail":     oldEmail,
			"oldEmailCode": oldEmailCode,
		})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			UpdateEmail model.User `json:"updateEmail"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.UpdateEmail)
	return &response.Data.UpdateEmail, nil
}

// RefreshToken
// 刷新当前用户的 token
func (c *Client) RefreshToken(token *string) (*model.RefreshToken, error) {

	b, err := c.SendHttpRequestCustomTokenManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, token, constant.RefreshUserTokenDocument,
		nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			RefreshToken model.RefreshToken `json:"refreshToken"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetToken(*response.Data.RefreshToken.Token)
	return &response.Data.RefreshToken, nil
}

func (c *Client) SendHttpRequestCustomTokenManage(url string, method string, token *string, query string, variables map[string]interface{}) ([]byte, error) {
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
	if token == nil {
		useToken, _ := GetAccessToken(c)
		req.Header.Add("Authorization", "Bearer "+useToken)
	} else {
		req.Header.Add("Authorization", "Bearer "+*token)

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

// LinkAccount
// 关联账号
func (c *Client) LinkAccount(primaryUserToken, secondaryUserToken string) (*model.CommonMessageAndCode, error) {

	url := fmt.Sprintf("%s/api/v2/users/link", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, nil, map[string]interface{}{
		"primaryUserToken":   primaryUserToken,
		"secondaryUserToken": secondaryUserToken,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := model.CommonMessageAndCode{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp, nil
}

// UnLinkAccount
// 主账号解绑社会化登录账号
func (c *Client) UnLinkAccount(primaryUserToken string, provider constant.SocialProviderType) (*model.CommonMessageAndCode, error) {

	url := fmt.Sprintf("%s/api/v2/users/unlink", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, nil, map[string]interface{}{
		"primaryUserToken": primaryUserToken,
		"provider":         provider,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := model.CommonMessageAndCode{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp, nil
}

// BindPhone
// 绑定手机号
func (c *Client) BindPhone(phone, phoneCode string) (*model.User, error) {

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.BindPhoneDocument,
		map[string]interface{}{"phone": phone, "phoneCode": phoneCode})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			BindPhone model.User `json:"bindPhone"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.BindPhone)
	return &response.Data.BindPhone, nil
}

// UnBindPhone
// 绑定手机号
func (c *Client) UnBindPhone() (*model.User, error) {
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UnBindPhoneDocument,
		nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			UnbindPhone model.User `json:"unbindPhone"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.UnbindPhone)
	return &response.Data.UnbindPhone, nil
}

// BindEmail
// 绑定邮箱号
func (c *Client) BindEmail(email, emailCode string) (*model.User, error) {

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.BindEmailDocument,
		map[string]interface{}{
			"email":     email,
			"emailCode": emailCode,
		})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			BindEmail model.User `json:"bindEmail"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.BindEmail)
	return &response.Data.BindEmail, nil
}

// UnBindEmail
// 解绑邮箱号
func (c *Client) UnBindEmail() (*model.User, error) {
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UnBindEmailDocument,
		nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			UnbindEmail model.User `json:"unbindEmail"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.UnbindEmail)
	return &response.Data.UnbindEmail, nil
}

// Logout
// 退出登录
func (c *Client) Logout() (*model.CommonMessageAndCode, error) {
	cacheToken, _ := cacheutil.GetCache(constant.TokenCacheKeyPrefix + c.userPoolId)
	if cacheToken == nil {
		return nil, errors.New("Please login first")
	}
	token := cacheToken.(string)

	url := fmt.Sprintf("%s/api/v2/logout?app_id=%s", c.Host, c.AppId)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, &token, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := model.CommonMessageAndCode{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	c.ClearUser()
	return &resp, nil
}

func (c *Client) ClearUser() {
	c.ClientUser = nil
	c.ClientToken = nil
	//cacheutil.DeleteCache(constant.TokenCacheKeyPrefix + c.userPoolId)
	//cacheutil.DeleteCache(constant.UserCacheKeyPrefix + c.userPoolId)
}

func (c *Client) getCacheUser() (*model.User, error) {
	//cache, _ := cacheutil.GetCache(constant.UserCacheKeyPrefix + c.userPoolId)
	//if cache == nil {
	//	return nil, errors.New("Please login first")
	//}
	//cacheUser := cache.(*model.User)
	if c.ClientUser == nil {
		return nil, errors.New("Please login first")
	}
	return c.ClientUser, nil
}

// ListUdv
// 获取当前用户的自定义数据列表
func (c *Client) ListUdv() (*[]model.UserDefinedData, error) {
	cacheUser, e := c.getCacheUser()
	if e != nil {
		return nil, e
	}
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UdvDocument, map[string]interface{}{
		"targetType": model.EnumUDFTargetTypeUSER,
		"targetId":   cacheUser.Id,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			Udv []model.UserDefinedData `json:"udv"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Udv, nil
}

// SetUdv
// 添加自定义数据
func (c *Client) SetUdv(udvList []model.KeyValuePair) (*[]model.UserDefinedData, error) {
	cacheUser, e := c.getCacheUser()
	if e != nil {
		return nil, e
	}
	variables := make(map[string]interface{})

	variables["targetType"] = model.EnumUDFTargetTypeUSER
	variables["targetId"] = cacheUser.Id
	variables["udvList"] = udvList

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.SetRoleUdfValueDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			SetUdvBatch []model.UserDefinedData `json:"setUdvBatch"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.SetUdvBatch, nil
}

// RemoveUdv
// 删除自定义数据
func (c *Client) RemoveUdv(key string) (*[]model.UserDefinedData, error) {
	cacheUser, e := c.getCacheUser()
	if e != nil {
		return nil, e
	}
	variables := make(map[string]interface{})
	variables["targetType"] = constant.USER
	variables["targetId"] = cacheUser.Id
	variables["key"] = key

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RemoveUdfValueDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			RemoveUdv []model.UserDefinedData `json:"removeUdv"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.RemoveUdv, nil
}

// ListOrg
// 获取用户所在组织机构
func (c *Client) ListOrg() (*struct {
	Code    int64            `json:"code"`
	Message string           `json:"message"`
	Data    []model.UserOrgs `json:"data"`
}, error) {

	if c.ClientToken == nil {
		return nil, errors.New("Please login first")
	}
	token := c.ClientToken

	url := fmt.Sprintf("%s/api/v2/users/me/orgs", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, token, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Code    int64            `json:"code"`
		Message string           `json:"message"`
		Data    []model.UserOrgs `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// LoginByLdap
// 使用 LDAP 用户名登录
func (c *Client) LoginByLdap(username, password string) (*struct {
	Code    int64      `json:"code"`
	Message string     `json:"message"`
	Data    model.User `json:"data"`
}, error) {

	url := fmt.Sprintf("%s/api/v2/ldap/verify-user", c.Host)
	b, err := c.SendHttpRestRequestNotToken(url, http.MethodPost, map[string]interface{}{
		"username": username,
		"password": password,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Code    int64      `json:"code"`
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// LoginByAd
// 使用 AD 用户名登录
func (c *Client) LoginByAd(username, password string) (*struct {
	Code    int64      `json:"code"`
	Message string     `json:"message"`
	Data    model.User `json:"data"`
}, error) {

	com, _ := regexp.Compile("(?:http.*://)?(?P<host>[^:/ ]+).?(?P<port>[0-9]*).*")
	domain := com.FindString(c.Host)

	lis := strings.Split(domain, ".")
	var wsHost string
	if len(lis) > 2 {
		wsHost = strings.Join(lis[1:], ".")
	} else {
		wsHost = domain
	}
	url := fmt.Sprintf("https://ws.%s/api/v2/ad/verify-user", wsHost)
	b, err := c.SendHttpRestRequestNotToken(url, http.MethodPost, map[string]interface{}{
		"username": username,
		"password": password,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Code    int64      `json:"code"`
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// GetSecurityLevel
// 用户安全等级
func (c *Client) GetSecurityLevel() (*struct {
	Code    int64                          `json:"code"`
	Message string                         `json:"message"`
	Data    model.GetSecurityLevelResponse `json:"data"`
}, error) {
	//cacheToken, _ := cacheutil.GetCache(constant.TokenCacheKeyPrefix + c.userPoolId)
	//if cacheToken == nil {
	//	return nil, errors.New("Please login first")
	//}
	//token := cacheToken.(string)
	if c.ClientToken == nil {
		return nil, errors.New("Please login first")
	}
	token := c.ClientToken
	url := fmt.Sprintf("%s/api/v2/users/me/security-level", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, token, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Code    int64                          `json:"code"`
		Message string                         `json:"message"`
		Data    model.GetSecurityLevelResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// ListAuthorizedResources
// 获取用户被授权的所有资源
func (c *Client) ListAuthorizedResources(namespace string, resourceType model.EnumResourceType) (*model.AuthorizedResources, error) {
	cacheUser, e := c.getCacheUser()
	if e != nil {
		return nil, e
	}
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListUserAuthorizedResourcesDocument,
		map[string]interface{}{
			"id":           cacheUser.Id,
			"namespace":    namespace,
			"resourceType": resourceType,
		})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			User struct {
				AuthorizedResources model.AuthorizedResources `json:"authorizedResources"`
			} `json:"user"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.User.AuthorizedResources, nil
}

func (c *Client) BuildAuthorizeUrlByOauth(scope, redirectUri, state, responseType string) (string, error) {

	if c.AppId == "" {
		return constant.StringEmpty, errors.New("请在初始化 AuthenticationClient 时传入 appId")
	}
	if c.Protocol != constant.OAUTH {
		return constant.StringEmpty, errors.New("初始化 AuthenticationClient 传入的 protocol 应为 ProtocolEnum.OAUTH")
	}
	if redirectUri == "" {
		return constant.StringEmpty, errors.New("redirectUri 不能为空")
	}

	if strings.Contains(scope, "offline_access") {
		scope = "consent"
	}
	dataMap := map[string]string{
		"client_id":     util.GetValidValue(c.AppId),
		"scope":         util.GetValidValue(scope, "openid profile email phone address"),
		"state":         util.GetValidValue(state, util.RandomString(12)),
		"response_type": util.GetValidValue(responseType),
		"redirect_uri":  util.GetValidValue(redirectUri),
	}
	return c.Host + "/oauth/auth?" + util.GetQueryString(dataMap), nil
}

func (c *Client) BuildAuthorizeUrlBySaml() string {
	return fmt.Sprintf("%s/api/v2/saml-idp/%s", c.Host, c.AppId)
}

func (c *Client) BuildAuthorizeUrlByCas(service *string) string {
	if service != nil {
		return fmt.Sprintf("%s/cas-idp/%s?service=%s", c.Host, c.AppId, *service)
	} else {
		return fmt.Sprintf("%s/cas-idp/%s?service", c.Host, c.AppId)
	}
}

// ValidateTicketV1
// 检验 CAS 1.0 Ticket 合法性
func (c *Client) ValidateTicketV1(ticket, service string) (*struct {
	Valid    bool   `json:"code"`
	Message  string `json:"message"`
	Username string `json:"username"`
}, error) {

	url := fmt.Sprintf("%s/cas-idp/%s/validate?service=%s&ticket=%s", c.Host, c.AppId, service, ticket)
	b, err := c.SendHttpRestRequestNotToken(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	sps := strings.Split(string(b), "\n")
	var username, message string

	valid := (sps[0] == "yes")
	username = sps[1]
	if !valid {
		message = "ticket is not valid"
	}
	resp := &struct {
		Valid    bool   `json:"code"`
		Message  string `json:"message"`
		Username string `json:"username"`
	}{
		Valid:    valid,
		Username: username,
		Message:  message,
	}

	return resp, nil
}

//BuildLogoutUrl
//拼接登出 URL
func (c *Client) BuildLogoutUrl(expert, redirectUri, idToken *string) string {
	var url string
	if c.Protocol == constant.OIDC {
		if expert == nil {
			if redirectUri != nil {
				url = fmt.Sprintf("%s/login/profile/logout?redirect_uri=%s", c.Host, *redirectUri)
			} else {
				url = fmt.Sprintf("%s/login/profile/logout", c.Host)
			}

		} else {
			if redirectUri != nil {
				url = fmt.Sprintf("%s/oidc/session/end?id_token_hint=%s&post_logout_redirect_uri=%s", c.Host, *idToken, *redirectUri)
			} else {
				url = fmt.Sprintf("%s/oidc/session/end", c.Host)
			}

		}
	}
	if c.Protocol == constant.CAS {
		if redirectUri != nil {
			url = fmt.Sprintf("%s/cas-idp/logout?url=%s", c.Host, *redirectUri)
		} else {
			url = fmt.Sprintf("%s/cas-idp/logout", c.Host)
		}
	}
	return url
}

// ListRole
// 获取用户拥有的角色列表
func (c *Client) ListRole(namespace string) (*struct {
	TotalCount int               `json:"totalCount"`
	List       []model.RoleModel `json:"list"`
}, error) {
	cacheUser, e := c.getCacheUser()
	if e != nil {
		return nil, e
	}
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.GetUserRolesDocument,
		map[string]interface{}{
			"id":        cacheUser.Id,
			"namespace": namespace,
		})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			User model.GetUserRolesResponse `json:"user"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.User.Roles, nil
}

// HasRole
// 判断当前用户是否有某个角色
func (c *Client) HasRole(code, namespace string) (*bool, error) {
	r, e := c.ListRole(namespace)
	if e != nil {
		return nil, e
	}
	hasRole := true
	notHas := false
	if r.TotalCount == 0 {
		return &notHas, nil
	}
	for _, d := range r.List {
		if d.Code == code {
			return &hasRole, nil
		}
	}
	return &notHas, nil
}

// ListApplications
// 获取当前用户能够访问的应用
func (c *Client) ListApplications(page, limit int) (*struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TotalCount int64               `json:"totalCount"`
		List       []model.Application `json:"list"`
	} `json:"data"`
}, error) {
	if c.ClientToken == nil {
		return nil, errors.New("Please login first")
	}
	token := c.ClientToken
	url := fmt.Sprintf("%s/api/v2/users/me/applications/allowed?page=%v&limit=%v", c.Host, page, limit)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, token, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
		Data    struct {
			TotalCount int64               `json:"totalCount"`
			List       []model.Application `json:"list"`
		} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// GenerateCodeChallenge
// 生成一个 PKCE 校验码，长度必须大于等于 43。
func (c *Client) GenerateCodeChallenge(size int) (string, error) {
	if size < 43 {
		return constant.StringEmpty, errors.New("code_challenge must be a string length grater than 43")
	}
	return util.RandomString(size), nil

}

// GetCodeChallengeDigest
// 生成一个 PKCE 校验码摘要值
func (c *Client) GetCodeChallengeDigest(codeChallenge string, method constant.GenerateCodeChallengeMethod) (string, error) {
	if len(codeChallenge) < 43 {
		return constant.StringEmpty, errors.New("code_challenge must be a string length grater than 43")
	}
	if method == constant.PLAIN {
		return codeChallenge, nil
	} else {
		hasher := sha256.New()
		hasher.Write([]byte(codeChallenge))
		base64Str := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		return strings.Replace(base64Str, "=", "", -1), nil
	}

}

// LoginBySubAccount
// 登录子账号
func (c *Client) LoginBySubAccount(req *model.LoginBySubAccountRequest) (*model.User, error) {
	req.Password = util.RsaEncrypt(req.Password)
	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.LoginBySubAccountDocument, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			LoginBySubAccount model.User `json:"loginBySubAccount"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	c.SetCurrentUser(&response.Data.LoginBySubAccount)
	return &response.Data.LoginBySubAccount, nil
}

// ResetPasswordByFirstLoginToken
// 通过首次登录的 Token 重置密码
func (c *Client) ResetPasswordByFirstLoginToken(token, password string) (*model.CommonMessageAndCode, error) {
	password = util.RsaEncrypt(password)

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ResetPasswordByTokenDocument,
		map[string]interface{}{
			"token":    token,
			"password": password,
		})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			ResetPasswordByFirstLoginToken model.CommonMessageAndCode `json:"resetPasswordByFirstLoginToken"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}

	return &response.Data.ResetPasswordByFirstLoginToken, nil
}

// ResetPasswordByForceResetToken
// 通过密码强制更新临时 Token 修改密码
func (c *Client) ResetPasswordByForceResetToken(token, password, newPassword string) (*model.CommonMessageAndCode, error) {
	password = util.RsaEncrypt(password)
	newPassword = util.RsaEncrypt(newPassword)
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ResetPasswordByForceResetTokenDocument,
		map[string]interface{}{
			"token":       token,
			"oldPassword": password,
			"newPassword": newPassword,
		})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			ResetPasswordByForceResetToken model.CommonMessageAndCode `json:"resetPasswordByForceResetToken"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}

	return &response.Data.ResetPasswordByForceResetToken, nil
}

// ListDepartments
// 获取用户所有部门
func (c *Client) ListDepartments() (*model.PaginatedDepartments, error) {
	cacheUser, e := c.getCacheUser()
	if e != nil {
		return nil, e
	}
	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.GetUserDepartmentsDocument,
		map[string]interface{}{"id": cacheUser.Id})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))

	var response = &struct {
		Data   model.UserDepartmentsData `json:"data"`
		Errors []model.GqlCommonErrors   `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return response.Data.User.Departments, nil

}

// IsUserExists
// 判断用户是否存在
func (c *Client) IsUserExists(req *model.IsUserExistsRequest) (*bool, error) {

	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)

	b, err := c.SendHttpRequestManage(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.IsUserExistsDocument,
		vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))

	var response = &struct {
		Data struct {
			IsUserExists *bool `json:"isUserExists"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return response.Data.IsUserExists, nil

}

// ValidateTicketV2
// 通过远端服务验证票据合法性
func (c *Client) ValidateTicketV2(ticket, service string, format constant.TicketFormat) (*struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}, error) {

	url := fmt.Sprintf("%s/cas-idp/%s/serviceValidate", c.Host, c.AppId)
	b, err := c.SendHttpRestRequestNotToken(url, http.MethodGet, map[string]interface{}{
		"service": service,
		"ticket":  ticket,
		"format":  format,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Code    int64       `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// TrackSession
// sso 检测登录态
func (c *Client) TrackSession(code string, country, lang, state *string) (*struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}, error) {

	url := fmt.Sprintf("%s/connection/social/wechat:mobile/%s/callback?code=%s", c.Host, c.AppId, code)
	if country != nil {
		url = url + "&country=" + *country
	}
	if lang != nil {
		url = url + "&lang=" + *lang
	}
	if state != nil {
		url = url + "&state=" + *state
	}
	b, err := c.SendHttpRestRequestNotToken(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Code    int64       `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}
