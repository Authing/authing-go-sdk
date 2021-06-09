package management

import (
	"authing-golang-sdk/lib/constant"
	"authing-golang-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
)

func (c *Client) ExportAll() ([]model.Node,error) {
	var q []model.Node
	b, err := c.SendHttpRequestV2(c.Host +"/api/v2/orgs/export", constant.HttpMethodGet, "", nil)
	if err != nil {
		return q, err
	}
	var response model.ExportAllOrganizationResponse
	jsoniter.Unmarshal(b, &response)
	return response.Data, nil
}

func (c *Client) ListMembers(req *model.ListMemberRequest) (*model.Node, error)  {
	variables := map[string]interface{}{
		"id":     req.NodeId,
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.NodeByIdWithMembersDocument, variables)
	if err != nil {
		return nil, err
	}
	var response model.NodeByIdResponse
	jsoniter.Unmarshal(b, &response)
	return &response.Data.NodeById, nil
}