package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

// CreateNamespace
// 创建权限分组
func (c *Client) CreateNamespace(request *model.EditNamespaceRequest) (*model.Namespace, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	url := fmt.Sprintf("%s/api/v2/resource-namespace/%s", c.Host, c.userPoolId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string          `json:"message"`
		Code    int64           `json:"code"`
		Data    model.Namespace `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// UpdateNamespace
// 修改权限分组
func (c *Client) UpdateNamespace(id string, request *model.EditNamespaceRequest) (*model.Namespace, error) {

	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/resource-namespace/%s/%s", c.Host, c.userPoolId, id)
	b, err := c.SendHttpRestRequest(url, http.MethodPut, variables)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string          `json:"message"`
		Code    int64           `json:"code"`
		Data    model.Namespace `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// DeleteNamespace
// 删除权限分组
func (c *Client) DeleteNamespace(id string) (*string, error) {

	url := fmt.Sprintf("%s/api/v2/resource-namespace/%s/%s", c.Host, c.userPoolId, id)
	b, err := c.SendHttpRestRequest(url, http.MethodDelete, nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Message, nil
}

// ListNamespace
// 权限分组列表
func (c *Client) ListNamespace(page, limit int) (*struct {
	List  []model.Namespace `json:"list"`
	Total int64             `json:"total"`
}, error) {

	url := fmt.Sprintf("%s/api/v2/resource-namespace/%s?page=%v&limit=%v", c.Host, c.userPoolId, page, limit)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
		Data    struct {
			List  []model.Namespace `json:"list"`
			Total int64             `json:"total"`
		} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}
