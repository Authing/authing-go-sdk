package management

import (
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

// PrincipalAuthDetail
// 获取主体认证详情
func (c *Client) PrincipalAuthDetail(userId string) (*struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
}, error) {

	url := fmt.Sprintf("%s/api/v2/users/%s/management/principal_authentication", c.Host, userId)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string      `json:"message"`
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// PrincipalAuthenticate
// 进行主体认证
func (c *Client) PrincipalAuthenticate(userId string, req *model.PrincipalAuthenticateRequest) (*struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
}, error) {
	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)
	url := fmt.Sprintf("%s/api/v2/users/%s/management/principal_authentication", c.Host, userId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, vars)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string      `json:"message"`
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}
