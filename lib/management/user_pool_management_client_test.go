package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_UserPoolDetail(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========用户池详情==========")
	resp, err := client.UserPoolDetail()
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_UpdateUserPool(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========修改用户池==========")
	userPoolName := "otherSdk9989995"
	req := &model.UpdateUserpoolInput{
		Name:   &userPoolName,
		Domain: &userPoolName,
	}
	resp, err := client.UpdateUserPool(*req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_UserPoolEnv(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========用户池环境变量==========")

	resp, err := client.ListUserPoolEnv()
	if err != nil {
		fmt.Println(err)
	}
	resp1, err1 := client.AddUserPoolEnv("qnm", "qnm")
	fmt.Println(resp1, err1)
	resp2, err2 := client.RemoveUserPoolEnv("qnm")
	fmt.Println(resp2, err2)
	resp, err = client.ListUserPoolEnv()
	log.Printf("%+v\n", resp)
}
