package management

import (
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_IsAllowed(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========判断某个用户是否对某个资源有某个操作权限==========")

	req := model.IsAllowedRequest{
		Resource:  "7629:read",
		Action:    "read",
		UserId:    "611b2ff477d701441c25e29e",
		Namespace: nil,
	}
	resp, _ := client.IsAllowed(req)
	log.Printf("%+v\n", resp)
}

func TestClient_Allow(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========允许某个用户对某个资源进行某个操作==========")
	req := model.AllowRequest{
		Resource:  "7629:read",
		Action:    "add",
		UserId:    "611b2ff477d701441c25e29e",
		Namespace: "6123528118b7794b2420b311",
	}
	resp, _ := client.Allow(req)
	log.Printf("%+v\n", resp)
}

func TestClient_AuthorizeResource(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========将一个（类）资源批量授权给用户、角色、分组、组织机构==========")
	var actions []string
	actions = append(actions, "*")
	opt := model.AuthorizeResourceOpt{
		TargetType:       model.EnumPolicyAssignmentTargetTypeUSER,
		TargetIdentifier: "611b2ff477d701441c25e29e",
		Actions:          actions,
	}
	var opts []model.AuthorizeResourceOpt
	opts = append(opts, opt)
	req := model.AuthorizeResourceRequest{
		Namespace:    "6123528118b7794b2420b311",
		Resource:     "7629:read",
		ResourceType: model.EnumResourceTypeBUTTON,
		Opts:         opts,
	}
	resp, _ := client.AuthorizeResource(req)
	log.Printf("%+v\n", resp)
}

func TestClient_RevokeResource(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========将一个（类）资源取消授权给用户、角色、分组、组织机构==========")
	var actions []string
	actions = append(actions, "*")
	opt := model.AuthorizeResourceOpt{
		TargetType:       model.EnumPolicyAssignmentTargetTypeUSER,
		TargetIdentifier: "61090ca34e01a3968d3e3b76",
		Actions:          actions,
	}
	var opts []model.AuthorizeResourceOpt
	opts = append(opts, opt)
	req := model.RevokeResourceRequest{
		Namespace:    "default",
		Resource:     "7629:read",
		ResourceType: model.EnumResourceTypeBUTTON,
		Opts:         opts,
	}
	resp, _ := client.RevokeResource(req)
	log.Printf("%+v\n", resp)
}

func TestClient_CheckResourcePermissionBatch(t *testing.T) {
	//client := NewClient(userPoolId, appSecret)
	client := NewClient(userPoolIdDev, appSecretDev, "http://localhost:3000")
	log.Println("==========获取用户对某些资源的权限==========")
	var resources []string
	resources = append(resources, "7629:read")
	req := model.CheckResourcePermissionBatchRequest{
		UserId:    "61436e13634d7bdc0fd7ce6e",
		Namespace: "default",
		Resources: resources,
	}
	resp, _ := client.CheckResourcePermissionBatch(req)
	log.Printf("%+v\n", resp)
}
