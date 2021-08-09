package authentication

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	"log"
	"testing"
)

const (
	AppId  = "60a6f980dd9a9a7642da768a"
	Secret = "5cd4ea7b3603b792aea9a00da9e18f44"
)

func TestClient_BuildAuthorizeUrlByOidc(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.Protocol = constant.OIDC
	authenticationClient.TokenEndPointAuthMethod = constant.None
	req := model.OidcParams{
		AppId:       AppId,
		RedirectUri: "https://mvnrepository.com/",
		Nonce:       "test",
	}
	resp, err := authenticationClient.BuildAuthorizeUrlByOidc(req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_GetAccessTokenByCode(t *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a", "5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.Host = "https://32l5hb-demo.authing.cn"
	authenticationClient.RedirectUri = "https://mvnrepository.com/"
	authenticationClient.Protocol = constant.OIDC
	authenticationClient.TokenEndPointAuthMethod = constant.None
	resp, err := authenticationClient.GetAccessTokenByCode("vj-MWd4eRZdmakwobde53RaFZpBON3-khElsrlEZRGm")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
	// {"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjQ0bnJHU05YQ3NDLTByd1J5Q0hENjBzdmc0elpLNF9iV2VnQjluOFRhQzQifQ.eyJqdGkiOiJ3NjJmNkVieHYxd19wbEV3YWMwWlIiLCJzdWIiOiI2MGUyNmI2ZjdiMGRkN2MwYWY4M2VjZDkiLCJpYXQiOjE2MjU0OTI3NjUsImV4cCI6MTYyNjcwMjM2NSwic2NvcGUiOiJvcGVuaWQgcGhvbmUgYWRkcmVzcyBwcm9maWxlIGVtYWlsIiwiaXNzIjoiaHR0cHM6Ly8zMmw1aGItZGVtby5hdXRoaW5nLmNuL29pZGMiLCJhdWQiOiI2MGE2Zjk4MGRkOWE5YTc2NDJkYTc2OGEifQ.KOMWqEtbyH3qdBv_bHX3Dof2t_3XBQ7QDg4-x7fIr9W2YtCnwNnqVehOVYjWpcF-pkVyzBlpmKIc6_X9F8GA-oYbdUKJzhxfoAATj1JnRCRs6Wsxpo3U41up1pgXs5B7JS7gVbiw_IucMg4vLYw_QJ_aPgBTkjCkBZVsPf3NRYCd2cVwiZwvoa8GT6jGP9PJ908rJSSSdsqt6JNzydVbJ9a7p4mBhV3WxUAckXePjIE0QDNDe_GxFwFDktkTbLBIJZBL4bSg3pHGQKHiF9wabfjBRfWV8ChRe8i95n7pq-Gw9fw2fKNv7ieC5bK52D1j6R9L5h7wRvTstgiR7p8krQ","expires_in":1209600,"id_token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjQ0bnJHU05YQ3NDLTByd1J5Q0hENjBzdmc0elpLNF9iV2VnQjluOFRhQzQifQ.eyJzdWIiOiI2MGUyNmI2ZjdiMGRkN2MwYWY4M2VjZDkiLCJwaG9uZV9udW1iZXIiOm51bGwsInBob25lX251bWJlcl92ZXJpZmllZCI6ZmFsc2UsImFkZHJlc3MiOnsiY291bnRyeSI6bnVsbCwicG9zdGFsX2NvZGUiOm51bGwsInJlZ2lvbiI6bnVsbCwiZm9ybWF0dGVkIjpudWxsfSwiYmlydGhkYXRlIjpudWxsLCJmYW1pbHlfbmFtZSI6bnVsbCwiZ2VuZGVyIjoiVSIsImdpdmVuX25hbWUiOm51bGwsImxvY2FsZSI6bnVsbCwibWlkZGxlX25hbWUiOm51bGwsIm5hbWUiOm51bGwsIm5pY2tuYW1lIjpudWxsLCJwaWN0dXJlIjoiaHR0cHM6Ly9maWxlcy5hdXRoaW5nLmNvL2F1dGhpbmctY29uc29sZS9kZWZhdWx0LXVzZXItYXZhdGFyLnBuZyIsInByZWZlcnJlZF91c2VybmFtZSI6bnVsbCwicHJvZmlsZSI6bnVsbCwidXBkYXRlZF9hdCI6IjIwMjEtMDctMDVUMTM6NDU6MjMuMTc0WiIsIndlYnNpdGUiOm51bGwsInpvbmVpbmZvIjpudWxsLCJlbWFpbCI6Imx1b2ppZWxpbkBhdXRoaW5nLmNuIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJub25jZSI6InRlc3QiLCJhdF9oYXNoIjoiaFVfa2o5NzdjSGUxXzdhdWI3OWY3dyIsImF1ZCI6IjYwYTZmOTgwZGQ5YTlhNzY0MmRhNzY4YSIsImV4cCI6MTYyNjcwMjM2NSwiaWF0IjoxNjI1NDkyNzY1LCJpc3MiOiJodHRwczovLzMybDVoYi1kZW1vLmF1dGhpbmcuY24vb2lkYyJ9.XtLA_hQZqqwUW2GyVwEVhO2BMqNCFMWCkxQGd1FP37tclxnHKsa26wz8oBKNPXsGwEUBIlcyzi9SCTOibl_UlG4hNrHASNkk_2zQcsjO8fidHfXjEyw2UjhDfxsyh1B6xcJIiM8AJIQi5BHJ1FcFzCLxRK81v_kPqQMMHagYXEQhaFNf-otxrBrf9Yc66wuMLKlgKUgAZLyhTqJFpXPIayzss00vIOvbQNTc5XY27M_uUP2-TInIG8dxY-rcxe06PqTWVvLkDx1CMsEC7Ume1wf6lKqGU4kGnSLlXBxrl1-MRd-Q01gosvBvP2r2Tuxb30ZD0-yG4QY9yD9ytTYSPA","scope":"openid phone address profile email","token_type":"Bearer"}
}

func TestClient_GetUserInfoByAccessToken(t *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a", "5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.Host = "https://32l5hb-demo.authing.cn"
	authenticationClient.RedirectUri = "https://mvnrepository.com/"
	authenticationClient.Protocol = constant.OIDC
	authenticationClient.TokenEndPointAuthMethod = constant.None
	resp, err := authenticationClient.GetUserInfoByAccessToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjQ0bnJHU05YQ3NDLTByd1J5Q0hENjBzdmc0elpLNF9iV2VnQjluOFRhQzQifQ.eyJqdGkiOiJ3NjJmNkVieHYxd19wbEV3YWMwWlIiLCJzdWIiOiI2MGUyNmI2ZjdiMGRkN2MwYWY4M2VjZDkiLCJpYXQiOjE2MjU0OTI3NjUsImV4cCI6MTYyNjcwMjM2NSwic2NvcGUiOiJvcGVuaWQgcGhvbmUgYWRkcmVzcyBwcm9maWxlIGVtYWlsIiwiaXNzIjoiaHR0cHM6Ly8zMmw1aGItZGVtby5hdXRoaW5nLmNuL29pZGMiLCJhdWQiOiI2MGE2Zjk4MGRkOWE5YTc2NDJkYTc2OGEifQ.KOMWqEtbyH3qdBv_bHX3Dof2t_3XBQ7QDg4-x7fIr9W2YtCnwNnqVehOVYjWpcF-pkVyzBlpmKIc6_X9F8GA-oYbdUKJzhxfoAATj1JnRCRs6Wsxpo3U41up1pgXs5B7JS7gVbiw_IucMg4vLYw_QJ_aPgBTkjCkBZVsPf3NRYCd2cVwiZwvoa8GT6jGP9PJ908rJSSSdsqt6JNzydVbJ9a7p4mBhV3WxUAckXePjIE0QDNDe_GxFwFDktkTbLBIJZBL4bSg3pHGQKHiF9wabfjBRfWV8ChRe8i95n7pq-Gw9fw2fKNv7ieC5bK52D1j6R9L5h7wRvTstgiR7p8krQ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_IntrospectToken(t *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a", "5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.Host = "https://32l5hb-demo.authing.cn"
	authenticationClient.RedirectUri = "https://mvnrepository.com/"
	authenticationClient.Protocol = constant.OIDC
	authenticationClient.TokenEndPointAuthMethod = constant.None

	resp, err := authenticationClient.IntrospectToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjQ0bnJHU05YQ3NDLTByd1J5Q0hENjBzdmc0elpLNF9iV2VnQjluOFRhQzQifQ.eyJqdGkiOiJ3NjJmNkVieHYxd19wbEV3YWMwWlIiLCJzdWIiOiI2MGUyNmI2ZjdiMGRkN2MwYWY4M2VjZDkiLCJpYXQiOjE2MjU0OTI3NjUsImV4cCI6MTYyNjcwMjM2NSwic2NvcGUiOiJvcGVuaWQgcGhvbmUgYWRkcmVzcyBwcm9maWxlIGVtYWlsIiwiaXNzIjoiaHR0cHM6Ly8zMmw1aGItZGVtby5hdXRoaW5nLmNuL29pZGMiLCJhdWQiOiI2MGE2Zjk4MGRkOWE5YTc2NDJkYTc2OGEifQ.KOMWqEtbyH3qdBv_bHX3Dof2t_3XBQ7QDg4-x7fIr9W2YtCnwNnqVehOVYjWpcF-pkVyzBlpmKIc6_X9F8GA-oYbdUKJzhxfoAATj1JnRCRs6Wsxpo3U41up1pgXs5B7JS7gVbiw_IucMg4vLYw_QJ_aPgBTkjCkBZVsPf3NRYCd2cVwiZwvoa8GT6jGP9PJ908rJSSSdsqt6JNzydVbJ9a7p4mBhV3WxUAckXePjIE0QDNDe_GxFwFDktkTbLBIJZBL4bSg3pHGQKHiF9wabfjBRfWV8ChRe8i95n7pq-Gw9fw2fKNv7ieC5bK52D1j6R9L5h7wRvTstgiR7p8krQ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_ValidateToken(t *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a", "5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.Host = "https://32l5hb-demo.authing.cn"
	authenticationClient.RedirectUri = "https://mvnrepository.com/"
	authenticationClient.Protocol = constant.OIDC
	authenticationClient.TokenEndPointAuthMethod = constant.None
	req := model.ValidateTokenRequest{
		AccessToken: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjQ0bnJHU05YQ3NDLTByd1J5Q0hENjBzdmc0elpLNF9iV2VnQjluOFRhQzQifQ.eyJqdGkiOiJ3NjJmNkVieHYxd19wbEV3YWMwWlIiLCJzdWIiOiI2MGUyNmI2ZjdiMGRkN2MwYWY4M2VjZDkiLCJpYXQiOjE2MjU0OTI3NjUsImV4cCI6MTYyNjcwMjM2NSwic2NvcGUiOiJvcGVuaWQgcGhvbmUgYWRkcmVzcyBwcm9maWxlIGVtYWlsIiwiaXNzIjoiaHR0cHM6Ly8zMmw1aGItZGVtby5hdXRoaW5nLmNuL29pZGMiLCJhdWQiOiI2MGE2Zjk4MGRkOWE5YTc2NDJkYTc2OGEifQ.KOMWqEtbyH3qdBv_bHX3Dof2t_3XBQ7QDg4-x7fIr9W2YtCnwNnqVehOVYjWpcF-pkVyzBlpmKIc6_X9F8GA-oYbdUKJzhxfoAATj1JnRCRs6Wsxpo3U41up1pgXs5B7JS7gVbiw_IucMg4vLYw_QJ_aPgBTkjCkBZVsPf3NRYCd2cVwiZwvoa8GT6jGP9PJ908rJSSSdsqt6JNzydVbJ9a7p4mBhV3WxUAckXePjIE0QDNDe_GxFwFDktkTbLBIJZBL4bSg3pHGQKHiF9wabfjBRfWV8ChRe8i95n7pq-Gw9fw2fKNv7ieC5bK52D1j6R9L5h7wRvTstgiR7p8krQ",
		IdToken:     "",
	}
	resp, err := authenticationClient.ValidateToken(req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_GetAccessTokenByClientCredentials(t *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a", "5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.Host = "https://32l5hb-demo.authing.cn"
	authenticationClient.RedirectUri = "https://mvnrepository.com/"
	authenticationClient.Protocol = constant.OIDC
	authenticationClient.TokenEndPointAuthMethod = constant.None
	input := model.ClientCredentialInput{
		AccessKey: "",
		SecretKey: "",
	}
	req := model.GetAccessTokenByClientCredentialsRequest{
		Scope:                 "openid",
		ClientCredentialInput: &input,
	}
	resp, err := authenticationClient.GetAccessTokenByClientCredentials(req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_RevokeToken(t *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a", "5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.Host = "https://32l5hb-demo.authing.cn"
	authenticationClient.RedirectUri = "https://mvnrepository.com/"
	authenticationClient.Protocol = constant.OIDC
	authenticationClient.TokenEndPointAuthMethod = constant.None
	resp, err := authenticationClient.RevokeToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjQ0bnJHU05YQ3NDLTByd1J5Q0hENjBzdmc0elpLNF9iV2VnQjluOFRhQzQifQ.eyJqdGkiOiJ3NjJmNkVieHYxd19wbEV3YWMwWlIiLCJzdWIiOiI2MGUyNmI2ZjdiMGRkN2MwYWY4M2VjZDkiLCJpYXQiOjE2MjU0OTI3NjUsImV4cCI6MTYyNjcwMjM2NSwic2NvcGUiOiJvcGVuaWQgcGhvbmUgYWRkcmVzcyBwcm9maWxlIGVtYWlsIiwiaXNzIjoiaHR0cHM6Ly8zMmw1aGItZGVtby5hdXRoaW5nLmNuL29pZGMiLCJhdWQiOiI2MGE2Zjk4MGRkOWE5YTc2NDJkYTc2OGEifQ.KOMWqEtbyH3qdBv_bHX3Dof2t_3XBQ7QDg4-x7fIr9W2YtCnwNnqVehOVYjWpcF-pkVyzBlpmKIc6_X9F8GA-oYbdUKJzhxfoAATj1JnRCRs6Wsxpo3U41up1pgXs5B7JS7gVbiw_IucMg4vLYw_QJ_aPgBTkjCkBZVsPf3NRYCd2cVwiZwvoa8GT6jGP9PJ908rJSSSdsqt6JNzydVbJ9a7p4mBhV3WxUAckXePjIE0QDNDe_GxFwFDktkTbLBIJZBL4bSg3pHGQKHiF9wabfjBRfWV8ChRe8i95n7pq-Gw9fw2fKNv7ieC5bK52D1j6R9L5h7wRvTstgiR7p8krQ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestClient_LoginByUserName(t *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a", "5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.userPoolId = "60e043f8cd91b87d712b6365"
	authenticationClient.Secret = "158c7679333bc196b524d78d745813e5"
	req := model.LoginByUsernameInput{
		Username:     "luojielin",
		Password:     "12341",
		CaptchaCode:  nil,
		AutoRegister: nil,
		ClientIp:     nil,
		Params:       nil,
		Context:      nil,
	}
	resp, err := authenticationClient.LoginByUserName(req)
	log.Println(resp, err)
}

func TestClient_LoginByEmail(t *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a", "5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.userPoolId = "60e043f8cd91b87d712b6365"
	authenticationClient.Secret = "158c7679333bc196b524d78d745813e5"
	req := model.LoginByEmailInput{
		Email:        "luojielin@authing.cn",
		Password:     "1234",
		CaptchaCode:  nil,
		AutoRegister: nil,
		ClientIp:     nil,
		Params:       nil,
		Context:      nil,
	}
	resp, err := authenticationClient.LoginByEmail(req)
	log.Println(resp, err)
}

func TestClient_LoginByPhonePassword(b *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a", "5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.userPoolId = "60e043f8cd91b87d712b6365"
	authenticationClient.Secret = "158c7679333bc196b524d78d745813e5"
	req := model.LoginByPhonePasswordInput{
		Phone:        "18310641137",
		Password:     "1234",
		CaptchaCode:  nil,
		AutoRegister: nil,
		ClientIp:     nil,
		Params:       nil,
		Context:      nil,
	}
	resp, err := authenticationClient.LoginByPhonePassword(req)
	log.Println(resp, err)
}

/*func TestClient_LoginByPhoneCode(b *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a","5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.userPoolId = "60e043f8cd91b87d712b6365"
	authenticationClient.Secret = "158c7679333bc196b524d78d745813e5"
	req := model.LoginByPhoneCodeInput{
		Phone:        "18310641137",
		Code:         "7458",
		AutoRegister: nil,
		ClientIp:     nil,
		Params:       nil,
		Context:      nil,
	}
	resp,err := authenticationClient.LoginByPhoneCode(req)
	log.Println(resp,err)
}

func TestClient_SendSmsCode(t *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a","5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.userPoolId = "60e043f8cd91b87d712b6365"
	authenticationClient.Secret = "158c7679333bc196b524d78d745813e5"
	resp,err := authenticationClient.SendSmsCode("15566416161")
	log.Println(resp,err)
}*/
