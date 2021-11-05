package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_ListUdf(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========自定义字段列表==========")
	resp, err := client.ListUdf(model.EnumUDFTargetTypeUSER)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_SetUdf(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========自定义字段列表==========")
	req := &model.SetUdfInput{
		TargetType: model.EnumUDFTargetTypeUSER,
		DataType:   model.EnumUDFDataTypeSTRING,
		Key:        "goSDK",
		Label:      "goSDK",
	}
	resp, err := client.SetUdf(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_RemoveUdf(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========自定义字段列表==========")

	resp, err := client.RemoveUdf(model.EnumUDFTargetTypeUSER, "goSDK")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListUdfValue(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========某对象自定义字段列表==========")

	resp, err := client.ListUdfValue(model.EnumUDFTargetTypeUSER, "616d41b7410a33da0cb70e65")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_SetUdvBatch(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========某对象自定义字段列表==========")

	resp, err := client.SetUdvBatch("616d41b7410a33da0cb70e65", model.EnumUDFTargetTypeUSER, &[]model.KeyValuePair{
		{Key: "goSDK", Value: "goSDK"},
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
