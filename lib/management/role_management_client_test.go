package management

import (
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
	req := model.GetRoleUserListRequest{
		Page:      1,
		Limit:     10,
		Code:      "develop",
		Namespace: "default",
	}
	resp, _ := client.GetRoleUserList(req)
	log.Printf("%+v\n", resp)
}
