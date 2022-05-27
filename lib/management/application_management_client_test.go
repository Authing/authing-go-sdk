package management

import (
	"fmt"
	"log"
	"testing"

	"github.com/Authing/authing-go-sdk/lib/model"
)

func TestClient_ListApplication(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========查询应用列表==========")

	req := &model.CommonPageRequest{
		Page:  1,
		Limit: 10,
	}
	resp, err := client.ListApplication(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_CreateApplication(t *testing.T) {
	log.Println(userPoolId)
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建应用==========")
	resp, err := client.CreateApplication("sqq12", "noww22", "http://locaqql", nil)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteApplication(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除应用==========")
	resp, err := client.DeleteApplication("616fbde39a4c5ce0518d87fc")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_RefreshApplicationSecret(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========刷新应用秘钥==========")
	resp, err := client.RefreshApplicationSecret("614bf4af279893d5ab645e58")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListApplicationActiveUsers(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取应用下登录用户==========")
	resp, err := client.ListApplicationActiveUsers("61527e0124a5f0df0eed7af2", 1, 100)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Println(resp.TotalCount)
}

func TestClient_FindApplicationById(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========根据ID获取应用==========")
	resp, err := client.FindApplicationById("614bf4af279893d5ab645e58")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_CreateApplicationAgreement(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建应用策略==========")
	resp, err := client.CreateApplicationAgreement("614bf4af279893d5ab645e58", "cccqq", nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListApplicationAgreement(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========遍历应用策略==========")
	resp, err := client.ListApplicationAgreement("614bf4af279893d5ab645e58")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ModifyApplicationAgreement(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========修改应用策略==========")
	resp, err := client.ModifyApplicationAgreement("614bf4af279893d5ab645e58", "249", "cccqq2", nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteApplicationAgreement(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除应用策略==========")
	resp, err := client.DeleteApplicationAgreement("614bf4af279893d5ab645e58", "249")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_SortApplicationAgreement(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除应用策略==========")
	resp, err := client.SortApplicationAgreement("614bf4af279893d5ab645e58", []string{"238"})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ApplicationTenants(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取应用关联租户==========")
	resp, err := client.ApplicationTenants("61b8366efa768b57d65b6394")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
