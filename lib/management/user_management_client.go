package management

import (
	"encoding/json"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"github.com/bitly/go-simplejson"
	jsoniter "github.com/json-iterator/go"
	"log"
)

func (c *Client) Detail(userId string) (*model.User, error) {
	b, err := c.SendHttpRequest(c.Host+"/api/v2/users/"+userId, constant.HttpMethodGet, "", nil)
	if err != nil {
		return nil, err
	}
	var userDetail model.UserDetailResponse
	jsoniter.Unmarshal(b, &userDetail)
	return &userDetail.Data, nil
}

func (c *Client) GetUserList(request model.QueryListRequest) (*model.PaginatedUsers, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.UsersDocument, variables)
	if err != nil {
		return nil, err
	}
	result := model.ListUserResponse{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data.Users, nil
}

func (c *Client) GetUserDepartments(request model.GetUserDepartmentsRequest) (*model.PaginatedDepartments, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.GetUserDepartmentsDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	result := model.GetUserDepartmentsResponse{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.User.Departments, nil
}

func (c *Client) CheckUserExists(request model.CheckUserExistsRequest) (bool, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+"/api/v2/users/is-user-exists", constant.HttpMethodGet, constant.StringEmpty, variables)
	result := model.CheckUserExistsResponse{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return false, err
	}
	return result.Data, err
}

func (c *Client) ListAuthorizedResources(request model.ListUserAuthorizedResourcesRequest) (*model.User, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.ListUserAuthorizedResourcesDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	result := model.User{}
	resultJson, err := simplejson.NewJson(b)
	byteUser, err := resultJson.Get("data").Get("user").MarshalJSON()
	err = json.Unmarshal(byteUser, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetUserRoleList(request model.GetUserRoleListRequest) (*model.PaginatedRoles, error) {
	variables := make(map[string]interface{}, 0)
	if request.Namespace != nil {
		variables["namespace"] = *request.Namespace
	}
	b, err := c.SendHttpRequest(c.Host+"/api/v2/users/"+request.UserId+"/roles", constant.HttpMethodGet, constant.StringEmpty, variables)
	log.Println(string(b))
	result := model.PaginatedRoles{}
	resultJson, err := simplejson.NewJson(b)
	byteUser, err := resultJson.Get("data").MarshalJSON()
	err = json.Unmarshal(byteUser, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (c *Client) GetUserGroupList(userId string) (*model.PaginatedGroups, error) {
	variables := make(map[string]interface{}, 0)
	variables["id"] = userId
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.GetUserGroupsDocument, variables)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	result := model.PaginatedGroups{}
	resultJson, err := simplejson.NewJson(b)
	byteUser, err := resultJson.Get("data").Get("user").Get("groups").MarshalJSON()
	err = json.Unmarshal(byteUser, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
