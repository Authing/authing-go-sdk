package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/enum"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

const (
	// prod

	//userPoolId = ""
	//appSecret  = ""
	//userPoolId = ""
	//appSecret = ""
	userPoolId = ""
	appSecret  = ""
)

func TestClient_ExportAll(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========导出所有组织机构数据==========")
	resp, err := client.ExportAll()
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_All(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========导出所有组织机构数据==========")
	resp, _ := client.ExportAll()
	log.Printf("%+v\n", resp)
	log.Println("==========获取节点成员==========")
	var req = &model.ListMemberRequest{
		NodeId:               "60cd9d3a4b96cfff16e7e5f4",
		Page:                 1,
		Limit:                10,
		IncludeChildrenNodes: true,
	}
	resp1, _ := client.ListMembers(req)
	log.Printf("%+v\n", resp1)
	log.Println("==========通过 ID 获取用户信息==========")
	resp2, _ := client.Detail("60a6f9ad5bcccc51834950c5")
	log.Printf("%+v\n", resp2)
}
func TestClient_GetOrganizationList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取用户池组织机构列表==========")
	req := model.QueryListRequest{
		Page:   1,
		Limit:  10,
		SortBy: enum.SortByCreatedAtAsc,
	}
	resp, _ := client.GetOrganizationList(req)
	log.Printf("%+v\n", resp)
}

func TestClient_GetOrganizationById(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取组织机构详情==========")
	resp, _ := client.GetOrganizationById("60cd9d3ab98280ce211bc834")
	log.Printf("%+v\n", resp)
}

func TestClient_GetOrganizationChildren(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取第 n 层组织机构==========")
	resp, _ := client.GetOrganizationChildren("60cd9d3a4b96cfff16e7e5f4", 1)
	log.Printf("%+v\n", resp)
}

func TestClient_CreateOrg(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建组织机构==========")
	req := &model.CreateOrgRequest{
		Name: "GoSDKOrg2",
	}
	resp, _ := client.CreateOrg(req)
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteOrgById(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除组织机构==========")
	resp, _ := client.DeleteOrgById("617224b00869fe94de9357de")
	log.Printf("%+v\n", resp)
}

func TestClient_ListOrg(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========遍历组织机构==========")
	resp, _ := client.ListOrg(1, 10)
	log.Printf("%+v\n", resp)
}

func TestClient_GetOrgNodeById(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========根据ID获取节点==========")
	resp, _ := client.GetOrgNodeById("61725b9f3ad07a44b85302b1")
	log.Printf("%+v\n", resp)
}

func TestClient_UpdateOrgNode(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========修改节点==========")
	updateName := "updateName"
	req := &model.UpdateOrgNodeRequest{
		Name: &updateName,
		Id:   "617230eba040848abb3689b7",
	}
	resp, _ := client.UpdateOrgNode(req)
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteOrgNode(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除节点==========")
	resp, _ := client.DeleteOrgNode("617230eba040848abb3689b7", "6172315f5371116d5ad5ead9")
	log.Printf("%+v\n", resp)
}

func TestClient_IsRootNode(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========判断是否根节点==========")
	resp, _ := client.IsRootNode("6142c2c41c6e6c6cc3edfd88", "6142e08f64d5a8873598e9fb")
	log.Printf("%+v\n", resp)
}

func TestClient_MoveOrgNode(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========移动节点==========")
	resp, _ := client.MoveOrgNode("6142c2c41c6e6c6cc3edfd88", "6142e08f64d5a8873598e9fb", "6142e03436f09aa7e66c1935")
	log.Printf("%+v\n", resp)
}

func TestClient_GetRootNode(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取根节点==========")
	resp, _ := client.GetRootNode("6142c2c41c6e6c6cc3edfd88")
	log.Printf("%+v\n", resp)
}

func TestClient_ImportNodeByJSON(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========通过JSON导入==========")
	json := `
	{	
		"name": "北京非凡科技有限公司",
		"code": "feifan",
		"children": []
	}`
	resp, _ := client.ImportNodeByJSON(json)
	log.Printf("%+v\n", resp)
}

func TestClient_AddMembers(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========新增节点成员==========")
	resp, _ := client.AddMembers("61722ece541df9301478b17d", []string{"6141876341abedef979c3740"})
	log.Printf("%+v\n", resp)
}

func TestClient_MoveNodeMembers(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========移动节点成员==========")
	resp, _ := client.MoveNodeMembers("61722ece541df9301478b17d", "617230eba040848abb3689b7", []string{"6141876341abedef979c3740"})
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteNodeMembers(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除节点成员==========")
	resp, _ := client.DeleteNodeMembers("617230eba040848abb3689b7", []string{"6141876341abedef979c3740"})
	log.Printf("%+v\n", resp)
}

func TestClient_SetMainDepartment(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========设置用户主部门==========")
	resp, _ := client.SetMainDepartment("6142e0483f54818690c99600", "6141876341abedef979c3740")
	log.Printf("%+v\n", resp)
}

func TestClient_ExportByOrgId(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========设置用户主部门==========")
	resp, _ := client.ExportByOrgId("6142c2c41c6e6c6cc3edfd88")
	log.Printf("%+v\n", resp)
}

func TestClient_ListAuthorizedResourcesByNodeId(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取机构授权资源==========")
	req := &model.ListAuthorizedResourcesByIdRequest{Id: "61725b9f321fcc1ca9e36ddc"}
	resp, _ := client.ListAuthorizedResourcesByNodeId(req)
	log.Printf("%+v\n", resp)
}

func TestClient_SearchNodes(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取机构授权资源==========")

	resp, _ := client.SearchNodes("qq")
	log.Printf("%+v\n", resp)
}

func TestClient_AddOrgNode(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========权限分组列表==========")

	req := &model.AddOrgNodeRequest{
		Name:         "qqqw",
		ParentNodeId: "617230eba040848abb3689b7",
		OrgId:        "61722ececf7cd66d1ec27075",
	}
	resp, err := client.AddOrgNode(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

//
//func TestClient_StartSync(t *testing.T) {
//	client := NewClient(userPoolId, appSecret)
//	log.Println("==========获取机构授权资源==========")
//
//	resp, _ := client.StartSync( constant.WechatWork,nil)
//	log.Printf("%+v\n", resp)
//}
