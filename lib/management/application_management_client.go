package management

import (
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
)

// ListApplication
// 获取应用列表
func (c *Client) ListApplication(req *model.CommonPageRequest) (*struct {
	List []model.Application `json:"list"`
}, error) {

	url := fmt.Sprintf("%v/api/v2/applications?page=%v&limit=%v", c.Host, req.Page, req.Limit)
	b, err := c.SendHttpRequest(url, constant.HttpMethodGet, "", nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
		Data    struct {
			List []model.Application `json:"list"`
		} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// CreateApplication
// 创建应用
func (c *Client) CreateApplication(name, identifier, redirectUris string, logo *string) (*model.Application, error) {
	vars := make(map[string]interface{})
	vars["name"] = name
	vars["identifier"] = identifier
	vars["redirectUris"] = redirectUris
	if logo != nil {
		vars["logo"] = logo
	}
	url := fmt.Sprintf("%v/api/v2/applications", c.Host)
	b, err := c.SendHttpRequest(url, constant.HttpMethodPost, "", vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string            `json:"message"`
		Code    int64             `json:"code"`
		Data    model.Application `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// DeleteApplication
// 删除应用
func (c *Client) DeleteApplication(appId string) (*string, error) {
	url := fmt.Sprintf("%v/api/v2/applications/%v", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodDelete, nil)
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
	return &resp.Message, nil
}

// RefreshApplicationSecret
// 刷新应用密钥
func (c *Client) RefreshApplicationSecret(appId string) (*model.Application, error) {
	url := fmt.Sprintf("%s/api/v2/application/%s/refresh-secret", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodPatch, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string            `json:"message"`
		Code    int64             `json:"code"`
		Data    model.Application `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// ListApplicationActiveUsers
// 查看应用下已登录用户
func (c *Client) ListApplicationActiveUsers(appId string, page, limit int) (*struct {
	List       []model.ApplicationActiveUsers `json:"list"`
	TotalCount int64                          `json:"totalCount"`
}, error) {
	url := fmt.Sprintf("%s/api/v2/applications/%s/active-users?page=%v&%v", c.Host, appId, page, limit)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
		Data    struct {
			List       []model.ApplicationActiveUsers `json:"list"`
			TotalCount int64                          `json:"totalCount"`
		} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// FindApplicationById
// 通过应用 id 查找应用详情
func (c *Client) FindApplicationById(appId string) (*model.Application, error) {
	url := fmt.Sprintf("%s/api/v2/applications/%s", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string            `json:"message"`
		Code    int64             `json:"code"`
		Data    model.Application `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// CreateApplicationAgreement
// 创建应用协议
func (c *Client) CreateApplicationAgreement(appId, title string, lang *string, required *bool) (*model.ApplicationAgreement, error) {
	if lang == nil {
		var def = "zh-CN"
		lang = &def
	}
	if required == nil {
		var def = true
		required = &def
	}
	vars := map[string]interface{}{
		"title":    title,
		"lang":     lang,
		"required": required,
	}
	url := fmt.Sprintf("%s/api/v2/applications/%s/agreements", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                     `json:"message"`
		Code    int64                      `json:"code"`
		Data    model.ApplicationAgreement `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// ListApplicationAgreement
// 应用协议列表
func (c *Client) ListApplicationAgreement(appId string) (*struct {
	List       []model.ApplicationAgreement `json:"list"`
	TotalCount int64                        `json:"totalCount"`
}, error) {

	url := fmt.Sprintf("%s/api/v2/applications/%s/agreements", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
		Data    struct {
			List       []model.ApplicationAgreement `json:"list"`
			TotalCount int64                        `json:"totalCount"`
		} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// ModifyApplicationAgreement
// 修改应用协议
func (c *Client) ModifyApplicationAgreement(appId, agreementId, title string, lang *string, required *bool) (*model.ApplicationAgreement, error) {
	if lang == nil {
		var def = "zh-CN"
		lang = &def
	}
	if required == nil {
		var def = true
		required = &def
	}
	vars := map[string]interface{}{
		"title":    title,
		"lang":     lang,
		"required": required,
	}
	url := fmt.Sprintf("%s/api/v2/applications/%s/agreements/%v", c.Host, appId, agreementId)
	b, err := c.SendHttpRestRequest(url, http.MethodPut, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                     `json:"message"`
		Code    int64                      `json:"code"`
		Data    model.ApplicationAgreement `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// DeleteApplicationAgreement
// 删除应用协议
func (c *Client) DeleteApplicationAgreement(appId, agreementId string) (*string, error) {

	url := fmt.Sprintf("%s/api/v2/applications/%s/agreements/%v", c.Host, appId, agreementId)
	b, err := c.SendHttpRestRequest(url, http.MethodDelete, nil)
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
	return &resp.Message, nil
}

// SortApplicationAgreement
// 排序应用协议
func (c *Client) SortApplicationAgreement(appId string, ids []string) (*string, error) {

	url := fmt.Sprintf("%s/api/v2/applications/%s/agreements/sort", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, map[string]interface{}{"ids": ids})
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
	return &resp.Message, nil
}
