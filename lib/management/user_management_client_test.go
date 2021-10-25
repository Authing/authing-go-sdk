package management

import (
	"github.com/Authing/authing-go-sdk/lib/enum"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"reflect"
	"testing"
)

func TestClient_GetUserList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========导出所有组织机构数据==========")
	custom := true
	req := model.QueryListRequest{
		Page:           1,
		Limit:          10,
		SortBy:         enum.SortByCreatedAtAsc,
		WithCustomData: &custom,
	}
	resp, _ := client.GetUserList(req)
	log.Printf("%+v\n", resp)
	log.Println(*resp)
}

func TestClient_GetUserDepartments(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取用户部门列表==========")
	req := model.GetUserDepartmentsRequest{
		Id:    "60e400c1701ea5b98dae628d",
		OrgId: nil,
	}
	resp, _ := client.GetUserDepartments(req)
	log.Printf("%+v\n", resp)
}

func TestClient_CheckUserExists(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========检查用户是否存在==========")
	//email := "t041gyqw0b@gmail.com"
	phone := "15761403457"
	req := model.CheckUserExistsRequest{
		Email:      nil,
		Phone:      &phone,
		Username:   nil,
		ExternalId: nil,
	}
	resp, _ := client.CheckUserExists(req)
	log.Println(resp)
}

func TestClient_CreateUser(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建用户==========")
	//email := "t041gyqw0b@gmail.com"
	phone := "15761403457222"
	username := "xx"
	pwd := "123456789"
	var userInfo = &model.CreateUserInput{
		Username: &username,
		Phone:    &phone,
		Password: &pwd,
	}
	req := model.CreateUserRequest{
		UserInfo: *userInfo,
	}
	resp, err := client.CreateUser(req)
	log.Println(resp)
	log.Println(err)
}

func TestClient_CreateUserWithCustom(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建用户包含自定义数据==========")
	//email := "t041gyqw0b@gmail.com"
	phone := "15761403457222122"
	username := "xxqq12"
	pwd := "123456789"
	var userInfo = &model.CreateUserInput{
		Username: &username,
		Phone:    &phone,
		Password: &pwd,
	}
	req := model.CreateUserRequest{
		UserInfo: *userInfo,
		CustomData: []model.KeyValuePair{
			model.KeyValuePair{
				Key:   "objhvfwdbi",
				Value: "qq",
			},
		},
	}
	resp, err := client.CreateUser(req)
	log.Println(resp)
	log.Println(err)
}

func TestClient_UpdateUser(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========更新用户==========")
	//email := "t041gyqw0b@gmail.com"
	phone := "15761403457222122"
	username := "xxqq123"
	//pwd:="123456789"
	var userInfo = &model.UpdateUserInput{
		Username: &username,
		Phone:    &phone,
		//Password: &pwd,
	}

	resp, err := client.UpdateUser("616d4333b809f9f4768db847", *userInfo)
	log.Println(resp)
	log.Println(err)
}

func TestClient_DeleteUser(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除用户==========")

	resp, err := client.DeleteUser("616d57e96dfa54908eda326f")
	log.Println(resp)
	log.Println(err)
}

func TestClient_BatchDeleteUser(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========批量删除用户==========")

	resp, err := client.BatchDeleteUser([]string{"616d430d58dbf82d1364453e"})
	log.Println(resp)
	log.Println(err)
}

func TestClient_BatchGetUser(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========批量获取用户==========")

	resp, err := client.BatchGetUser([]string{"xxq", "xx"}, "username", true)
	log.Println(resp)
	log.Println(err)
}

func TestClient_ListArchivedUsers(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取归档用户==========")

	resp, err := client.ListArchivedUsers(model.CommonPageRequest{
		Page:  1,
		Limit: 10,
	})
	log.Println(resp)
	log.Println(err)
}

