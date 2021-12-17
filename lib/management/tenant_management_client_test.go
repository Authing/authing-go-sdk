package management

import (
	"fmt"
	"log"
	"testing"

	"github.com/Authing/authing-go-sdk/lib/model"
)

func TestClient_GetTenantList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取用户池下租户列表==========")
	resp, err := client.GetTenantList(&model.CommonPageRequest{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_GetTenantDetails(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========根据租户 ID 查询租户==========")
	resp, err := client.GetTenantDetails("61b83950c110f5a2955221df")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_CreateTenant(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建租户==========")
	resp, err := client.CreateTenant(&model.CreateTenantRequest{
		Name:   "测试lnoi",
		AppIds: "61503af19ddff2aa185b665a",
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_UpdateTenant(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========修改租户==========")
	resp, err := client.UpdateTenant("61b95412098eb8dd16d5a7f4", &model.CreateTenantRequest{
		Name: "测试 go 修改eve",
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteTenant(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除租户==========")
	resp, err := client.RemoveTenant("61b95412098eb8dd16d5a7f4")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ConfigTenant(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========配置租户品牌化==========")
	resp, err := client.ConfigTenant("61b83950c110f5a2955221df", &model.ConfigTenantRequest{
		CSS: ".btnId {\n text-color: #ffff}",
		SsoPageCustomizationSettings: &model.TenantSsoPageCustomizationSettings{
			AutoRegisterThenLogin: false,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_GetTenantMembers(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取租户成员列表==========")
	resp, err := client.GetTenantMembers("61b83950c110f5a2955221df", &model.CommonPageRequest{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_AddTenantMembers(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========添加租户成员==========")
	resp, err := client.AddTenantMembers("61b83950c110f5a2955221df", []string{"61b85b9da80ac34ac3a9451d", "61b85b945468e9865acae737"})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_RemoveTenantMembers(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除租户成员==========")
	resp, err := client.RemoveTenantMembers("61b83950c110f5a2955221df", "61b85b9da80ac34ac3a9451d")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListExtIdp(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取身份源列表==========")
	resp, err := client.ListExtIdp("61b83950c110f5a2955221df")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ExtIdpDetail(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取身份源==========")

	resp, err := client.ExtIdpDetail("61b868aea25030db174529f1")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_CreateExtIdp(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建身份源==========")
	fields := map[string]string{
		"displayName":  "测试创建",
		"baseURL":      "https://gitlab.com/wfr",
		"clientID":     "everwew",
		"clientSecret": "everwew",
	}

	resp, err := client.CreateExtIdp(&model.CreateExtIdpRequest{
		Name:     "GitLab",
		Type:     "gitlab",
		TenantUd: "61b83950c110f5a2955221df",
		Connections: []model.ExtIdpConnection{{
			Identifier:  "nboenboei",
			Type:        "gitlab",
			DisplayName: "测试创建envoengoie",
			Fields:      fields,
		}},
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_UpdateExtIdp(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========更新身份源==========")
	resp, err := client.UpdateExtIdp("61b958a18a3f153bf3674e5b", &model.UpdateExtIdpRequest{
		Name: "cscwecw",
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteExtIdp(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除身份源==========")
	resp, err := client.DeleteExtIdp("61b958a18a3f153bf3674e5b")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_CreateExtIdpConnection(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建身份源==========")
	fields := map[string]string{
		"displayName":  "测试创建连接",
		"baseURL":      "https://gitlab.com/123456",
		"clientID":     "123456",
		"clientSecret": "123456",
	}

	resp, err := client.CreateExtIdpConnection(&model.CreateExtIdpConnectionRequest{
		ExtIdpId:    "61b955fd8f70040602f8ebe4",
		Identifier:  "prmoroorobrnro",
		Type:        "gitlab",
		DisplayName: "测试创建envoengoioi",
		Fields:      fields,
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_UpdateExtIdpConnection(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("=========更新身份源连接==========")
	fields := map[string]string{
		"displayName":  "测试连接修改2",
		"baseURL":      "https://gitlab.com/123456",
		"clientID":     "123456",
		"clientSecret": "123456",
	}

	resp, err := client.UpdateExtIdpConnection("61b9602bac8e32162db6d9d5", &model.UpdateExtIdpConnectionRequest{
		DisplayName: "测试连接修改2",
		Fields:      fields,
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteExtIdpConnection(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除身份源连接==========")

	resp, err := client.DeleteExtIdpConnection("61b9602bac8e32162db6d9d5")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_CheckExtIdpConnectionIdentifierUnique(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========检查连接唯一标识是否冲突==========")

	resp, err := client.CheckExtIdpConnectionIdentifierUnique("emoo")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ChangeExtIdpConnectionState(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========开关身份源连接==========")

	resp, err := client.ChangeExtIdpConnectionState("61b868ae560f5e2ef2bd9e91", &model.ChangeExtIdpConnectionStateRequest{
		Enabled:  true,
		TenantID: "61b83950c110f5a2955221df",
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_BatchChangeExtIdpConnectionState(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========批量开关身份源连接==========")

	resp, err := client.BatchChangeExtIdpConnectionState("61b98798fab83706ed7f853f", &model.ChangeExtIdpConnectionStateRequest{
		Enabled:  false,
		TenantID: "61b83950c110f5a2955221df",
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
