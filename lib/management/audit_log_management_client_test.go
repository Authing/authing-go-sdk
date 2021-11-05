package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_ListAuditLogs(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========审计日志列表==========")
	var userIds = []string{"xx", "xxq"}
	page := 1
	limit := 10
	req := &model.ListAuditLogsRequest{
		Page:    &page,
		Limit:   &limit,
		UserIds: &userIds,
	}
	resp, err := client.ListAuditLogs(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListUserActionLogs(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========用户操作日志列表==========")
	var userIds = []string{"xx", "xxq"}
	page := 1
	limit := 10
	req := &model.ListUserActionRequest{
		Page:    &page,
		Limit:   &limit,
		UserIds: &userIds,
	}
	resp, err := client.ListUserAction(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
