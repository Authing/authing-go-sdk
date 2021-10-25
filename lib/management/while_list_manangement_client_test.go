package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_GetWhileList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取白名单==========")

	resp, err := client.GetWhileList(model.EnumWhitelistTypeUSERNAME)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_AddWhileList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========获取白名单==========")

	resp, err := client.AddWhileList(model.EnumWhitelistTypeUSERNAME, []string{"qqxccx", "qweqwe"})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_RemoveWhileList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========移除白名单==========")

	resp, err := client.RemoveWhileList(model.EnumWhitelistTypeUSERNAME, []string{"qqxccx"})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_EnableWhileList(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========移除白名单==========")

	resp, err := client.EnableWhileList(model.EnumWhitelistTypeUSERNAME)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
