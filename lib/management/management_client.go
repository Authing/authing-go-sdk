package management

import (
	"authing-golang-sdk/lib/constant"
	"authing-golang-sdk/lib/model"
	"authing-golang-sdk/lib/util/cacheutil"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kelvinji2009/graphql"
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
	Client *graphql.Client
	HttpClient *http.Client
	userPoolId string
	secret string
	Host string

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
		secret: secret,
		Host: clientHost,
	}
	if c.HttpClient == nil {
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

func (c *Client)SendHttpRequest(url string, method string, query string, variables map[string]interface{}) ([]byte, error) {
	client := &http.Client {
	}
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
	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	//增加header选项
	if !strings.HasPrefix(query, "query accessToken") {
		token, _ := GetAccessToken(c)
		req.Header.Add("Authorization", "Bearer " + token)
	}
	req.Header.Add("x-authing-userpool-id", "" + c.userPoolId)
	req.Header.Add("x-authing-request-from", constant.SdkType)
	req.Header.Add("x-authing-sdk-version", constant.SdkVersion)
	req.Header.Add("x-authing-app-id", "" + constant.AppId)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	return body, nil
}


func (c *Client)httpGet(url string, client *http.Client) (string,error)  {
	reqest, err := http.NewRequest(constant.HttpMethodGet, c.Host +url, nil)
	if err != nil {
		return "",err
	}

	//增加header选项
	token, _ := GetAccessToken(c)
	reqest.Header.Add("Authorization", "Bearer " + token)
	reqest.Header.Add("x-authing-userpool-id", "" + c.userPoolId)
	reqest.Header.Add("x-authing-request-from", constant.SdkType)
	reqest.Header.Add("x-authing-sdk-version", constant.SdkVersion)
	reqest.Header.Add("x-authing-app-id", "" + constant.AppId)

	resp, err := client.Do(reqest)
	if err != nil {
		return "",err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	result := string(body)
	fmt.Printf(result)
	return result,nil
}

func (c *Client)SendHttpRequestV2(url string, method string, query string, variables map[string]interface{}) ([]byte, error) {
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
	req.Header.Add("Authorization", "Bearer " + token)
	req.Header.Add("x-authing-userpool-id", "" + c.userPoolId)
	req.Header.Add("x-authing-request-from", constant.SdkType)
	req.Header.Add("x-authing-sdk-version", constant.SdkVersion)
	req.Header.Add("x-authing-app-id", "" + constant.AppId)
	req.Header.SetMethod(method)
	req.SetBody(buf.Bytes())

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)
	body := resp.Body()
	return body,err
}


func QueryAccessToken(client *Client) (*model.AccessTokenRes, error) {
	type Data struct {
		AccessToken model.AccessTokenRes `json:"accessToken"`
	}
	type Result struct {
		Data Data `json:"data"`
	}

	variables := map[string]interface{}{
		"userPoolId":     client.userPoolId,
		"secret": 		client.secret,
	}

	b, err := client.SendHttpRequest(client.Host + constant.CoreAuthingGraphqlPath,constant.HttpMethodPost, constant.AccessTokenDocument, variables)
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
	cacheutil.SetCache(constant.TokenCacheKeyPrefix + client.userPoolId, *token.AccessToken, time.Duration(expire * int64(time.Second)))
	return *token.AccessToken, nil
}

