package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_CreateGroups(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建分组==========")
	req := &model.CreateGroupsRequest{
		Code: "goSDK",
		Name: "goSDK",
	}
	resp, err := client.CreateGroups(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_UpdateGroups(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========更新分组==========")
	newCode := "newGoSdk"
	req := &model.UpdateGroupsRequest{
		Code:    "goSDK",
		NewCode: &newCode,
	}
	resp, err := client.UpdateGroups(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DetailGroups(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========分组详情==========")

	resp, err := client.DetailGroups("newGoSdk")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteGroups(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除分组==========")

	resp, err := client.DeleteGroups("newGoSdk")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListGroups(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========分组列表==========")

	resp, err := client.ListGroups(1, 10)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListGroupsUser(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========分组下的用户列表==========")

	resp, err := client.ListGroupsUser("jjwjl", 1, 10, false)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListGroupsAuthorizedResources(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取分组授权资源==========")
	cc := model.EnumResourceTypeAPI
	nm := "default"
	req := &model.ListGroupsAuthorizedResourcesRequest{
		Code: "kcerb",
		//Code: "kmvnk",
		ResourceType: &cc,
		Namespace:    &nm,
	}
	resp, err := client.ListGroupsAuthorizedResources(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
