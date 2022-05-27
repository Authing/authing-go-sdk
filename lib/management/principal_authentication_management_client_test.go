package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_PrincipalAuthDetail(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========主体认证详情==========")

	resp, err := client.PrincipalAuthDetail("6139c4d24e78a4d706b7545b")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
func TestClient_PrincipalAuthenticate(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========主体认证详情==========")
	req := &model.PrincipalAuthenticateRequest{
		Name:   "xx",
		Type:   constant.P,
		IdCard: "123123",
	}
	resp, err := client.PrincipalAuthenticate("6139c4d24e78a4d706b7545b", req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
