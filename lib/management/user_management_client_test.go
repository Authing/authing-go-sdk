package management

import (
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/enum"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"reflect"
	"testing"
)

func TestClient_GetUserList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========导出所有组织机构数据==========")
	req := model.QueryListRequest{
		Page:   1,
		Limit:  10,
		SortBy: enum.SortByCreatedAtAsc,
	}
	resp, _ := client.GetUserList(req)
	log.Printf("%+v\n", resp)
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

//func TestClient_CreateUser(t *testing.T) {
//	client := NewClient(userPoolId, appSecret)
//	log.Println("==========创建用户==========")
//	//email := "t041gyqw0b@gmail.com"
//	phone :=  "15761403457222"
//	username := "xx"
//	pwd:="123456789"
//	var userInfo = &model.CreateUserInput{
//		Username: &username,
//		Phone: &phone,
//		Password: &pwd,
//	}
//	req := model.CreateUserRequest{
//		 UserInfo: *userInfo,
//	}
//	resp, err := client.CreateUser(req)
//	log.Println(resp)
//	log.Println(err)
//}

//func TestClient_CreateUserWithCustom(t *testing.T) {
//	client := NewClient(userPoolId, appSecret)
//	log.Println("==========创建用户包含自定义数据==========")
//	//email := "t041gyqw0b@gmail.com"
//	phone :=  "15761403457222122"
//	username := "xxqq12"
//	pwd:="123456789"
//	var userInfo = &model.CreateUserInput{
//		Username: &username,
//		Phone: &phone,
//		Password: &pwd,
//	}
//	req := model.CreateUserRequest{
//		UserInfo: *userInfo,
//		CustomData: []model.KeyValuePair{
//			 model.KeyValuePair{
//				Key: "objhvfwdbi",
//				Value: "qq",
//			},
//		},
//	}
//	resp, err := client.CreateUser(req)
//	log.Println(resp)
//	log.Println(err)
//}

//func TestClient_UpdateUser(t *testing.T) {
//	client := NewClient(userPoolId, appSecret)
//	log.Println("==========更新用户==========")
//	//email := "t041gyqw0b@gmail.com"
//	phone :=  "15761403457222122"
//	username := "xxqq123"
//	//pwd:="123456789"
//	var userInfo = &model.UpdateUserInput{
//		Username: &username,
//		Phone: &phone,
//		//Password: &pwd,
//	}
//
//	resp, err := client.UpdateUser("616d4333b809f9f4768db847",*userInfo)
//	log.Println(resp)
//	log.Println(err)
//}

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

	resp, err := client.FindUser(&model.FindUserRequest{
		Username: "xxqq",
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
		ResourceType: constant.API,
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

//func TestClient_KickUser(t *testing.T) {
//	client := NewClient(userPoolId, appSecret)
//	log.Println("==========强制用户下线==========")
//
//	resp, err := client.KickUser([]string{"5a597f35085a2000144a10ed"})
//	log.Println(resp)
//	log.Println(err)
//}
