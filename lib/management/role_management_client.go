package management

import (
	"encoding/json"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"log"
)

func (c *Client) GetRoleList(request model.GetRoleListRequest) (*model.PaginatedRoles, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.RolesDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response model.GetRoleListResponse
	jsoniter.Unmarshal(b, &response)
	return &response.Data.Roles, nil
}

func (c *Client) GetRoleUserList(request model.GetRoleUserListRequest) (*model.PaginatedRoles, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.RoleWithUsersDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var response model.GetRoleListResponse
	jsoniter.Unmarshal(b, &response)
	return &response.Data.Roles, nil
}
