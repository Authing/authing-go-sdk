package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	"github.com/Authing/authing-go-sdk/lib/util"
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

func TestClient_ListMembers(t *testing.T) {
	dataMap := map[string]string{
		"a": "1",
		"b": "2",
	}
	fmt.Println(util.GetQueryString(dataMap))
}
