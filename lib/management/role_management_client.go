package management

import (
	"encoding/json"
	"errors"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

// GetRoleList
// 获取角色列表
func (c *Client) GetRoleList(request model.GetRoleListRequest) (*model.PaginatedRoles, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.RolesDocument, variables)
	if err != nil {
		return nil, err
	}
	var response model.GetRoleListResponse
	jsoniter.Unmarshal(b, &response)
	return &response.Data.Roles, nil
}

// GetRoleUserList
// 获取角色用户列表
func (c *Client) GetRoleUserList(request model.GetRoleUserListRequest) (*struct {
	TotalCount int64        `json:"totalCount"`
	List       []model.User `json:"list"`
}, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.RoleWithUsersDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Role struct {
				Users struct {
					TotalCount int64        `json:"totalCount"`
					List       []model.User `json:"list"`
				} `json:"users"`
			} `json:"role"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	return &response.Data.Role.Users, nil
}

// CreateRole 创建角色
func (c *Client) CreateRole(request model.CreateRoleRequest) (*model.Role, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.CreateRoleDocument, variables)
	if err != nil {
		return nil, err
	}
	//var response model.CreateRoleResponse
	var response = &struct {
		Data struct {
			CreateRole model.Role `json:"createRole"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if response.Errors != nil {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.CreateRole, nil
}

// DeleteRole
// 删除角色
func (c *Client) DeleteRole(request model.DeleteRoleRequest) (*model.CommonMessageAndCode, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.DeleteRoleDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data   struct{ DeleteRole model.CommonMessageAndCode } `json:"data"`
		Errors []model.GqlCommonErrors                         `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if response.Errors != nil {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.DeleteRole, nil
}

// BatchDeleteRole
// 批量删除角色
func (c *Client) BatchDeleteRole(request model.BatchDeleteRoleRequest) (*model.CommonMessageAndCode, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.BatchDeleteRoleDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data   struct{ DeleteRoles model.CommonMessageAndCode } `json:"data"`
		Errors []model.GqlCommonErrors                          `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.DeleteRoles, nil
}

// RoleDetail
// 角色详情
func (c *Client) RoleDetail(request model.RoleDetailRequest) (*model.Role, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RoleDetailDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Role model.Role `json:"role"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Role, nil
}

// UpdateRole
// 更新角色
func (c *Client) UpdateRole(request model.UpdateRoleRequest) (*model.Role, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UpdateRoleDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data   struct{ UpdateRole model.Role } `json:"data"`
		Errors []model.GqlCommonErrors         `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.UpdateRole, nil
}

// AssignRole
// 角色 添加用户
func (c *Client) AssignRole(request model.AssignAndRevokeRoleRequest) (*model.CommonMessageAndCode, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.AssignRoleDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data   struct{ AssignRole model.CommonMessageAndCode } `json:"data"`
		Errors []model.GqlCommonErrors                         `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.AssignRole, nil
}

// RevokeRole
// 角色 移除用户
func (c *Client) RevokeRole(request model.AssignAndRevokeRoleRequest) (*model.CommonMessageAndCode, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RevokeRoleDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data   struct{ RevokeRole model.CommonMessageAndCode } `json:"data"`
		Errors []model.GqlCommonErrors                         `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.RevokeRole, nil
}

// ListRolePolicies
// 获取角色策略列表
func (c *Client) ListRolePolicies(request model.ListPoliciesRequest) (*model.ListPoliciesResponse, error) {

	if request.Page == 0 {
		request.Page = 1
	}
	if request.Limit == 0 {
		request.Limit = 10
	}

	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	variables["targetType"] = constant.ROLE
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListPoliciesDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			PolicyAssignments model.ListPoliciesResponse `json:"policyAssignments"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.PolicyAssignments, nil
}

// AddRolePolicies
// 给角色授权策略
func (c *Client) AddRolePolicies(code string, policiesCode []string) (*model.CommonMessageAndCode, error) {

	variables := make(map[string]interface{})

	variables["policies"] = policiesCode
	variables["targetType"] = constant.ROLE
	variables["targetIdentifiers"] = []string{code}

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.AddPoliciesDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			AddPolicyAssignments model.CommonMessageAndCode `json:"addPolicyAssignments"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.AddPolicyAssignments, nil
}

// RemoveRolePolicies
// 角色移除策略
func (c *Client) RemoveRolePolicies(code string, policiesCode []string) (*model.CommonMessageAndCode, error) {

	variables := make(map[string]interface{})

	variables["policies"] = policiesCode
	variables["targetType"] = constant.ROLE
	variables["targetIdentifiers"] = []string{code}

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RemovePoliciesDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			AddPolicyAssignments model.CommonMessageAndCode `json:"removePolicyAssignments"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.AddPolicyAssignments, nil
}

// ListRoleAuthorizedResources
// 获取角色被授权的所有资源
func (c *Client) ListRoleAuthorizedResources(code, namespace string, resourceType model.EnumResourceType) (*model.AuthorizedResources, error) {

	variables := make(map[string]interface{})

	variables["code"] = code
	variables["resourceType"] = resourceType
	variables["namespace"] = namespace

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListRoleAuthorizedResourcesDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Role struct {
				AuthorizedResources model.AuthorizedResources `json:"authorizedResources"`
			} `json:"role"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Role.AuthorizedResources, nil
}

// GetRoleUdfValue
// 获取某个角色扩展字段列表
func (c *Client) GetRoleUdfValue(id string) (*[]model.UserDefinedData, error) {

	variables := make(map[string]interface{})

	variables["targetType"] = constant.ROLE
	variables["targetId"] = id

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.GetRoleUdfValueDocument, variables)
	if err != nil {
		return nil, err
	}
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

// GetRoleSpecificUdfValue
// 获取某个角色某个扩展字段
func (c *Client) GetRoleSpecificUdfValue(id string) (*[]model.UserDefinedData, error) {
	return c.GetRoleUdfValue(id)
}

// BatchGetRoleUdfValue
// 获取多个角色扩展字段列表
func (c *Client) BatchGetRoleUdfValue(ids []string) (map[string][]model.UserDefinedData, error) {

	variables := make(map[string]interface{})

	variables["targetType"] = constant.ROLE
	variables["targetIds"] = ids

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.BatchGetRoleUdfValueDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			UdfValueBatch []model.BatchRoleUdv `json:"udfValueBatch"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	resultMap := make(map[string][]model.UserDefinedData)
	for _, v := range response.Data.UdfValueBatch {
		resultMap[v.TargetId] = v.Data
	}
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return resultMap, nil
}

// SetRoleUdfValue
// 设置某个角色扩展字段列表
func (c *Client) SetRoleUdfValue(id string, udv *model.KeyValuePair) (*[]model.UserDefinedData, error) {

	variables := make(map[string]interface{})

	v, _ := json.Marshal(udv.Value)
	udv.Value = string(v)
	variables["targetType"] = constant.ROLE
	variables["targetId"] = id
	variables["udvList"] = []model.KeyValuePair{*udv}

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.SetRoleUdfValueDocument, variables)
	if err != nil {
		return nil, err
	}
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

// BatchSetRoleUdfValue
// 设置多个角色扩展字段列表
func (c *Client) BatchSetRoleUdfValue(request *[]model.SetUdfValueBatchInput) (*model.CommonMessageAndCode, error) {

	variables := make(map[string]interface{})
	input := make([]model.SetUdfValueBatchInput, 0)
	for _, req := range *request {
		v, _ := json.Marshal(&req.Value)
		req.Value = string(v)
		input = append(input, req)
	}

	variables["targetType"] = constant.ROLE
	variables["input"] = input
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.BatchSetUdfValueDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			SetUdfValueBatch model.CommonMessageAndCode `json:"setUdfValueBatch"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.SetUdfValueBatch, nil
}

// RemoveRoleUdfValue
// 删除用户的扩展字段
func (c *Client) RemoveRoleUdfValue(id, key string) (*[]model.UserDefinedData, error) {

	variables := make(map[string]interface{})
	variables["targetType"] = constant.ROLE
	variables["targetId"] = id
	variables["key"] = key

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RemoveUdfValueDocument, variables)
	if err != nil {
		return nil, err
	}
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
