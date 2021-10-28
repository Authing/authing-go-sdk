package authentication

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

func TestClient_GetMfaAuthenticators(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	//loginReq:= model.LoginByEmailInput{
	//	Email: "fptvmzqyxn@authing.cn",
	//	Password: "12345678",
	//}
	//u,e:=authenticationClient.LoginByEmail(loginReq)
	//log.Println(u)
	//log.Println(e)
	resp, err := authenticationClient.GetMfaAuthenticators(&model.MfaInput{
		//MfaToken: u.Token,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_AssociateMfaAuthenticator(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	loginReq := model.LoginByEmailInput{
		Email:    "fptvmzqyxn@authing.cn",
		Password: "12345678",
	}
	u, e := authenticationClient.LoginByEmail(loginReq)
	log.Println(e)
	resp, err := authenticationClient.AssociateMfaAuthenticator(&model.MfaInput{
		MfaToken: u.Token,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_DeleteMfaAuthenticator(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	resp, err := authenticationClient.DeleteMfaAuthenticator()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_ConfirmAssociateMfaAuthenticator(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	//loginReq:= model.LoginByEmailInput{
	//	Email: "fptvmzqyxn@authing.cn",
	//	Password: "12345678",
	//}
	//u,e:=authenticationClient.LoginByEmail(loginReq)
	//log.Println(e)
	resp, err := authenticationClient.ConfirmAssociateMfaAuthenticator(&model.ConfirmAssociateMfaAuthenticatorRequest{
		Totp: "D5LH4GQQGEEWEHKX",
		//Totp: "c833-cbb6-9180-7240-a048-ebe6",
		//MfaToken: u.Token,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_VerifyTotpMfa(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	mfaToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJQb29sSWQiOiI2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODYiLCJ1c2VySWQiOiI2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJhcm4iOiJhcm46Y246YXV0aGluZzo2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODY6dXNlcjo2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJzdGFnZSI6MX0sImlhdCI6MTYzNTE0OTQ2MiwiZXhwIjoxNjM1MTQ5ODIyfQ.2DbmVf1-JQeiRMpZBk-3y-uPIN15FL-ranE4UlMKMoM"

	resp, err := authenticationClient.VerifyTotpMfa("q", mfaToken)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_VerifyAppSmsMfa(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	loginReq := model.LoginByEmailInput{
		Email:    "gosdk@mail.com",
		Password: "123456789",
	}
	u, e := authenticationClient.LoginByEmail(loginReq)
	log.Println(e)
	resp, err := authenticationClient.VerifyAppSmsMfa("777777", "q", *u.Token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_VerifyAppEmailMfa(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	loginReq := model.LoginByEmailInput{
		Email:    "gosdk@mail.com",
		Password: "123456789",
	}
	u, e := authenticationClient.LoginByEmail(loginReq)
	log.Println(u, e)
	mfaToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJQb29sSWQiOiI2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODYiLCJ1c2VySWQiOiI2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJhcm4iOiJhcm46Y246YXV0aGluZzo2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODY6dXNlcjo2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJzdGFnZSI6MX0sImlhdCI6MTYzNTE0OTQ2MiwiZXhwIjoxNjM1MTQ5ODIyfQ.2DbmVf1-JQeiRMpZBk-3y-uPIN15FL-ranE4UlMKMoM"

	resp, err := authenticationClient.VerifyAppEmailMfa("gosdk@mail.com", "q", mfaToken)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_PhoneOrEmailBindable(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	mfaToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJQb29sSWQiOiI2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODYiLCJ1c2VySWQiOiI2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJhcm4iOiJhcm46Y246YXV0aGluZzo2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODY6dXNlcjo2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJzdGFnZSI6MX0sImlhdCI6MTYzNTE0OTQ2MiwiZXhwIjoxNjM1MTQ5ODIyfQ.2DbmVf1-JQeiRMpZBk-3y-uPIN15FL-ranE4UlMKMoM"
	email := "gosdk@mail.com"
	resp, err := authenticationClient.PhoneOrEmailBindable(&email, nil, mfaToken)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_VerifyFaceMfa(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool

	mfaToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJQb29sSWQiOiI2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODYiLCJ1c2VySWQiOiI2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJhcm4iOiJhcm46Y246YXV0aGluZzo2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODY6dXNlcjo2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJzdGFnZSI6MX0sImlhdCI6MTYzNTE0OTQ2MiwiZXhwIjoxNjM1MTQ5ODIyfQ.2DbmVf1-JQeiRMpZBk-3y-uPIN15FL-ranE4UlMKMoM"

	resp, err := authenticationClient.VerifyFaceMfa("http://face", mfaToken)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_AssociateFaceByUrl(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool

	mfaToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJQb29sSWQiOiI2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODYiLCJ1c2VySWQiOiI2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJhcm4iOiJhcm46Y246YXV0aGluZzo2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODY6dXNlcjo2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJzdGFnZSI6MX0sImlhdCI6MTYzNTE0OTQ2MiwiZXhwIjoxNjM1MTQ5ODIyfQ.2DbmVf1-JQeiRMpZBk-3y-uPIN15FL-ranE4UlMKMoM"

	resp, err := authenticationClient.AssociateFaceByUrl("http://tp", "http://zp", mfaToken)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_VerifyTotpRecoveryCode(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool

	mfaToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJQb29sSWQiOiI2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODYiLCJ1c2VySWQiOiI2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJhcm4iOiJhcm46Y246YXV0aGluZzo2MGMxN2IzZDcyYjkyNTA5N2E3MzhkODY6dXNlcjo2MTc2NWYxMDI5MThhOGZjNjUyNDU2NDAiLCJzdGFnZSI6MX0sImlhdCI6MTYzNTE0OTQ2MiwiZXhwIjoxNjM1MTQ5ODIyfQ.2DbmVf1-JQeiRMpZBk-3y-uPIN15FL-ranE4UlMKMoM"

	resp, err := authenticationClient.VerifyTotpMfa("eedc-58ed-931b-8967-a092-46ae", mfaToken)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
