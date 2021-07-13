package management

import (
	"github.com/Authing/authing-go-sdk/lib/enum"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

const (
	userPoolId = "60a6f97f3f50c7a9483e313d"
	appSecret  = "d254623f808ba850d5d5ea7b07bead60"
)

func TestClient_ExportAll(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========导出所有组织机构数据==========")
	resp, _ := client.ExportAll()
	log.Printf("%+v\n", resp)
	log.Println("==========获取节点成员==========")
	var req = &model.ListMemberRequest{
		NodeId:               "60bdde221e3d90c0ac5efd16",
		Page:                 0,
		Limit:                0,
		IncludeChildrenNodes: false,
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
