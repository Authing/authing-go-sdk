package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"github.com/Authing/authing-go-sdk/lib/util"
	"github.com/bitly/go-simplejson"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
)

//IsAllowed
//判断某个用户是否对某个资源有某个操作权限
func (c *Client) IsAllowed(request model.IsAllowedRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.IsActionAllowedDocument, variables)
	if err != nil {
		return false, err
	}
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("data").Get("isActionAllowed").Bool()
	if err != nil {
		return false, err
	}
	return result, nil
}

//Allow
//允许某个用户对某个资源进行某个操作
func (c *Client) Allow(request model.AllowRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.AllowDocument, variables)
	if err != nil {
		return false, err
	}
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("data").Get("isActionAllowed").Bool()
	if err != nil {
		return false, err
	}
	return result, nil

}

//AuthorizeResource
//将一个（类）资源授权给用户、角色、分组、组织机构，且可以分别指定不同的操作权限。
func (c *Client) AuthorizeResource(request model.AuthorizeResourceRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.AuthorizeResourceDocument, variables)
	if err != nil {
		return false, err
	}
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("data").Get("authorizeResource").Get("code").Int64()
	if err != nil {
		return false, err
	}
	return result == 200, nil
}

//RevokeResource
//批量撤销资源的授权
func (c *Client) RevokeResource(request model.RevokeResourceRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+"/api/v2/acl/revoke-resource", constant.HttpMethodPost, constant.StringEmpty, variables)
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("code").Int64()
	if err != nil {
		return false, err
	}
	return result == 200, nil
}

