package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_IsAllowed(t *testing.T) {
	client, err := NewClientWithError(userPoolId, appSecret)
	if err != nil {
		fmt.Println(err)
		return
	}
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
		TargetType:       model.EnumPolicyAssignmentTargetTypeUser,
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
		TargetType:       model.EnumPolicyAssignmentTargetTypeUser,
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

func TestClient_ListAuthorizedResourcesForCustom(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========将一个（类）资源取消授权给用户、角色、分组、组织机构==========")
	rt := model.EnumResourceTypeAPI
	req := model.ListAuthorizedResourcesRequest{
		Namespace:        "default",
		ResourceType:     &rt,
		TargetIdentifier: "616d41b7410a33da0cb70e65",
		TargetType:       constant.USER,
	}
	resp, _ := client.ListAuthorizedResourcesForCustom(req)
	log.Printf("%+v\n", resp)
}

func TestClient_ProgrammaticAccessAccountList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========编程账号列表==========")

	resp, _ := client.ProgrammaticAccessAccountList("6168f95e81d5e20f9cb72f22", 1, 10)
	log.Printf("%+v\n", resp)
}

func TestClient_CreateProgrammaticAccessAccount(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========编程账号列表==========")

	resp, _ := client.CreateProgrammaticAccessAccount("6168f95e81d5e20f9cb72f22", nil, nil)
	log.Printf("%+v\n", resp)
}

func TestClient_DisableProgrammaticAccessAccount(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========禁用编程账号==========")

	resp, _ := client.DisableProgrammaticAccessAccount("617109c03d185a5092395cab")
	log.Printf("%+v\n", resp)
}

func TestClient_EnableProgrammaticAccessAccount(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========启用编程账号==========")

	resp, _ := client.EnableProgrammaticAccessAccount("617109c03d185a5092395cab")
	log.Printf("%+v\n", resp)
}

func TestClient_RefreshProgrammaticAccessAccountSecret(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========刷新编程账号访问秘钥==========")

	resp, _ := client.RefreshProgrammaticAccessAccountSecret("617109c03d185a5092395cab", nil)
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteProgrammaticAccessAccount(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========刷新编程账号访问秘钥==========")

	resp, _ := client.DeleteProgrammaticAccessAccount("617109c03d185a5092395cab")
	log.Printf("%+v\n", resp)
}

func TestClient_ListNamespaceResources(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取Namespace下资源列表==========")

	req := model.ListResourceRequest{
		ResourceType: model.EnumResourceTypeAPI,
		Namespace:    "default",
		Page:         1,
		Limit:        10,
	}
	resp, _ := client.ListNamespaceResources(req)
	log.Printf("%+v\n", resp)
}

func TestClient_GetResourceById(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========根据ID获取资源==========")

	resp, _ := client.GetResourceById("616cdf9d1642b20d8c2ec555")
	log.Printf("%+v\n", resp)
}

func TestClient_GetResourceByCode(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========根据 Code 获取资源==========")

	resp, _ := client.GetResourceByCode("ddddd", "default")
	log.Printf("%+v\n", resp)
}

func TestClient_CreateResource(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建资源==========")
	req := &model.CreateResourceRequest{
		Code:      "nmw",
		Namespace: "default",
		Actions: []model.ActionsModel{{
			Name:        "qqw",
			Description: "qwe",
		}},
	}
	resp, _ := client.CreateResource(req)
	log.Printf("%+v\n", resp)
}

func TestClient_UpdateResource(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========修改资源==========")
	req := &model.UpdateResourceRequest{

		Namespace: "default",
		Actions: []model.ActionsModel{{
			Name:        "qqwcc",
			Description: "qwe",
		}},
	}
	resp, _ := client.UpdateResource("nmw", req)
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteResource(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除资源==========")

	resp, _ := client.DeleteResource("nmw", "default")
	log.Printf("%+v\n", resp)
}

func TestClient_GetApplicationAccessPolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取应用访问策略==========")

	resp, _ := client.GetApplicationAccessPolicies("6168f95e81d5e20f9cb72f22", 1, 10)
	log.Printf("%+v\n", resp)
}

func TestClient_EnableApplicationAccessPolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========启用应用访问策略==========")
	req := &model.ApplicationAccessPoliciesRequest{
		TargetType:        constant.USER,
		InheritByChildren: true,
		TargetIdentifiers: []string{"616e905ebc18f0f106973a29"},
		Namespace:         "default",
	}
	resp, _ := client.EnableApplicationAccessPolicies("6168f95e81d5e20f9cb72f22", req)
	log.Printf("%+v\n", resp)
}

func TestClient_DisableApplicationAccessPolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========禁用应用访问策略==========")
	req := &model.ApplicationAccessPoliciesRequest{
		TargetType:        constant.USER,
		InheritByChildren: true,
		TargetIdentifiers: []string{"616e905ebc18f0f106973a29"},
		Namespace:         "default",
	}
	resp, _ := client.DisableApplicationAccessPolicies("6168f95e81d5e20f9cb72f22", req)
	log.Printf("%+v\n", resp)
}

func TestClient_AllowApplicationAccessPolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========允许应用访问策略==========")
	req := &model.ApplicationAccessPoliciesRequest{
		TargetType:        constant.USER,
		InheritByChildren: true,
		TargetIdentifiers: []string{"616e905ebc18f0f106973a29"},
		Namespace:         "default",
	}
	resp, _ := client.AllowApplicationAccessPolicies("6168f95e81d5e20f9cb72f22", req)
	log.Printf("%+v\n", resp)
}

func TestClient_DenyApplicationAccessPolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========拒绝应用访问策略==========")
	req := &model.ApplicationAccessPoliciesRequest{
		TargetType:        constant.USER,
		InheritByChildren: true,
		TargetIdentifiers: []string{"616e905ebc18f0f106973a29"},
		Namespace:         "default",
	}
	resp, _ := client.DenyApplicationAccessPolicies("6168f95e81d5e20f9cb72f22", req)
	log.Printf("%+v\n", resp)
}

func TestClient_UpdateDefaultApplicationAccessPolicy(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========修改应用默认访问策略==========")

	resp, _ := client.UpdateDefaultApplicationAccessPolicy("6168f95e81d5e20f9cb72f22", constant.AllowAll)
	log.Printf("%+v\n", resp)
}

func TestClient_GetAuthorizedTargets(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取拥有资源的对象列表==========")

	req := &model.GetAuthorizedTargetsRequest{
		TargetType:   constant.ROLE,
		Resource:     "cccccc",
		Namespace:    "default",
		ResourceType: model.EnumResourceTypeAPI,
	}
	resp, _ := client.GetAuthorizedTargets(req)
	log.Printf("%+v\n", resp)
}

/*func TestClient_CheckResourcePermissionBatch(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
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
}*/
