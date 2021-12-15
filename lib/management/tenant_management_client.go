package management

import (
	"errors"
	"fmt"

	// "github.com/Authing/authing-go-sdk/lib/constant"
	"net/http"

	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
)

// GetTenantList
// 获取用户池下租户列表
func (c *Client) GetTenantList(request *model.CommonPageRequest) (*model.GetTenantListResponse, error) {

	url := fmt.Sprintf("%s/api/v2/tenants?page=%v&limit=%v", c.Host, request.Page, request.Limit)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string                      `json:"message"`
		Code    int64                       `json:"code"`
		Data    model.GetTenantListResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// GetTenantDetails
// 获取租户详情
func (c *Client) GetTenantDetails(tenantId string) (*model.TenantDetails, error) {

	url := fmt.Sprintf("%s/api/v2/tenant/%s", c.Host, tenantId)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string              `json:"message"`
		Code    int64               `json:"code"`
		Data    model.TenantDetails `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// CreateTenant
// 创建租户
func (c *Client) CreateTenant(request *model.CreateTenantRequest) (*model.TenantDetails, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/tenant", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string              `json:"message"`
		Code    int64               `json:"code"`
		Data    model.TenantDetails `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// UpdateTenant
// 修改租户
func (c *Client) UpdateTenant(tenantId string, request *model.CreateTenantRequest) (*bool, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/tenant/%s", c.Host, tenantId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
		Data    bool   `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// RemoveTenant
// 删除租户
func (c *Client) RemoveTenant(tenantId string) (*string, error) {

	url := fmt.Sprintf("%s/api/v2/tenant/%s", c.Host, tenantId)
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

// ConfigTenant
// 配置租户品牌化
func (c *Client) ConfigTenant(tenantId string, request *model.ConfigTenantRequest) (*string, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/tenant/%s", c.Host, tenantId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
		Data    bool   `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Message, nil
}

// GetTenantMembers
// 获取租户成员列表
func (c *Client) GetTenantMembers(tenantId string, request *model.CommonPageRequest) (*model.TenantMembersResponse, error) {

	url := fmt.Sprintf("%s/api/v2/tenant/%s/users?page=%v&limit=%v", c.Host, tenantId, request.Page, request.Limit)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string                      `json:"message"`
		Code    int64                       `json:"code"`
		Data    model.TenantMembersResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// AddTenantMembers
// 添加租户成员
func (c *Client) AddTenantMembers(tenantId string, userIds []string) (*model.AddTenantMembersResponse, error) {

	url := fmt.Sprintf("%s/api/v2/tenant/%s/user", c.Host, tenantId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, map[string]interface{}{
		"userIds": userIds,
	})
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string                         `json:"message"`
		Code    int64                          `json:"code"`
		Data    model.AddTenantMembersResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// RemoveTenantMembers
// 删除租户成员
func (c *Client) RemoveTenantMembers(tenantId string, userId string) (*string, error) {

	url := fmt.Sprintf("%s/api/v2/tenant/%s/user?userId=%s", c.Host, tenantId, userId)
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

// ListExtIdp
// 获取身份源列表
func (c *Client) ListExtIdp(tenantId string) (*[]model.ListExtIdpResponse, error) {

	url := fmt.Sprintf("%s/api/v2/extIdp?tenantId=%s", c.Host, tenantId)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string                     `json:"message"`
		Code    int64                      `json:"code"`
		Data    []model.ListExtIdpResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// ExtIdpDetail
// 获取身份源详细信息
func (c *Client) ExtIdpDetail(extIdpId string) (*model.ExtIdpDetailResponse, error) {

	url := fmt.Sprintf("%s/api/v2/extIdp/%s", c.Host, extIdpId)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string                     `json:"message"`
		Code    int64                      `json:"code"`
		Data    model.ExtIdpDetailResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// CreateExtIdp
// 创建身份源
func (c *Client) CreateExtIdp(request *model.CreateExtIdpRequest) (*model.ExtIdpDetailResponse, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/extIdp", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
	if err != nil {
		return nil, err
	}
	resp := &struct {
		Message string                     `json:"message"`
		Code    int64                      `json:"code"`
		Data    model.ExtIdpDetailResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// UpdateExtIdp
// 更新身份源
func (c *Client) UpdateExtIdp(extIdpId string, request *model.UpdateExtIdpRequest) (*string, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/extIdp/%v", c.Host, extIdpId)
	b, err := c.SendHttpRestRequest(url, http.MethodPut, variables)

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

// DeleteExtIdp
// 删除身份源
func (c *Client) DeleteExtIdp(extIdpId string) (*string, error) {

	url := fmt.Sprintf("%s/api/v2/extIdp/%v", c.Host, extIdpId)
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

// CreateExtIdpConnection
// 创建身份源连接
func (c *Client) CreateExtIdpConnection(request *model.CreateExtIdpConnectionRequest) (*model.ExtIdpConnectionDetails, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/extIdpConn", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)

	if err != nil {
		return nil, err
	}

	resp := &struct {
		Message string                        `json:"message"`
		Code    int64                         `json:"code"`
		Data    model.ExtIdpConnectionDetails `json:"data"`
	}{}

	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// UpdateExtIdpConnection
// 更新身份源连接
func (c *Client) UpdateExtIdpConnection(extIdpConnectionId string, request *model.UpdateExtIdpConnectionRequest) (*string, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/extIdpConn/%v", c.Host, extIdpConnectionId)
	b, err := c.SendHttpRestRequest(url, http.MethodPut, variables)

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

// DeleteExtIdpConnection
// 删除身份源连接
func (c *Client) DeleteExtIdpConnection(extIdpConnectionId string) (*string, error) {

	url := fmt.Sprintf("%s/api/v2/extIdpConn/%v", c.Host, extIdpConnectionId)
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

// CheckExtIdpConnectionIdentifierUnique
// 检查连接唯一标识是否冲突
func (c *Client) CheckExtIdpConnectionIdentifierUnique(identifier string) (bool, error) {

	url := fmt.Sprintf("%s/api/v2/check/extIdpConn/identifier", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, map[string]interface{}{
		"identifier": identifier,
	})

	if err != nil {
		return true, err
	}

	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
	}{}

	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return true, errors.New(resp.Message)
	}
	return false, nil
}

// ChangeExtIdpConnectionState
// 开关身份源连接
func (c *Client) ChangeExtIdpConnectionState(extIdpConnectionId string, request *model.ChangeExtIdpConnectionStateRequest) (*string, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/extIdpConn/%v/state", c.Host, extIdpConnectionId)
	b, err := c.SendHttpRestRequest(url, http.MethodPut, variables)

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

// BatchChangeExtIdpConnectionState
// 批量开关身份源连接
func (c *Client) BatchChangeExtIdpConnectionState(extIdpId string, request *model.ChangeExtIdpConnectionStateRequest) (*string, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/extIdp/%v/connState", c.Host, extIdpId)
	b, err := c.SendHttpRestRequest(url, http.MethodPut, variables)

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
