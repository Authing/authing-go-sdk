package management

import (
	"errors"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

// CreatePolicy
// 添加策略
func (c *Client) CreatePolicy(req *model.PolicyRequest) (*model.CreatePolicyResponse, error) {
	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.CreatePolicyDocument, vars)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			CreatePolicy model.CreatePolicyResponse `json:"createPolicy"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.CreatePolicy, nil
}

// ListPolicy
// 获取策略列表
func (c *Client) ListPolicy(page, limit int) (*model.PaginatedPolicies, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListPolicyDocument,
		map[string]interface{}{"page": page, "limit": limit})
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Policies model.PaginatedPolicies `json:"policies"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Policies, nil
}

// DetailPolicy
// 获取策略详情
func (c *Client) DetailPolicy(code string) (*model.Policy, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost,
		constant.DetailPolicyDocument, map[string]interface{}{"code": code})
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Policy model.Policy `json:"policy"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Policy, nil
}

// UpdatePolicy
// 修改策略
func (c *Client) UpdatePolicy(req *model.PolicyRequest) (*model.UpdatePolicyResponse, error) {
	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UpdatePolicyDocument, vars)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			UpdatePolicy model.UpdatePolicyResponse `json:"updatePolicy"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.UpdatePolicy, nil
}

// DeletePolicy
// 删除策略
func (c *Client) DeletePolicy(code string) (*model.CommonMessageAndCode, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost,
		constant.DeletePolicyDocument, map[string]interface{}{"code": code})
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			DeletePolicy model.CommonMessageAndCode `json:"deletePolicy"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.DeletePolicy, nil
}

// BatchDeletePolicy
// 删除策略
func (c *Client) BatchDeletePolicy(codeList []string) (*model.CommonMessageAndCode, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost,
		constant.BatchDeletePolicyDocument, map[string]interface{}{"codeList": codeList})
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			DeletePolicies model.CommonMessageAndCode `json:"deletePolicies"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.DeletePolicies, nil
}

// ListAssignments
// 获取策略授权记录
func (c *Client) ListAssignments(code string, page, limit int) (*model.PaginatedPolicyAssignments, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.PolicyAssignmentsDocument,
		map[string]interface{}{"code": code, "page": page, "limit": limit})
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			PolicyAssignments model.PaginatedPolicyAssignments `json:"policyAssignments"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.PolicyAssignments, nil
}

// AddAssignments
// 添加策略授权
func (c *Client) AddAssignments(req *model.PolicyAssignmentsRequest) (*model.CommonMessageAndCode, error) {
	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.AddAssignmentsDocument, vars)
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

// RemoveAssignments
// 撤销策略授权
func (c *Client) RemoveAssignments(req *model.PolicyAssignmentsRequest) (*model.CommonMessageAndCode, error) {
	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RemoveAssignmentsDocument, vars)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			RemovePolicyAssignments model.CommonMessageAndCode `json:"removePolicyAssignments"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.RemovePolicyAssignments, nil
}

// EnableAssignments
// 设置策略授权状态为开启
func (c *Client) EnableAssignments(req *model.SwitchPolicyAssignmentsRequest) (*model.CommonMessageAndCode, error) {
	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.EnablePolicyAssignmentDocument, vars)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			EnablePolicyAssignment model.CommonMessageAndCode `json:"enablePolicyAssignment"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.EnablePolicyAssignment, nil
}

// DisableAssignments
// 设置策略授权状态为关闭
func (c *Client) DisableAssignments(req *model.SwitchPolicyAssignmentsRequest) (*model.CommonMessageAndCode, error) {
	data, _ := jsoniter.Marshal(req)
	vars := make(map[string]interface{})
	jsoniter.Unmarshal(data, &vars)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.DisablePolicyAssignmentDocument, vars)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			DisablePolicyAssignment model.CommonMessageAndCode `json:"disbalePolicyAssignment"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.DisablePolicyAssignment, nil
}
