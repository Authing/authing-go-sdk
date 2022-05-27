package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_CreateNamespace(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建权限分组==========")
	code := "qCode"
	name := "qName"
	req := &model.EditNamespaceRequest{
		Code: &code,
		Name: &name,
	}
	resp, err := client.CreateNamespace(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_UpdateNamespace(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========修改权限分组==========")
	code := "qCodeww"
	name := "qNameww"
	req := &model.EditNamespaceRequest{
		Code: &code,
		Name: &name,
	}
	resp, err := client.UpdateNamespace("54156", req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListNamespace(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========权限分组列表==========")

	resp, err := client.ListNamespace(1, 10)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DeleteNamespace(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========权限分组列表==========")

	resp, err := client.DeleteNamespace("54156")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
