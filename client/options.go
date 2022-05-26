package client

import (
	"authing-go-sdk/constant"
	"authing-go-sdk/dto"
	"authing-go-sdk/util/cache"
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	HttpClient *http.Client
	options    *ManagementClientOptions
	userPoolId string
}

type ManagementClientOptions struct {
	AccessKeyId     string
	AccessKeySecret string
	TenantId        string
	Timeout         int
	RequestFrom     string
	Lang            string
	Host            string
	Headers         fasthttp.RequestHeader
}

func NewClient(options *ManagementClientOptions) (*Client, error) {
	if options.Host == "" {
		options.Host = constant.ApiServiceUrl
	}
	c := &Client{
		options: options,
	}
	if c.HttpClient == nil {
		c.HttpClient = &http.Client{}
		_, err := GetAccessToken(c)
		if err != nil {
			return nil, err
		}
		/*src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken},
		)
		c.HttpClient = oauth2.NewClient(context.Background(), src)*/
	}
	return c, nil
}

type JwtClaims struct {
	*jwt.StandardClaims
	//用户编号
	UID      string
	Username string
}

func GetAccessToken(client *Client) (string, error) {
	// 从缓存获取token
	cacheToken, b := cache.GetCache(constant.TokenCacheKeyPrefix + client.options.AccessKeyId)
	if b && cacheToken != nil {
		return cacheToken.(string), nil
	}
	// 从服务获取token，加锁
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	cacheToken, b = cache.GetCache(constant.TokenCacheKeyPrefix + client.options.AccessKeyId)
	if b && cacheToken != nil {
		return cacheToken.(string), nil
	}
	resp, err := QueryAccessToken(client)
	if err != nil {
		return "", err
	}
	/*var jwtclaim = &JwtClaims{}
	_, err := jwt.ParseWithClaims(resp.Data.AccessToken, &jwtclaim, func(*jwt.Token) (interface{}, error) {
		//得到盐
		return secret, nil
	})*/
	if token, _ := jwt.Parse(resp.Data.AccessToken, nil); token != nil {
		userPoolId := token.Claims.(jwt.MapClaims)["scoped_userpool_id"]
		client.userPoolId = userPoolId.(string)
	}
	//fmt.Println(token)
	//var expire = (*(token.Exp) - time.Now().Unix() - 259200) * int64(time.Second)
	// TODO 时间戳类型转换
	cache.SetCache(constant.TokenCacheKeyPrefix+client.options.AccessKeyId, resp.Data.AccessToken, time.Duration(resp.Data.ExpiresIn*int(time.Second)))
	return resp.Data.AccessToken, nil
}

func QueryAccessToken(client *Client) (*dto.GetManagementTokenRespDto, error) {
	variables := map[string]interface{}{
		"accessKeyId":     client.options.AccessKeyId,
		"accessKeySecret": client.options.AccessKeySecret,
	}

	b, err := client.SendHttpRequest("/api/v3/get-management-token", fasthttp.MethodPost, variables)
	if err != nil {
		return nil, err
	}
	var r dto.GetManagementTokenRespDto
	if b != nil {
		json.Unmarshal(b, &r)
	}
	return &r, nil
}

func (c *Client) SendHttpRequest(url string, method string, variables interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(variables)
	if err != nil {
		return nil, err
	}
	req := fasthttp.AcquireRequest()

	req.SetRequestURI(c.options.Host + url)

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("x-authing-app-tenant-id", ""+c.options.TenantId)
	req.Header.Add("x-authing-request-from", c.options.RequestFrom)
	req.Header.Add("x-authing-sdk-version", constant.SdkVersion)
	req.Header.Add("x-authing-lang", c.options.Lang)
	if url != "/api/v3/get-management-token" {
		token, _ := GetAccessToken(c)
		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Add("x-authing-userpool-id", c.userPoolId)
	}
	req.Header.SetMethod(method)
	req.SetBody(buf.Bytes())

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)
	body := resp.Body()
	return body, err
}
