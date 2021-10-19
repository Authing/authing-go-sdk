package management

import (
	"github.com/Authing/authing-go-sdk/lib/enum"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
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

func TestClient_ListAuthorizedResources(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取用户被授权的所有资源列表==========")

	req := model.ListUserAuthorizedResourcesRequest{
		UserId:       "611b2ff477d701441c25e29e",
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