// ListAuthorizedResourcesForCustom
// 获取某个主体（用户、角色、分组、组织机构节点）被授权的所有资源。
func (c *Client) ListAuthorizedResourcesForCustom(request model.ListAuthorizedResourcesRequest) (*struct {
	TotalCount int64                      `json:"totalCount"`
	List       []model.AuthorizedResource `json:"list"`
}, error) {

	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListAuthorizedResourcesDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			AuthorizedResources struct {
				TotalCount int64                      `json:"totalCount"`
				List       []model.AuthorizedResource `json:"list"`
			} `json:"authorizedResources"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.AuthorizedResources, nil
}

// ProgrammaticAccessAccountList
// 编程访问账号列表
func (c *Client) ProgrammaticAccessAccountList(appId string, page, limit int) (*struct {
	TotalCount int64                             `json:"totalCount"`
	List       []model.ProgrammaticAccessAccount `json:"list"`
}, error) {

	url := fmt.Sprintf("%s/api/v2/applications/%s/programmatic-access-accounts?limit=%v&page=%v", c.Host, appId, limit, page)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
		Data    struct {
			TotalCount int64                             `json:"totalCount"`
			List       []model.ProgrammaticAccessAccount `json:"list"`
		} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// CreateProgrammaticAccessAccount
// 添加编程访问账号
func (c *Client) CreateProgrammaticAccessAccount(appId string, remark *string, tokenLifetime *int) (*model.ProgrammaticAccessAccount, error) {

	vars := make(map[string]interface{})
	if tokenLifetime == nil {
		vars["tokenLifetime"] = 600
	} else {
		vars["tokenLifetime"] = tokenLifetime
	}
	if remark != nil {
		vars["remark"] = remark
	}
	url := fmt.Sprintf("%s/api/v2/applications/%s/programmatic-access-accounts", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                          `json:"message"`
		Code    int64                           `json:"code"`
		Data    model.ProgrammaticAccessAccount `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// DisableProgrammaticAccessAccount
// 禁用编程访问账号
func (c *Client) DisableProgrammaticAccessAccount(programmaticAccessAccountId string) (*model.ProgrammaticAccessAccount, error) {

	url := fmt.Sprintf("%s/api/v2/applications/programmatic-access-accounts", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPatch, map[string]interface{}{
		"id":      programmaticAccessAccountId,
		"enabled": false,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                          `json:"message"`
		Code    int64                           `json:"code"`
		Data    model.ProgrammaticAccessAccount `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// EnableProgrammaticAccessAccount
// 启用编程访问账号
func (c *Client) EnableProgrammaticAccessAccount(programmaticAccessAccountId string) (*model.ProgrammaticAccessAccount, error) {

	url := fmt.Sprintf("%s/api/v2/applications/programmatic-access-accounts", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPatch, map[string]interface{}{
		"id":      programmaticAccessAccountId,
		"enabled": true,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                          `json:"message"`
		Code    int64                           `json:"code"`
		Data    model.ProgrammaticAccessAccount `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// RefreshProgrammaticAccessAccountSecret
// 刷新编程访问账号密钥
func (c *Client) RefreshProgrammaticAccessAccountSecret(programmaticAccessAccountId string, secret *string) (*model.ProgrammaticAccessAccount, error) {

	vars := map[string]interface{}{
		"id": programmaticAccessAccountId,
	}
	if secret == nil {
		vars["secret"] = util.RandomString(32)
	} else {
		vars["secret"] = secret
	}
	url := fmt.Sprintf("%s/api/v2/applications/programmatic-access-accounts", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPatch, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                          `json:"message"`
		Code    int64                           `json:"code"`
		Data    model.ProgrammaticAccessAccount `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// DeleteProgrammaticAccessAccount
// 删除编程访问账号
func (c *Client) DeleteProgrammaticAccessAccount(programmaticAccessAccountId string) (*string, error) {

	url := fmt.Sprintf("%s/api/v2/applications/programmatic-access-accounts?id=%s", c.Host, programmaticAccessAccountId)
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

// ListNamespaceResources
// 获取资源列表
func (c *Client) ListNamespaceResources(req model.ListResourceRequest) (*model.ListNamespaceResourceResponse, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/resources", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                              `json:"message"`
		Code    int64                               `json:"code"`
		Data    model.ListNamespaceResourceResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// GetResourceById
// 根据 ID 获取单个资源
func (c *Client) GetResourceById(id string) (*model.ResourceResponse, error) {
	url := fmt.Sprintf("%s/api/v2/resources/detail", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                 `json:"message"`
		Code    int64                  `json:"code"`
		Data    model.ResourceResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// GetResourceByCode
// 根据 Code 获取单个资源
func (c *Client) GetResourceByCode(code, namespace string) (*model.ResourceResponse, error) {
	url := fmt.Sprintf("%s/api/v2/resources/detail", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, map[string]interface{}{
		"code":      code,
		"namespace": namespace,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                 `json:"message"`
		Code    int64                  `json:"code"`
		Data    model.ResourceResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// CreateResource
// 创建资源
func (c *Client) CreateResource(req *model.CreateResourceRequest) (*model.ResourceResponse, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	url := fmt.Sprintf("%s/api/v2/resources", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                 `json:"message"`
		Code    int64                  `json:"code"`
		Data    model.ResourceResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// UpdateResource
// 更新资源
func (c *Client) UpdateResource(code string, req *model.UpdateResourceRequest) (*model.ResourceResponse, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	url := fmt.Sprintf("%s/api/v2/resources/%s", c.Host, code)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                 `json:"message"`
		Code    int64                  `json:"code"`
		Data    model.ResourceResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// DeleteResource
// 删除资源
func (c *Client) DeleteResource(code, namespace string) (*string, error) {

	url := fmt.Sprintf("%s/api/v2/resources/%s", c.Host, code)
	b, err := c.SendHttpRestRequest(url, http.MethodDelete, map[string]interface{}{"namespace": namespace})
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

// GetApplicationAccessPolicies
// 获取应用访问控制策略列表
func (c *Client) GetApplicationAccessPolicies(appId string, page, limit int) (*model.GetApplicationAccessPoliciesResponse, error) {

	url := fmt.Sprintf("%s/api/v2/applications/%s/authorization/records", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, map[string]interface{}{
		"page":  page,
		"limit": limit,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                                     `json:"message"`
		Code    int64                                      `json:"code"`
		Data    model.GetApplicationAccessPoliciesResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// EnableApplicationAccessPolicies
// 启用应用访问控制策略
func (c *Client) EnableApplicationAccessPolicies(appId string, req *model.ApplicationAccessPoliciesRequest) (*string, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	url := fmt.Sprintf("%s/api/v2/applications/%s/authorization/enable-effect", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
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

// DisableApplicationAccessPolicies
// 停用应用访问控制策略
func (c *Client) DisableApplicationAccessPolicies(appId string, req *model.ApplicationAccessPoliciesRequest) (*string, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	url := fmt.Sprintf("%s/api/v2/applications/%s/authorization/disable-effect", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
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

// DeleteApplicationAccessPolicies
// 删除应用访问控制策略
func (c *Client) DeleteApplicationAccessPolicies(appId string, req *model.ApplicationAccessPoliciesRequest) (*string, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	url := fmt.Sprintf("%s/api/v2/applications/%s/authorization/revoke", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
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

// AllowApplicationAccessPolicies
// 配置「允许主体（用户、角色、分组、组织机构节点）访问应用」的控制策略
func (c *Client) AllowApplicationAccessPolicies(appId string, req *model.ApplicationAccessPoliciesRequest) (*string, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	url := fmt.Sprintf("%s/api/v2/applications/%s/authorization/allow", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
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

// DenyApplicationAccessPolicies
// 配置「拒绝主体（用户、角色、分组、组织机构节点）访问应用」的控制策略
func (c *Client) DenyApplicationAccessPolicies(appId string, req *model.ApplicationAccessPoliciesRequest) (*string, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	url := fmt.Sprintf("%s/api/v2/applications/%s/authorization/deny", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, variables)
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

// UpdateDefaultApplicationAccessPolicy
// 更改默认应用访问策略（默认拒绝所有用户访问应用、默认允许所有用户访问应用）
func (c *Client) UpdateDefaultApplicationAccessPolicy(appId string, strategy constant.ApplicationDefaultAccessPolicies) (*model.Application, error) {

	url := fmt.Sprintf("%s/api/v2/applications/%s", c.Host, appId)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, map[string]interface{}{
		"permissionStrategy": map[string]interface{}{
			"defaultStrategy": strategy,
		},
	})
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

// GetAuthorizedTargets
// 获取具备某些资源操作权限的主体
func (c *Client) GetAuthorizedTargets(req *model.GetAuthorizedTargetsRequest) (*struct {
	TotalCount int64 `json:"totalCount"`
	List       []struct {
		Actions          string `json:"actions"`
		TargetType       string `json:"targetType"`
		TargetIdentifier string `json:"targetIdentifier"`
	} `json:"list"`
}, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.GetAuthorizedTargetsDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			AuthorizedTargets struct {
				TotalCount int64 `json:"totalCount"`
				List       []struct {
					Actions          string `json:"actions"`
					TargetType       string `json:"targetType"`
					TargetIdentifier string `json:"targetIdentifier"`
				} `json:"list"`
			} `json:"authorizedTargets"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.AuthorizedTargets, nil
}

/*func (c *Client) CheckResourcePermissionBatch(request model.CheckResourcePermissionBatchRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+"/api/v2/acl/check-resource-permission-batch", constant.HttpMethodPost, constant.StringEmpty, variables)
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("code").Int64()
	if err != nil {
		return false, err
	}
	return result == 200, nil
}

func (c *Client) GetAuthorizedResourcesOfResourceKind(request model.GetAuthorizedResourcesOfResourceKindRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+"/api/v2/acl/get-authorized-resources-of-resource-kind", constant.HttpMethodPost, constant.StringEmpty, variables)
	log.Println(string(b))
	resultJson, err := simplejson.NewJson(b)
	result, err := resultJson.Get("code").Int64()
	if err != nil {
		return false, err
	}
	return result == 200, nil
}*/
