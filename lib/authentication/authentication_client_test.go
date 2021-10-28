package authentication

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"log"
	"testing"
)

const (
	//UserPool="60c17b3d72b925097a738d86"
	//Secret="6a350fe221596e96125e9375452da606"
	//AppId ="60c17b536f0f06def12dfec4"
	AppId    = "6168f95e81d5e20f9cb72f22"
	Secret   = "ff053c05e4fb664a560556ea7c2cb715"
	UserPool = "61384d3e302f1f75e69ce95a"
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

func TestClient_GetCurrentUser(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool

	resp, err := authenticationClient.GetCurrentUser(nil)
	log.Println(resp, err)
}

func TestClient_RegisterByEmail(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	data, e := jsoniter.Marshal([]model.KeyValuePair{{Key: "custom", Value: "qq"}})
	log.Println(data, e)
	p := string(data)
	req := &model.RegisterByEmailInput{
		Email:    "5304950622@qq.com",
		Password: "123456",
		Params:   &p,
	}
	resp, err := authenticationClient.RegisterByEmail(req)
	log.Println(resp, err)
}

func TestClient_RegisterByUsername(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	data, e := jsoniter.Marshal([]model.KeyValuePair{{Key: "custom", Value: "qq"}})
	log.Println(data, e)
	p := string(data)
	req := &model.RegisterByUsernameInput{
		Username: "gosdk",
		Password: "123456",
		Params:   &p,
	}
	resp, err := authenticationClient.RegisterByUsername(req)
	log.Println(resp, err)
}

func TestClient_RegisterByPhoneCode(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	data, e := jsoniter.Marshal([]model.KeyValuePair{{Key: "custom", Value: "qq"}})
	log.Println(data, e)
	p := string(data)
	req := &model.RegisterByPhoneCodeInput{
		Phone:  "15865561492",
		Code:   "123456",
		Params: &p,
	}
	resp, err := authenticationClient.RegisterByPhoneCode(req)
	log.Println(resp, err)
}

func TestClient_CheckPasswordStrength(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	data, e := jsoniter.Marshal([]model.KeyValuePair{{Key: "custom", Value: "qq"}})
	log.Println(data, e)

	resp, err := authenticationClient.CheckPasswordStrength("12345678")
	log.Println(resp, err)
}

func TestClient_SendSmsCode(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool

	resp, err := authenticationClient.SendSmsCode("18910471835")
	log.Println(resp, err)
}

func TestClient_LoginByPhoneCode(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByPhoneCodeInput{
		Code:  "3289",
		Phone: "18910471835",
	}
	resp, err := authenticationClient.LoginByPhoneCode(req)
	log.Println(resp, err)
}

func TestClient_CheckLoginStatus(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	reginter := &model.RegisterByUsernameInput{
		Username: "testGoSDK",
		Password: "123456789",
	}
	ru, re := authenticationClient.RegisterByUsername(reginter)
	log.Println(ru, re)
	req := &model.LoginByUsernameInput{
		Username: "testGoSDK",
		Password: "123456789",
	}
	u, e := authenticationClient.LoginByUserName(*req)
	log.Println(u, e)
	resp, err := authenticationClient.CheckLoginStatus(*u.Token)
	log.Println(resp, err)
}

func TestClient_SendEmail(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool

	resp, err := authenticationClient.SendEmail(" mail@qq.com", model.EnumEmailSceneCHANGE_EMAIL)
	log.Println(resp, err)
}

func TestClient_UpdateProfile(t *testing.T) {
	authenticationClient := NewClient("6139c4d24e78a4d706b7545b", Secret)
	authenticationClient.userPoolId = UserPool

	req := &model.LoginByUsernameInput{
		Username: "updateProfile",
		Password: "123456",
	}
	resp, err := authenticationClient.LoginByUserName(*req)
	log.Println(resp)
	username := "goSdkTestUpdateProfile"
	updateReq := &model.UpdateUserInput{
		Username: &username,
	}
	resp1, err := authenticationClient.UpdateProfile(updateReq)
	log.Println(resp1, err)
}

func TestClient_UpdatePassword(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "goSdkTestUpdateProfile",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.UpdatePassword("123456", "654321")

	log.Println(resp, err)
	loginResp, loginErr := authenticationClient.LoginByUserName(model.LoginByUsernameInput{
		Username: "goSdkTestUpdateProfile",
		Password: "654321",
	})
	log.Println(loginResp, loginErr)
}

func TestClient_UpdatePhone(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "goSdkTestUpdateProfile",
		Password: "654321",
	}
	//authenticationClient.SendSmsCode("18515006338")
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.UpdatePhone("18515006338", "7757", nil, nil)

	log.Println(resp, err)

}