func TestClient_FindUser(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查找用户==========")
	userName := "xxqq"
	resp, err := client.FindUser(&model.FindUserRequest{
		Username: &userName,
	})
	log.Println(resp)
	log.Println(err)
}

func TestClient_SearchUser(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查找用户==========")

	resp, err := client.SearchUser(&model.SearchUserRequest{
		Query: "xxqq",
	})
	log.Println(resp)
	log.Println(err)
}

func TestClient_UpdateUser2(t *testing.T) {
	username := "111"
	phone := "222"
	var userInfo = &model.UpdateUserInput{
		Username: &username,
		Phone:    &phone,
		//Password: &pwd,
	}
	u := "U"
	var defVal *string
	defVal = &u
	target := reflect.ValueOf(*userInfo)
	rUsername := target.FieldByName("Username")
	rIsVal := target.FieldByName("Gender")

	defaultVal := reflect.ValueOf(&defVal)
	log.Println(defaultVal.CanAddr())
	rIsVal.Set(defaultVal)

	log.Println(rUsername, rIsVal)
	log.Println(*userInfo.Gender)
}

func TestClient_RefreshUserToken(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========刷新用户Token==========")

	resp, err := client.RefreshUserToken("616d41b7410a33da0cb70e65")
	log.Println(*resp)
	log.Println(err)
}

func TestClient_GetUserGroups(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取用户所属组==========")

	resp, err := client.GetUserGroups("616d41b7410a33da0cb70e65")
	log.Println(resp)

	for k, v := range resp.List {
		log.Println(k)
		log.Println(v)
	}
	log.Println(err)
}

func TestClient_AddUserToGroup(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========添加用户到组==========")

	resp, err := client.AddUserToGroup("616d41b7410a33da0cb70e65", "pngrn")
	log.Println(resp)
	log.Println(err)
}

func TestClient_RemoveUserInGroup(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========移除组内用户==========")

	resp, err := client.RemoveUserInGroup("616d41b7410a33da0cb70e65", "pngrn")
	log.Println(resp)
	log.Println(err)
}

func TestClient_AddUserToRoles(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========用户分配角色==========")
	request := &model.UserRoleOptRequest{
		UserIds:   []string{"616d41b7410a33da0cb70e65"},
		RoleCodes: []string{"wwqhd"},
	}
	resp, err := client.AddUserToRoles(*request)
	log.Println(resp)
	log.Println(err)
}

func TestClient_GetUserRoles(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查询用户角色==========")
	request := &model.GetUserRolesRequest{
		Id: "616d41b7410a33da0cb70e65",
	}
	resp, err := client.GetUserRoles(*request)
	log.Println(resp)
	log.Println(err)
}

func TestClient_RemoveUserInRoles(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========撤回用户角色==========")
	request := &model.UserRoleOptRequest{
		UserIds:   []string{"616d41b7410a33da0cb70e65"},
		RoleCodes: []string{"wwqhd"},
	}
	resp, err := client.RemoveUserInRoles(*request)
	log.Println(resp)
	log.Println(err)
}

func TestClient_ListUserOrg(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查询用户机构==========")
	resp, err := client.ListUserOrg("616d41b7410a33da0cb70e65")
	log.Println(resp)
	log.Println(err)
}

func TestClient_GetUserUdfValue(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查询用户自定义字段==========")
	resp, err := client.GetUserUdfValue("616d41b7410a33da0cb70e65")
	log.Println(resp)
	log.Println(err)
}

func TestClient_ListUserAuthorizedResources(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查询用户授权资源==========")

	req := &model.ListUserAuthResourceRequest{
		Id:           "616d41b7410a33da0cb70e65",
		Namespace:    "default",
		ResourceType: model.EnumResourceTypeAPI,
	}
	resp, err := client.ListUserAuthorizedResources(*req)
	log.Println(resp)
	log.Println(err)
}

