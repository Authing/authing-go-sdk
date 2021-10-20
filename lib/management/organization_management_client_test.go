package management

import (
	"github.com/Authing/authing-go-sdk/lib/enum"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

const (
	// prod
	userPoolId = "60e043f8cd91b87d712b6365"
	appSecret  = "158c7679333bc196b524d78d745813e5"
	// dev
	userPoolIdDev = "61090ca2ae21b81053abbd07"
	appSecretDev  = "db3e0a32cd5629fe12c9d29911abb9b7"
	//userPoolId = "60e043f8cd91b87d712b6365"
	//appSecret  = "158c7679333bc196b524d78d745813e5"
	//userPoolId = "6114ea3b25851f2e44db357f"
	//appSecret  = "4f673a16f53cbbf54633212b1a882a2a"
	//userPoolId = "61384d3e302f1f75e69ce95a"
	//appSecret  = "ff053c05e4fb664a560556ea7c2cb715"
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
