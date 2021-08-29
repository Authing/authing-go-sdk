package management

import (
	"encoding/json"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/enum"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"log"
)

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
