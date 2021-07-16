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
