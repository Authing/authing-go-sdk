package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/enum"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_GetRoleList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取角色列表==========")
	req := model.GetRoleListRequest{
		Page:   1,
		Limit:  10,
		SortBy: enum.SortByCreatedAtAsc,
	}
	resp, _ := client.GetRoleList(req)
	log.Printf("%+v\n", resp)
}

func TestClient_GetRoleUserList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取角色列表==========")
	defaultNamespace := "default"
	req := model.GetRoleUserListRequest{
		Page:      1,
		Limit:     10,
		Code:      "develop",
		Namespace: &defaultNamespace,
	}
	resp, _ := client.GetRoleUserList(req)
	log.Printf("%+v\n", resp)
}

func TestClient_CreateRole(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建角色==========")
	req := model.CreateRoleRequest{
		Code: "develop123456",
	}
	resp, err := client.CreateRole(req)

	log.Printf("%+v\n %+v\n", resp, err)
}

func TestClient_DeleteRole(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除角色==========")
	req := model.DeleteRoleRequest{
		Code: "develop123456",
	}
	resp, err := client.DeleteRole(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteRoles(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========批量删除角色==========")

	req := model.BatchDeleteRoleRequest{
		CodeList: []string{"develop123456", "develop1234562"},
	}
	resp, err := client.BatchDeleteRole(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_UpdateRole(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========更新角色==========")

	req := model.CreateRoleRequest{
		Code: "ttCode",
	}
	resp, err := client.CreateRole(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)

	updateRequest := model.UpdateRoleRequest{
		Code: "ttCode",
	}
	resp, err = client.UpdateRole(updateRequest)
	log.Printf("%+v\n", resp)
}

func TestClient_RoleDetail(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========角色详情==========")

	req := model.RoleDetailRequest{
		Code: "NewCode",
	}
	resp, err := client.RoleDetail(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_AssignRole(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========分配角色==========")

	req := model.AssignAndRevokeRoleRequest{
		RoleCodes: []string{"NewCode"},
		UserIds:   []string{"615551a3dcdd486139a917b1"},
	}
	resp, err := client.AssignRole(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_RevokeRole(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========撤回角色==========")

	req := model.AssignAndRevokeRoleRequest{
		RoleCodes: []string{"NewCode"},
		UserIds:   []string{"615551a3dcdd486139a917b1"},
	}
	resp, err := client.RevokeRole(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListRolePolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查询角色策略==========")

	req := model.ListPoliciesRequest{
		Code: "NewCode",
		//Code: "rndyxyjuan",
	}
	resp, err := client.ListRolePolicies(req)
	//resp, err := client.ListRolePolicies("rndyxyjuan", 1, 10)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_AddRolePolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查询角色策略==========")
	resp, err := client.AddRolePolicies("develop1234", []string{"ehsncbahxr"})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_RemoveRolePolicies(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查询角色策略==========")
	resp, err := client.RemoveRolePolicies("develop1234", []string{"ehsncbahxr"})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListRoleAuthorizedResources(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查询角色被授权资源==========")
	resp, err := client.ListRoleAuthorizedResources("NewCode", "default", model.EnumResourceTypeAPI)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_GetRoleUdfValue(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查询角色自定义字段==========")
	resp, err := client.GetRoleSpecificUdfValue("61692d23d17aec55f4cfcfa6")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println((*resp)[0].Key)
	log.Printf("%+v\n", resp)
}

func TestClient_BatchGetRoleUdfValue(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========批量查询角色自定义字段==========")
	resp, err := client.BatchGetRoleUdfValue([]string{"61692d23d17aec55f4cfcfa6", "61386f82e3a0b1c8a5bd7491"})
	if err != nil {
		fmt.Println(err)
	}
	d := resp["61692d23d17aec55f4cfcfa6"]
	fmt.Println(d[0].Key)
	log.Printf("%+v\n", resp)

}

func TestClient_SetRoleUdfValue(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========设置角色自定义字段==========")
	kv := &model.KeyValuePair{
		Key:   "school",
		Value: "西财",
	}
	resp, err := client.SetRoleUdfValue("624298162086c052b6dc8e5f", kv)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_BatchSetRoleUdfValue(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========批量设置角色自定义字段==========")
	f := &model.SetUdfValueBatchInput{
		Key:      "lhucskosfr",
		Value:    "123",
		TargetId: "616d112b7e387494d1ed0676",
	}
	tc := &model.SetUdfValueBatchInput{
		Key:      "lhucskosfr",
		Value:    "1235",
		TargetId: "61692d23d17aec55f4cfcfa6",
	}
	param := []model.SetUdfValueBatchInput{*f, *tc}
	resp, err := client.BatchSetRoleUdfValue(&param)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_RemoveRoleUdfValue(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除角色自定义字段==========")
	resp, err := client.RemoveRoleUdfValue("61692d23d17aec55f4cfcfa6", "lhucskosfr")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
