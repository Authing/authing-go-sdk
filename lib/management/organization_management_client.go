package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/enum"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
)

// ExportAll
// 导出所有组织机构
func (c *Client) ExportAll() ([]model.OrgNode, error) {
	var q []model.OrgNode
	b, err := c.SendHttpRequest(c.Host+"/api/v2/orgs/export", constant.HttpMethodGet, "", nil)
	if err != nil {
		return q, err
	}
	var response model.ExportAllOrganizationResponse
	log.Println(string(b))
	err = jsoniter.Unmarshal(b, &response)
	if err != nil {
		log.Println(err)
	}
	return response.Data, nil
}

// ListMembers
// 获取节点成员
func (c *Client) ListMembers(req *model.ListMemberRequest) (*model.Node, error) {
	if req.SortBy == "" {
		req.SortBy = enum.SortByCreatedAtAsc
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}
	variables := map[string]interface{}{
		"id":                   req.NodeId,
		"limit":                req.Limit,
		"sortBy":               req.SortBy,
		"page":                 req.Page,
		"includeChildrenNodes": req.IncludeChildrenNodes,
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.NodeByIdWithMembersDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println("___" + string(b))
	var response model.NodeByIdResponse
	jsoniter.Unmarshal(b, &response)
	return &response.Data.NodeById, nil
}

// TODO
func (c *Client) GetOrganizationList(request model.QueryListRequest) (model.PaginatedOrgs, error) {
	var result model.PaginatedOrgs
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+"/api/v2/orgs/pagination", constant.HttpMethodGet, "", variables)
	if err != nil {
		return result, err
	}
	var response model.ListOrganizationResponse
	jsoniter.Unmarshal(b, &response)
	return response.Data, nil
}

