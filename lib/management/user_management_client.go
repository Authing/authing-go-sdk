package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"github.com/Authing/authing-go-sdk/lib/util"
	"github.com/bitly/go-simplejson"
	jsoniter "github.com/json-iterator/go"
)

// Detail
// 获取用户详情
func (c *Client) Detail(userId string) (*model.User, error) {
	b, err := c.SendHttpRequest(c.Host+"/api/v2/users/"+userId, constant.HttpMethodGet, "", nil)
	if err != nil {
		return nil, err
	}
	var userDetail model.UserDetailResponse
	jsoniter.Unmarshal(b, &userDetail)
	return &userDetail.Data, nil
}

// GetUserList
// 获取用户列表
func (c *Client) GetUserList(request model.QueryListRequest) (*model.PaginatedUsers, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	query := constant.UsersDocument
	if request.WithCustomData != nil && *request.WithCustomData == true {
		query = constant.UsersWithCustomDocument
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, query, variables)
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

// GetUserDepartments
// 获取用户所在部门
func (c *Client) GetUserDepartments(request model.GetUserDepartmentsRequest) (*model.PaginatedDepartments, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.GetUserDepartmentsDocument, variables)
	if err != nil {
		return nil, err
	}
	result := model.GetUserDepartmentsResponse{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.User.Departments, nil
}

// CheckUserExists
// 检查用户是否存在
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

// CreateUser
// 创建用户
func (c *Client) CreateUser(request model.CreateUserRequest) (*model.User, error) {
	if request.UserInfo.Password != nil {
		pwd := util.RsaEncrypt(*request.UserInfo.Password)
		request.UserInfo.Password = &pwd
	}
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	query := constant.CreateUserDocument
	if request.CustomData != nil {
		query = constant.CreateUserWithCustomDataDocument
		customData, _ := json.Marshal(&request.CustomData)
		variables["params"] = customData
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, query, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			CreateUser model.User `json:"createUser"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.CreateUser, nil
}

//UpdateUser
//修改用户资料
func (c *Client) UpdateUser(id string, updateInfo model.UpdateUserInput) (*model.User, error) {
	if updateInfo.Password != nil {
		pwd := util.RsaEncrypt(*updateInfo.Password)
		updateInfo.Password = &pwd
	}
	variables := make(map[string]interface{})
	variables["id"] = id
	variables["input"] = updateInfo
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.UpdateUserDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			UpdateUser model.User `json:"updateUser"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.UpdateUser, nil
}

//DeleteUser
//删除用户
func (c *Client) DeleteUser(id string) (*model.CommonMessageAndCode, error) {

	variables := make(map[string]interface{})
	variables["id"] = id

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.DeleteUserDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			DeleteUser model.CommonMessageAndCode `json:"deleteUser"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.DeleteUser, nil
}

//BatchDeleteUser
//批量删除用户
func (c *Client) BatchDeleteUser(ids []string) (*model.CommonMessageAndCode, error) {
	variables := make(map[string]interface{})
	variables["ids"] = ids
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.BatchDeleteUserDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			DeleteUsers model.CommonMessageAndCode `json:"deleteUsers"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.DeleteUsers, nil
}

//BatchGetUser
//通过 ID、username、email、phone、email、externalId 批量获取用户详情
func (c *Client) BatchGetUser(ids []string, queryField string, withCustomData bool) (*[]model.User, error) {

	variables := make(map[string]interface{})
	variables["ids"] = ids
	variables["type"] = queryField
	query := constant.BatchGetUserDocument
	if withCustomData {
		query = constant.BatchGetUserWithCustomDocument
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, query, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			BatchGetUsers []model.User `json:"userBatch"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.BatchGetUsers, nil
}

//ListArchivedUsers
//获取已归档用户列表
func (c *Client) ListArchivedUsers(request model.CommonPageRequest) (*model.CommonPageUsersResponse, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListArchivedUsersDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			ArchivedUsers model.CommonPageUsersResponse `json:"archivedUsers"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.ArchivedUsers, nil
}

//FindUser
//查找用户
func (c *Client) FindUser(request *model.FindUserRequest) (*model.User, error) {

	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	query := constant.FindUserDocument
	if request.WithCustomData {
		query = constant.FindUserWithCustomDocument
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, query, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			FindUser model.User `json:"findUser"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.FindUser, nil
}

//SearchUser
//搜索用户
func (c *Client) SearchUser(request *model.SearchUserRequest) (*model.CommonPageUsersResponse, error) {
	if request.Page == 0 {
		request.Page = 1
	}
	if request.Limit == 0 {
		request.Limit = 10
	}

	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	query := constant.SearchUserDocument
	if request.WithCustomData {
		query = constant.SearchUserWithCustomDocument
	}
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, query, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			SearchUser model.CommonPageUsersResponse `json:"searchUser"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.SearchUser, nil
}

//RefreshUserToken
//刷新用户 token
func (c *Client) RefreshUserToken(userId string) (*model.RefreshToken, error) {
	variables := make(map[string]interface{})
	variables["id"] = userId

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RefreshUserTokenDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			RefreshToken model.RefreshToken `json:"refreshToken"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.RefreshToken, nil
}

//GetUserGroups
//获取用户分组列表
func (c *Client) GetUserGroups(userId string) (*struct {
	TotalCount int                `json:"totalCount"`
	List       []model.GroupModel `json:"list"`
}, error) {
	variables := make(map[string]interface{})
	variables["id"] = userId

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.GetUserGroupsDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			User model.GetUserGroupsResponse `json:"user"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.User.Groups, nil
}

//AddUserToGroup
//将用户加入分组
func (c *Client) AddUserToGroup(userId, groupCode string) (*model.CommonMessageAndCode, error) {

	variables := make(map[string]interface{})
	variables["userIds"] = []string{userId}
	variables["code"] = groupCode
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

//RemoveUserInGroup
//将用户退出分组
func (c *Client) RemoveUserInGroup(userId, groupCode string) (*model.CommonMessageAndCode, error) {

	variables := make(map[string]interface{})
	variables["userIds"] = []string{userId}
	variables["code"] = groupCode
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

//GetUserRoles
//获取用户角色列表
func (c *Client) GetUserRoles(request model.GetUserRolesRequest) (*struct {
	TotalCount int               `json:"totalCount"`
	List       []model.RoleModel `json:"list"`
}, error) {

	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.GetUserRolesDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			User model.GetUserRolesResponse `json:"user"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.User.Roles, nil
}

//AddUserToRoles
//将用户加入角色
func (c *Client) AddUserToRoles(request model.UserRoleOptRequest) (*model.CommonMessageAndCode, error) {
	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.AddUserToRoleDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			AssignRole model.CommonMessageAndCode `json:"assignRole"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.AssignRole, nil
}

//RemoveUserInRoles
//将用户从角色中移除
func (c *Client) RemoveUserInRoles(request model.UserRoleOptRequest) (*model.CommonMessageAndCode, error) {
	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.RemoveUserInRoleDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			RevokeRole model.CommonMessageAndCode `json:"revokeRole"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}
	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.RevokeRole, nil
}

//ListUserOrg
//获取用户所在组织机构
func (c *Client) ListUserOrg(userId string) (*[][]model.OrgModel, error) {

	url := fmt.Sprintf("%v/api/v2/users/%v/orgs", c.Host, userId)
	b, err := c.SendHttpRequest(url, http.MethodGet, "", nil)
	if err != nil {
		return nil, err
	}

	var response [][]model.OrgModel
	var resultMap map[string]interface{}
	e := jsoniter.Unmarshal(b, &resultMap)

	if e != nil || resultMap["code"].(float64) != 200 {
		return nil, errors.New("ListUserOrg Error")
	}
	data, _ := jsoniter.Marshal(resultMap["data"])
	jsoniter.Unmarshal(data, &response)
	return &response, nil
}

//GetUserUdfValue
//获取某个用户的所有自定义数据
func (c *Client) GetUserUdfValue(userId string) (*[]model.UserDefinedData, error) {
	variables := make(map[string]interface{})

	variables["targetType"] = constant.USER
	variables["targetId"] = userId

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

// ListUserAuthorizedResources
// 获取用户被授权的所有资源
func (c *Client) ListUserAuthorizedResources(request model.ListUserAuthResourceRequest) (*model.AuthorizedResources, error) {

	data, _ := jsoniter.Marshal(request)
	variables := make(map[string]interface{})
	jsoniter.Unmarshal(data, &variables)

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.ListUserAuthorizedResourcesDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			User struct {
				AuthorizedResources model.AuthorizedResources `json:"authorizedResources"`
			} `json:"user"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.User.AuthorizedResources, nil
}

// BatchGetUserUdfValue
// 批量获取多个用户的自定义数据
func (c *Client) BatchGetUserUdfValue(ids []string) (map[string][]model.UserDefinedData, error) {

	variables := make(map[string]interface{})

	variables["targetType"] = constant.USER
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

// SetUserUdfValue
// 设置某个用户的自定义数据
func (c *Client) SetUserUdfValue(id string, udv *model.KeyValuePair) (*[]model.UserDefinedData, error) {

	variables := make(map[string]interface{})

	variables["targetType"] = constant.USER
	variables["targetId"] = id
	variables["key"] = udv.Key
	variables["value"] = udv.Value

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.SetUdvDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			SetUdvBatch []model.UserDefinedData `json:"setUdv"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)
	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.SetUdvBatch, nil
}

// BatchSetUserUdfValue
// 批量设置自定义数据
func (c *Client) BatchSetUserUdfValue(request *[]model.SetUdfValueBatchInput) (*model.CommonMessageAndCode, error) {

	variables := make(map[string]interface{})
	variables["targetType"] = constant.USER
	variables["input"] = request
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

// RemoveUserUdfValue
// 清除用户的自定义数据
func (c *Client) RemoveUserUdfValue(id, key string) (*[]model.UserDefinedData, error) {

	variables := make(map[string]interface{})
	variables["targetType"] = constant.USER
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

// ListUserPolicies
// 获取策略列表
func (c *Client) ListUserPolicies(request model.ListPoliciesOnIdRequest) (*model.ListPoliciesResponse, error) {

	if request.Page == 0 {
		request.Page = 1
	}
	if request.Limit == 0 {
		request.Limit = 10
	}

	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	variables["targetType"] = constant.USER
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

// AddUserPolicies
// 批量添加策略
func (c *Client) AddUserPolicies(userId string, policiesCode []string) (*model.CommonMessageAndCode, error) {

	variables := make(map[string]interface{})

	variables["policies"] = policiesCode
	variables["targetType"] = constant.USER
	variables["targetIdentifiers"] = []string{userId}

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

// RemoveUserPolicies
// 批量移除策略
func (c *Client) RemoveUserPolicies(userId string, policiesCode []string) (*model.CommonMessageAndCode, error) {

	variables := make(map[string]interface{})

	variables["policies"] = policiesCode
	variables["targetType"] = constant.USER
	variables["targetIdentifiers"] = []string{userId}

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

// UserHasRole
// 判断用户是否有某个角色
func (c *Client) UserHasRole(userId, roleCode, namespace string) (bool, error) {
	req := model.GetUserRolesRequest{
		Id:        userId,
		Namespace: namespace,
	}
	hasRole := false
	list, err := c.GetUserRoles(req)
	if err != nil {
		return false, err
	}
	if list.TotalCount == 0 {
		return false, nil
	}
	for _, v := range list.List {
		if v.Code == roleCode {
			hasRole = true
			break
		}
	}
	return hasRole, nil
}

//KickUser
//强制一批用户下线
func (c *Client) KickUser(userIds []string) (*model.CommonMessageAndCode, error) {

	url := fmt.Sprintf("%v/api/v2/users/kick", c.Host)
	json := make(map[string]interface{})
	json["userIds"] = userIds
	b, err := c.SendHttpRequest(url, http.MethodPost, "", json)
	if err != nil {
		return nil, err
	}
	var response model.CommonMessageAndCode
	jsoniter.Unmarshal(b, &response)
	return &response, nil
}

func (c *Client) ListAuthorizedResources(request model.ListAuthorizedResourcesByIdRequest) (*model.User, error) {
	data, _ := json.Marshal(&request)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, constant.HttpMethodPost, constant.ListUserAuthorizedResourcesDocument, variables)
	if err != nil {
		return nil, err
	}
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
	result := model.PaginatedGroups{}
	resultJson, err := simplejson.NewJson(b)
	byteUser, err := resultJson.Get("data").Get("user").Get("groups").MarshalJSON()
	err = json.Unmarshal(byteUser, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//CheckLoginStatus
//检查用户登录状态
func (c *Client) CheckLoginStatus(userId string, appId, deviceId *string) (*model.CommonMessageAndCode, error) {
	variables := make(map[string]interface{}, 0)
	if appId != nil {
		variables["appId"] = appId
	}
	if deviceId != nil {
		variables["deviceId"] = deviceId
	}
	variables["userId"] = userId

	url := fmt.Sprintf("%v/api/v2/users/login-status", c.Host)
	b, err := c.SendHttpRequest(url, constant.HttpMethodGet, constant.StringEmpty, variables)
	result := model.CommonMessageAndCode{}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

//LogOut
//用户退出
func (c *Client) LogOut(userId string, appId *string) (*model.CommonMessageAndCode, error) {
	variables := make(map[string]interface{}, 0)
	if appId != nil {
		variables["appId"] = appId
	}

	variables["userId"] = userId

	url := fmt.Sprintf("%v/logout", c.Host)
	b, err := c.SendHttpRequest(url, http.MethodGet, constant.StringEmpty, variables)
	result := model.CommonMessageAndCode{}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// SendFirstLoginVerifyEmail
// 发送首次登录验证邮件
func (c *Client) SendFirstLoginVerifyEmail(userId, appId string) (*model.CommonMessageAndCode, error) {

	variables := make(map[string]interface{})
	variables["appId"] = appId
	variables["userId"] = userId

	b, err := c.SendHttpRequest(c.Host+constant.CoreAuthingGraphqlPath, http.MethodPost, constant.SendFirstLoginVerifyEmailDocument, variables)
	if err != nil {
		return nil, err
	}
	var response = &struct {
		Data struct {
			SendFirstLoginVerifyEmail model.CommonMessageAndCode `json:"sendFirstLoginVerifyEmail"`
		} `json:"data"`
		Errors []model.GqlCommonErrors `json:"errors"`
	}{}

	jsoniter.Unmarshal(b, &response)

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Message.Message)
	}
	return &response.Data.SendFirstLoginVerifyEmail, nil
}

// GetUserTenants
// 获取用户所在租户
func (c *Client) GetUserTenants(userId string) (*model.GetUserTenantsResponse, error) {

	url := fmt.Sprintf("%s/api/v2/users/%v/tenants", c.Host, userId)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	resp := &struct {
		Message string                       `json:"message"`
		Code    int64                        `json:"code"`
		Data    model.GetUserTenantsResponse `json:"data"`
	}{}

	jsoniter.Unmarshal(b, &resp)

	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp.Data, nil
}