func TestClient_UpdateEmail(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "goSdkTestUpdateProfile",
		Password: "654321",
	}
	//authenticationClient.SendSmsCode("18515006338")
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.UpdateEmail("530495062@qq.com", "7757", nil, nil)

	log.Println(resp, err)

}

func TestClient_RefreshToken(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "goSdkTestUpdateProfile",
		Password: "654321",
	}
	//authenticationClient.SendSmsCode("18515006338")
	user, _ := authenticationClient.LoginByUserName(*req)
	oldToken := user.Token
	log.Println(oldToken)
	resp, err := authenticationClient.RefreshToken(user.Token)
	log.Println(resp.Token)

	log.Println(resp, err)

}

func TestClient_LinkAccount(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "goSdkTestUpdateProfile",
		Password: "654321",
	}

	user, _ := authenticationClient.LoginByUserName(*req)

	resp, err := authenticationClient.LinkAccount(*user.Token, "qqwe")

	log.Println(resp, err)

}

func TestClient_UnLinkAccount(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "goSdkTestUpdateProfile",
		Password: "654321",
	}

	user, _ := authenticationClient.LoginByUserName(*req)

	resp, err := authenticationClient.UnLinkAccount(*user.Token, constant.WECHATPC)

	log.Println(resp, err)

}

func TestClient_BindPhone(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.BindPhone("18515006338", "1453")
	log.Println(resp, err)

}
func TestClient_SendSmsCode2(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	d, e := authenticationClient.SendSmsCode("18515006338")
	log.Println(d, e)
}

func TestClient_UnBindPhone(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.UnBindPhone()
	log.Println(resp, err)

}

func TestClient_BindEmail(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.BindEmail("email", "code")
	log.Println(resp, err)

}

func TestClient_UnBindEmail(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.UnBindEmail()
	log.Println(resp, err)

}

func TestClient_Logout(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.Logout()
	log.Println(resp, err)

}

func TestClient_ListUdv(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.ListUdv()
	log.Println(resp, err)

}

func TestClient_SetUdv(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.SetUdv([]model.KeyValuePair{
		{Key: "age", Value: "18"},
	})
	log.Println(resp, err)

}

func TestClient_RemoveUdv(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.RemoveUdv("school")
	log.Println(resp, err)

}

func TestClient_ListOrg(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.ListOrg()
	log.Println(resp, err)

}

func TestClient_LoginByLdap(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	resp, err := authenticationClient.LoginByLdap("18515006338", "123456")
	log.Println(resp, err)
}

func TestClient_LoginByAd(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	resp, err := authenticationClient.LoginByAd("18515006338", "123456")
	log.Println(resp, err)
}

func TestClient_GetSecurityLevel(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.GetSecurityLevel()
	log.Println(resp, err)
}

func TestClient_ListAuthorizedResources(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.ListAuthorizedResources("default", model.EnumResourceTypeDATA)
	log.Println(resp, err)
}

func TestClient_BuildAuthorizeUrlByOauth(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	authenticationClient.Protocol = constant.OAUTH
	resp, ee := authenticationClient.BuildAuthorizeUrlByOauth("email", "qq", "ww", "cc")
	log.Println(resp, ee)
}

func TestClient_ValidateTicketV1(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	authenticationClient.Protocol = constant.OAUTH
	resp, ee := authenticationClient.ValidateTicketV1("email", "qq")
	log.Println(resp, ee)
}

func TestClient_ListRole(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.ListRole("default")
	log.Println(resp, err)
}
func TestClient_HasRole(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.HasRole("NewCode", "default")
	log.Println(resp, err)
}
func TestClient_ListApplications(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.ListApplications(1, 10)
	log.Println(resp, err)
}

func TestClient_GetCodeChallengeDigest(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool

	resp, err := authenticationClient.GetCodeChallengeDigest("wpaiscposrovkquicztfmftripjocybgmphyqtucmoz", constant.S256)

	log.Println(resp, err)
}

func TestClient_LoginBySubAccount(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginBySubAccountRequest{
		Account:  "123456789",
		Password: "8558781",
	}
	resp, err := authenticationClient.LoginBySubAccount(req)

	log.Println(resp, err)
}

func TestClient_ListDepartments(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.ListDepartments()
	log.Println(resp, err)
}

func TestClient_IsUserExists(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool
	req := &model.LoginByUsernameInput{
		Username: "18515006338",
		Password: "123456",
	}
	userName := "18515006338"
	authenticationClient.LoginByUserName(*req)
	resp, err := authenticationClient.IsUserExists(&model.IsUserExistsRequest{
		Username: &userName,
	})
	log.Println(resp, err)
}

func TestClient_ValidateTicketV2(t *testing.T) {
	authenticationClient := NewClient(AppId, Secret)
	authenticationClient.userPoolId = UserPool

	resp, err := authenticationClient.ValidateTicketV2("ss", "ss", constant.XML)
	log.Println(resp, err)
}
