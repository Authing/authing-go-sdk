package management

import (
	"errors"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

// CreateGroups
// 创建分组
func (c *Client) CreateGroups(req *model.CreateGroupsRequest) (*model.GroupModel, error) {
	data, _ := jsoniter.Marshal(req)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.CreateGroupsDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			CreateGroup model.GroupModel `json:"createGroup"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.CreateGroup, nil
}

// UpdateGroups
// 修改分组
func (c *Client) UpdateGroups(req *model.UpdateGroupsRequest) (*model.GroupModel, error) {
	data, _ := jsoniter.Marshal(req)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UpdateGroupsDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			UpdateGroup model.GroupModel `json:"updateGroup"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.UpdateGroup, nil
}

// DetailGroups
// 获取分组详情
func (c *Client) DetailGroups(code string) (*model.GroupModel, error) {

	variables := map[string]interface{}{"code": code}

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.DetailGroupsDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Group model.GroupModel `json:"group"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Group, nil
}

// DeleteGroups
// 删除分组
func (c *Client) DeleteGroups(code string) (*model.CommonMessageAndCode, error) {
	r, e := c.BatchDeleteGroups([]string{code})
	return r, e
}

// ListGroups
// 获取分组列表
func (c *Client) ListGroups(page, limit int) (*struct {
	TotalCount int64              `json:"totalCount"`
	List       []model.GroupModel `json:"list"`
}, error) {

	variables := map[string]interface{}{"page": page, "limit": limit}

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListGroupsDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Groups struct {
				TotalCount int64              `json:"totalCount"`
				List       []model.GroupModel `json:"list"`
			} `json:"groups"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Groups, nil
}

// BatchDeleteGroups
// 批量删除分组
func (c *Client) BatchDeleteGroups(codes []string) (*model.CommonMessageAndCode, error) {
	variables := map[string]interface{}{"codeList": codes}

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.DeleteGroupsDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			DeleteGroups model.CommonMessageAndCode `json:"deleteGroups"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.DeleteGroups, nil
}

// ListGroupsUser
// 获取分组用户列表
func (c *Client) ListGroupsUser(code string, page, limit int, withCustomData bool) (*struct {
	TotalCount int          `json:"totalCount"`
	List       []model.User `json:"list"`
}, error) {
	variables := map[string]interface{}{
		"code":  code,
		"page":  page,
		"limit": limit,
	}
	query := constant.ListGroupUserDocument
	if withCustomData {
		query = constant.ListGroupUserWithCustomDocument
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, query, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Group model.GetGroupUserResponse `json:"group"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Group.Users, nil
}

// AddUserToGroups
// 添加用户
func (c *Client) AddUserToGroups(code string, userIds []string) (*model.CommonMessageAndCode, error) {
	variables := map[string]interface{}{
		"code":    code,
		"userIds": userIds,
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.AddUserToGroupDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			AddUserToGroup model.CommonMessageAndCode `json:"addUserToGroup"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.AddUserToGroup, nil
}

//RemoveGroupUsers
//移除用户
func (c *Client) RemoveGroupUsers(code string, userIds []string) (*model.CommonMessageAndCode, error) {

	variables := map[string]interface{}{
		"code":    code,
		"userIds": userIds,
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RemoveUserInGroupDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			RemoveUserFromGroup model.CommonMessageAndCode `json:"removeUserFromGroup"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.RemoveUserFromGroup, nil
}

//ListGroupsAuthorizedResources
//获取分组被授权的所有资源
func (c *Client) ListGroupsAuthorizedResources(req *model.ListGroupsAuthorizedResourcesRequest) (*struct {
	TotalCount int64                      `json:"totalCount"`
	List       []model.AuthorizedResource `json:"list"`
}, error) {
	data, _ := jsoniter.Marshal(req)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListGroupAuthorizedResourcesDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			Group struct {
				AuthorizedResources struct {
					TotalCount int64                      `json:"totalCount"`
					List       []model.AuthorizedResource `json:"list"`
				} `json:"authorizedResources"`
			} `json:"group"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.Group.AuthorizedResources, nil
}