// GetOrganizationById
// 获取组织机构详情
func (c *Client) GetOrganizationById(orgId string) (*model.Org, error) {
	variables := map[string]interface{}{
		"id": orgId,
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.OrgDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response model.GetOrganizationByIdResponse
	jsoniter.Unmarshal(b, &response)
	return &response.Data.Org, nil
}

// GetOrganizationChildren
// 获取子节点列表
func (c *Client) GetOrganizationChildren(nodeId string, depth int) (*[]model.Node, error) {
	var result *[]model.Node
	variables := map[string]interface{}{
		"nodeId": nodeId,
		"depth":  depth,
	}
	b, err := c.SendHttpRequest(c.Host+"/api/v2/orgs/children", constant.HttpMethodGet, "", variables)
	if err != nil {
		return result, err
	}
	log.Println(string(b))
	var response model.GetOrganizationChildrenResponse
	jsoniter.Unmarshal(b, &response)
	return &response.Data, nil
}

// CreateOrg
// 创建组织机构
func (c *Client) CreateOrg(req *model.CreateOrgRequest) (*model.OrgResponse, error) {
	data, _ := jsoniter.Marshal(req)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.CreateOrgDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			CreateOrg model.OrgResponse `json:"createOrg"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.CreateOrg, nil
}

// DeleteOrgById
// 删除组织机构
func (c *Client) DeleteOrgById(id string) (*model.CommonMessageAndCode, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.DeleteOrgDocument, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			DeleteOrg model.CommonMessageAndCode `json:"deleteOrg"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.DeleteOrg, nil
}

// ListOrg
// 获取用户池组织机构列表
func (c *Client) ListOrg(page, limit int) (*model.PaginatedOrgs, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListOrgDocument, map[string]interface{}{
		"page":  page,
		"limit": limit,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			Orgs model.PaginatedOrgs `json:"orgs"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Orgs, nil
}

// AddOrgNode
// 在组织机构中添加一个节点
func (c *Client) AddOrgNode(req *model.AddOrgNodeRequest) (*model.AddNodeOrg, error) {
	data, _ := jsoniter.Marshal(req)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.AddOrgNodeDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			AddNode model.AddNodeOrg `json:"addNode"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.AddNode, nil
}

// GetOrgNodeById
// 获取某个节点详情
func (c *Client) GetOrgNodeById(id string) (*model.OrgNodeChildStr, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.GetOrgNodeDocument, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			NodeById model.OrgNodeChildStr `json:"nodeById"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.NodeById, nil
}

// UpdateOrgNode
// 修改节点
func (c *Client) UpdateOrgNode(req *model.UpdateOrgNodeRequest) (*model.Node, error) {
	data, _ := jsoniter.Marshal(req)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UpdateOrgNodeDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			UpdateNode model.Node `json:"updateNode"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.UpdateNode, nil
}

// DeleteOrgNode
// 删除节点
func (c *Client) DeleteOrgNode(orgId, nodeId string) (*model.CommonMessageAndCode, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.DeleteOrgNodeDocument, map[string]interface{}{
		"orgId":  orgId,
		"nodeId": nodeId,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			DeleteNode model.CommonMessageAndCode `json:"deleteNode"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.DeleteNode, nil
}

// IsRootNode
// 判断是否为根节点
func (c *Client) IsRootNode(orgId, nodeId string) (*bool, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.IsRootNodeDocument, map[string]interface{}{
		"orgId":  orgId,
		"nodeId": nodeId,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			IsRootNode bool `json:"isRootNode"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.IsRootNode, nil
}

// MoveOrgNode
// 移动节点
func (c *Client) MoveOrgNode(orgId, nodeId, targetParentId string) (*model.AddNodeOrg, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.MoveNodeDocument, map[string]interface{}{
		"orgId":          orgId,
		"nodeId":         nodeId,
		"targetParentId": targetParentId,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			MoveNode model.AddNodeOrg `json:"moveNode"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.MoveNode, nil
}

// GetRootNode
// 获取根节点
func (c *Client) GetRootNode(orgId string) (*model.OrgNodeChildStr, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.GetRootNodeDocument, map[string]interface{}{
		"orgId": orgId,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			RootNode model.OrgNodeChildStr `json:"rootNode"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.RootNode, nil
}

// ImportNodeByJSON
// 通过 JSON 导入
func (c *Client) ImportNodeByJSON(jsonStr string) (*string, error) {

	url := fmt.Sprintf("%s/api/v2/orgs/import", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, map[string]interface{}{
		"filetype": "json",
		"file":     jsonStr,
	})
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

// AddMembers
// 节点添加成员
func (c *Client) AddMembers(nodeId string, userIds []string) (*model.Node, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.AddMembersDocument, map[string]interface{}{
		"nodeId":  nodeId,
		"userIds": userIds,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			AddMember model.Node `json:"addMember"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.AddMember, nil
}

// MoveNodeMembers
// 移动节点成员
func (c *Client) MoveNodeMembers(nodeId, targetNodeId string, userIds []string) (*model.CommonMessageAndCode, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.MoveNodeMembersDocument, map[string]interface{}{
		"userIds":      userIds,
		"targetNodeId": targetNodeId,
		"sourceNodeId": nodeId,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			MoveMembers model.CommonMessageAndCode `json:"moveMembers"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.MoveMembers, nil
}

// DeleteNodeMembers
// 删除节点成员
func (c *Client) DeleteNodeMembers(nodeId string, userIds []string) (*model.Node, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RemoveNodeMembersDocument, map[string]interface{}{
		"userIds": userIds,
		"nodeId":  nodeId,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			RemoveMembers model.Node `json:"removeMember"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.RemoveMembers, nil
}

// SetMainDepartment
// 设置用户主部门
func (c *Client) SetMainDepartment(departmentId, userId string) (*model.CommonMessageAndCode, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.SetUserMainDepartmentDocument, map[string]interface{}{
		"userId":       userId,
		"departmentId": departmentId,
	})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			SetMainDepartment model.CommonMessageAndCode `json:"setMainDepartment"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.SetMainDepartment, nil
}

// ExportByOrgId
// 导出某个组织机构
func (c *Client) ExportByOrgId(orgId string) (*model.OrgNode, error) {

	url := fmt.Sprintf("%s/api/v2/orgs/export?org_id=%s", c.Host, orgId)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string        `json:"message"`
		Code    int64         `json:"code"`
		Data    model.OrgNode `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}

// ListAuthorizedResourcesByNodeId
// 获取组织机构节点被授权的所有资源
func (c *Client) ListAuthorizedResourcesByNodeId(req *model.ListAuthorizedResourcesByIdRequest) (*struct {
	TotalCount int64                      `json:"totalCount"`
	List       []model.AuthorizedResource `json:"list"`
}, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListNodeByIdAuthorizedResourcesDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			NodeByCode struct {
				AuthorizedResources struct {
					TotalCount int64                      `json:"totalCount"`
					List       []model.AuthorizedResource `json:"list"`
				} `json:"authorizedResources"`
			} `json:"nodeByCode"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.NodeByCode.AuthorizedResources, nil
}

// ListAuthorizedResourcesByNodeCode
// 获取组织机构节点被授权的所有资源
func (c *Client) ListAuthorizedResourcesByNodeCode(req *model.ListAuthorizedResourcesByNodeCodeRequest) (*struct {
	TotalCount int64                      `json:"totalCount"`
	List       []model.AuthorizedResource `json:"list"`
}, error) {
	data, _ := json.Marshal(&req)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListNodeByIdAuthorizedResourcesDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			NodeById struct {
				AuthorizedResources struct {
					TotalCount int64                      `json:"totalCount"`
					List       []model.AuthorizedResource `json:"list"`
				} `json:"authorizedResources"`
			} `json:"nodeById"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.NodeById.AuthorizedResources, nil
}

// SearchNodes
// 搜索组织机构节点
func (c *Client) SearchNodes(keywords string) (*[]model.OrgNodeChildStr, error) {

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost,
		constant.SearchNodesDocument, map[string]interface{}{"keyword": keywords})
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response = &struct {
		Data struct {
			SearchNodes []model.OrgNodeChildStr `json:"searchNodes"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.SearchNodes, nil
}

//
//// StartSync
//// 组织机构同步
//func (c *Client) StartSync(providerType constant.ProviderTypeEnum, connectionId *string) (*interface{}, error) {
//
//	url:=fmt.Sprintf("%s/connections/enterprise/%s/start-sync",c.Host,providerType)
//	vars:=make(map[string]interface{})
//	if providerType == constant.AD {
//		url = fmt.Sprintf("%s/api/v2/ad/sync",c.Host)
//		vars["connectionId"]=connectionId
//	}
//	b, err := c.SendHttpRestRequest(url, http.MethodPost,  vars)
//	if err != nil {
//		return nil, err
//	}
//	log.Println(string(b))
//	resp :=&struct {
//		Message string `json:"message"`
//		Code    int64  `json:"code"`
//		Data interface{} `json:"data"`
//	}{}
//	jsoniter.Unmarshal(b, &resp)
//	if resp.Code != 200 {
//		return nil, errors.New(resp.Message)
//	}
//	return &resp.Data, nil
//}
