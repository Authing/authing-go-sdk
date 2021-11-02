package management

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_CreatePolicy(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建策略==========")
	ef := model.EnumPolicyEffectAllow
	stateMents := &model.PolicyStatement{
		Resource: "book:222c",
		Effect:   &ef,
		Actions:  []string{"'booksc:read'"},
	}
	req := &model.PolicyRequest{
		Code:       "qqx",
		Statements: []model.PolicyStatement{*stateMents},
	}
	resp, err := client.CreatePolicy(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListPolicy(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========创建策略==========")

	resp, err := client.ListPolicy(1, 10)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DetailPolicy(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========策略详情==========")

	resp, err := client.DetailPolicy("qqx")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", *resp.Statements[0].Effect)
}

func TestClient_UpdatePolicy(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========修改策略==========")
	ef := model.EnumPolicyEffectAllow
	stateMents := &model.PolicyStatement{
		Resource: "book:222cw",
		Effect:   &ef,
		Actions:  []string{"'booksc:read'"},
	}
	req := &model.PolicyRequest{
		Code:       "qqx",
		Statements: []model.PolicyStatement{*stateMents},
	}
	resp, err := client.UpdatePolicy(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DeletePolicy(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除策略==========")

	resp, err := client.DeletePolicy("qqx")
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ListAssignments(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========删除策略==========")

	resp, err := client.ListAssignments("tliewdutrn", 1, 10)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_AddAssignments(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========授权策略==========")
	req := &model.PolicyAssignmentsRequest{
		Policies:          []string{"tliewdutrn"},
		TargetType:        model.EnumPolicyAssignmentTargetTypeUser,
		TargetIdentifiers: []string{"616e905ebc18f0f106973a29"},
	}
	resp, err := client.AddAssignments(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_RemoveAssignments(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========撤销策略==========")
	req := &model.PolicyAssignmentsRequest{
		Policies:          []string{"tliewdutrn"},
		TargetType:        model.EnumPolicyAssignmentTargetTypeUser,
		TargetIdentifiers: []string{"616e905ebc18f0f106973a29"},
	}
	resp, err := client.RemoveAssignments(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_EnableAssignments(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========撤销策略==========")
	req := &model.SwitchPolicyAssignmentsRequest{
		Policy:           "tliewdutrn",
		TargetType:       model.EnumPolicyAssignmentTargetTypeUser,
		TargetIdentifier: "616e905ebc18f0f106973a29",
	}
	resp, err := client.EnableAssignments(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_DisableAssignments(t *testing.T) {
	client := NewClient(userPoolId, appSecret)
	log.Println("==========撤销策略==========")
	req := &model.SwitchPolicyAssignmentsRequest{
		Policy:           "tliewdutrn",
		TargetType:       model.EnumPolicyAssignmentTargetTypeUser,
		TargetIdentifier: "616e905ebc18f0f106973a29",
	}
	resp, err := client.DisableAssignments(req)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v\n", resp)
}
