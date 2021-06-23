package management

import (
	"encoding/json"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"log"
)

func (c *Client) ExportAll() ([]model.Node, error) {
	var q []model.Node
	b, err := c.SendHttpRequest(c.Host+"/api/v2/orgs/export", constant.HttpMethodGet, "", nil)
	if err != nil {
		return q, err
	}
	var response model.ExportAllOrganizationResponse
	jsoniter.Unmarshal(b, &response)
	return response.Data, nil
}

func (c *Client) ListMembers(req *model.ListMemberRequest) (*model.Node, error) {
	variables := map[string]interface{}{
		"id": req.NodeId,
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.NodeByIdWithMembersDocument, variables)
	if err != nil {
		return nil, err
	}
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

func (c *Client) GetOrganizationFirstLevel(orgId string, depth int) (*[]model.Org, error) {
	var result *[]model.Org
	variables := map[string]interface{}{
		"orgId": orgId,
		"depth": depth,
	}
	b, err := c.SendHttpRequest(c.Host+"/api/v2/orgs/children", constant.HttpMethodGet, "", variables)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return result, err
	}
	log.Println(string(b))
	var response model.ListOrganizationResponse
	jsoniter.Unmarshal(b, &response)
	return result, nil
}
