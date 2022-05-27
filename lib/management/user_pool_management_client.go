package management

import (
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

// UserPoolDetail
// 查询用户池配置
func (c *Client) UserPoolDetail() (*model.UserPool, error) {

	url := fmt.Sprintf("%s/api/v2/userpools/detail", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string         `json:"message"`
		Code    int64          `json:"code"`
		Data    model.UserPool `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// UpdateUserPool
// 更新用户池配置
func (c *Client) UpdateUserPool(request model.UpdateUserpoolInput) (*model.UserPool, error) {
	variables := make(map[string]interface{})
	variables["input"] = request
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UpdateUserPoolDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			UpdateUserPool model.UserPool `json:"updateUserpool"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.UpdateUserPool, nil

}

// ListUserPoolEnv
// 获取环境变量列表
func (c *Client) ListUserPoolEnv() (*[]model.UserPoolEnv, error) {

	url := fmt.Sprintf("%s/api/v2/env", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string              `json:"message"`
		Code    int64               `json:"code"`
		Data    []model.UserPoolEnv `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// RemoveUserPoolEnv
// 移除环境变量列表
func (c *Client) RemoveUserPoolEnv(key string) (*model.CommonMessageAndCode, error) {

	url := fmt.Sprintf("%s/api/v2/env/%s", c.Host, key)
	b, err := c.SendHttpRestRequest(url, http.MethodDelete, nil)
	if err != nil {
		return nil, err
	}
	var resp model.CommonMessageAndCode
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp, nil
}

// AddUserPoolEnv
// 新增环境变量列表
func (c *Client) AddUserPoolEnv(key, value string) (*model.UserPoolEnv, error) {

	url := fmt.Sprintf("%s/api/v2/env", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, map[string]interface{}{
		"key": key, "value": value,
	})
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string            `json:"message"`
		Code    int64             `json:"code"`
		Data    model.UserPoolEnv `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}
