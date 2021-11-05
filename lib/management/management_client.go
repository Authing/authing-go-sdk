package management

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"github.com/Authing/authing-go-sdk/lib/util/cacheutil"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Client is a client for interacting with the GraphQL API of `Authing`
type Client struct {
	HttpClient *http.Client
	userPoolId string
	secret     string
	Host       string

	// Log is called with various debug information.
	// To log to standard out, use:
	//  client.Log = func(s string) { log.Println(s) }
	Log func(s string)
}

func NewClient(userPoolId string, secret string, host ...string) *Client {
	var clientHost string
	if len(host) == 0 {
		clientHost = constant.CoreAuthingDefaultUrl
	} else {
		clientHost = host[0]
	}
	c := &Client{
		userPoolId: userPoolId,
		secret:     secret,
		Host:       clientHost,
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

// NewHttpClient creates a new Authing user endpoint GraphQL API client
func NewHttpClient(userPoolId string, appSecret string, isDev bool) *Client {
	c := &Client{
		userPoolId: userPoolId,
	}

	/*if c.Client == nil {
		var endpointURL string
		if isDev {
			endpointURL = constant.CoreEndPointDevUrl + "/graphql/v2"
		} else {
			endpointURL = constant.CoreEndPointProdUrl + "/graphql/v2"
		}
		accessToken, err := GetAccessToken(userPoolId, appSecret)
		if err != nil {
			log.Println(err)
			//return nil
		}
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken},
		)
		c.HttpClient = oauth2.NewClient(context.Background(), src)

		c.Client = graphql.NewClient(endpointURL, c.HttpClient)
	}*/

	return c
}

// NewOauthClient creates a new Authing oauth endpoint GraphQL API client
func NewOauthClient(userPoolId string, appSecret string, isDev bool) *Client {
	c := &Client{
		userPoolId: userPoolId,
	}

	/*if c.Client == nil {
		var endpointURL string
		if isDev {
			endpointURL = constant.CoreEndPointDevUrl
		} else {
			endpointURL = constant.CoreEndPointProdUrl
		}
		accessToken, err := GetAccessToken(userPoolId, appSecret)
		if err != nil {
			log.Println(err)
			return nil
		}

		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken},
		)

		httpClient := oauth2.NewClient(context.Background(), src)

		if isDev {
			endpointURL = constant.CoreEndPointDevUrl
		} else {
			endpointURL = constant.CoreEndPointProdUrl
		}

		c.Client = graphql.NewClient(endpointURL, httpClient)
	}*/

	return c
}

func (c *Client) SendHttpRequest(url string, method string, query string, variables map[string]interface{}) ([]byte, error) {
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
		var err error
		if query == constant.StringEmpty {
			err = json.NewEncoder(&buf).Encode(variables)
		} else {
			err = json.NewEncoder(&buf).Encode(in)
		}
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

func (c *Client) SendHttpRestRequest(url string, method string, variables map[string]interface{}) ([]byte, error) {
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

	token, _ := GetAccessToken(c)
	req.Header.Add("Authorization", "Bearer "+token)

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

func (c *Client) httpGet(url string, client *http.Client) (string, error) {
	reqest, err := http.NewRequest(constant.HttpMethodGet, c.Host+url, nil)
	if err != nil {
		return "", err
	}

	//增加header选项
	token, _ := GetAccessToken(c)
	reqest.Header.Add("Authorization", "Bearer "+token)
	reqest.Header.Add("x-authing-userpool-id", ""+c.userPoolId)
	reqest.Header.Add("x-authing-request-from", constant.SdkType)
	reqest.Header.Add("x-authing-sdk-version", constant.SdkVersion)
	reqest.Header.Add("x-authing-app-id", ""+constant.AppId)

	resp, err := client.Do(reqest)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	result := string(body)
	return result, nil
}

func (c *Client) SendHttpRequestV2(url string, method string, query string, variables map[string]interface{}) ([]byte, error) {
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
	req := fasthttp.AcquireRequest()

	req.SetRequestURI(url)
	token, _ := GetAccessToken(c)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("x-authing-userpool-id", ""+c.userPoolId)
	req.Header.Add("x-authing-request-from", constant.SdkType)
	req.Header.Add("x-authing-sdk-version", constant.SdkVersion)
	req.Header.Add("x-authing-app-id", ""+constant.AppId)
	req.Header.SetMethod(method)
	req.SetBody(buf.Bytes())

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)
	body := resp.Body()
	return body, err
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
		"secret":     client.secret,
	}

	b, err := client.SendHttpRequest(client.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.AccessTokenDocument, variables)
	if err != nil {
		return nil, err
	}
	var r Result
	if b != nil {
		json.Unmarshal(b, &r)
	}
	log.Println(string(b))
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

func CreateRequestParam(param struct{}) map[string]interface{} {
	data, _ := json.Marshal(&param)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	return variables
}

// SendEmail
// 发送邮件
func (c *Client) SendEmail(email string, scene model.EnumEmailScene) (*model.CommonMessageAndCode, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.SendMailDocument,
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

// CheckLoginStatusByToken
// 检测登录状态
func (c *Client) CheckLoginStatusByToken(token string) (*model.CheckLoginStatusResponse, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.CheckLoginStatusDocument,
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

// IsPasswordValid
// 检测密码是否合法
func (c *Client) IsPasswordValid(password string) (*struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}, error) {

	url := fmt.Sprintf("%s/api/v2/password/check", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, map[string]interface{}{"password": password})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
		Data    struct {
			Valid   bool   `json:"valid"`
			Message string `json:"message"`
		} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}
