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

func TestClient_CheckUserExists(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========检查用户是否存在==========")
	//email := "t041gyqw0b@gmail.com"
	phone := "15761403457"
	req := model.CheckUserExistsRequest{
		Email:      nil,
		Phone:      &phone,
		Username:   nil,
		ExternalId: nil,
	}
	resp, _ := client.CheckUserExists(req)
	log.Println(resp)
}