func TestClient_BatchGetUserUdfValue(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========批量查询用户自定义字段==========")
	resp, err := client.BatchGetUserUdfValue([]string{"616d41b7410a33da0cb70e65"})
	log.Println(resp)
	log.Println(err)
}

func TestClient_SetUserUdfValue(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========批量查询用户自定义字段==========")
	udv := model.KeyValuePair{
		Key:   "school",
		Value: "1x1",
	}
	resp, err := client.SetUserUdfValue("616d41b7410a33da0cb70e65", &udv)
	log.Println(resp)
	log.Println(err)
}

func TestClient_AddUserPolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========授权用户策略==========")

	resp, err := client.AddUserPolicies("616d41b7410a33da0cb70e65", []string{"ehsncbahxr"})
	log.Println(resp)
	log.Println(err)
}

func TestClient_ListUserPolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========遍历用户策略==========")
	req := model.ListPoliciesOnIdRequest{
		Id: "616d41b7410a33da0cb70e65",
	}
	resp, err := client.ListUserPolicies(req)
	log.Println(resp)
	log.Println(err)
}

func TestClient_RemoveUserPolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========撤销用户策略==========")

	resp, err := client.RemoveUserPolicies("616d41b7410a33da0cb70e65", []string{"ehsncbahxr"})
	log.Println(resp)
	log.Println(err)
}

func TestClient_UserHasRole(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========判断用户是否具有角色==========")

	resp, err := client.UserHasRole("616d41b7410a33da0cb70e65", "NewCode", "default")
	log.Println(resp)
	log.Println(err)
}

func TestClient_KickUser(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========强制用户下线==========")

	resp, err := client.KickUser([]string{"5a597f35085a2000144a10ed"})
	log.Println(resp)
	log.Println(err)
}

func TestClient_ListAuthorizedResources(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取用户被授权的所有资源列表==========")

	req := model.ListAuthorizedResourcesByIdRequest{
		Id:           "611b2ff477d701441c25e29e",
		Namespace:    "6123528118b7794b2420b311",
		ResourceType: nil,
	}
	resp, _ := client.ListAuthorizedResources(req)
	log.Printf("%+v\n", resp.AuthorizedResources)
}

func TestClient_GetUserRoleList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取用户角色列表==========")
	namespace := "default"
	req := model.GetUserRoleListRequest{
		UserId:    "611a149db64310ca4764ab15",
		Namespace: &namespace,
	}
	resp, _ := client.GetUserRoleList(req)
	log.Printf("%+v\n", resp)
}

func TestClient_GetUserGroupList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取用户分组列表==========")
	resp, _ := client.GetUserGroupList("611a149db64310ca4764ab15")
	log.Printf("%+v\n", resp)
}

func TestClient_CheckLoginStatus(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========检查用户登录状态==========")

	resp, err := client.CheckLoginStatus("5a597f35085a2000144a10ed", nil, nil)
	log.Println(resp)
	log.Println(err)
}

func TestClient_LogOut(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========用户退出==========")

	resp, err := client.LogOut("5a597f35085a2000144a10ed", nil)
	log.Println(resp)
	log.Println(err)
}

func TestClient_SendFirstLoginVerifyEmail(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========发送用户首次登录邮件==========")

	resp, err := client.SendFirstLoginVerifyEmail("616d4333b809f9f4768db847", "6168f95e81d5e20f9cb72f22")
	log.Println(resp)
	log.Println(err)
}

func TestClient_CheckLoginStatus2(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========检验登录状态根据Token==========")
	tx, e := GetAccessToken(client)
	log.Println(tx, e)
	resp, err := client.CheckLoginStatusByToken(tx)
	log.Println(resp)
	log.Println(err)
}

func TestClient_IsPasswordValid(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========检验登录状态根据Token==========")
	tx, e := GetAccessToken(client)
	log.Println(tx, e)
	resp, err := client.IsPasswordValid("tx")
	log.Println(resp)
	log.Println(err)
}
