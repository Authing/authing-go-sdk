package authentication

import (
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"testing"
)

func TestClient_GetAccessTokenByCode(t *testing.T) {
	authenticationClient := NewClient("60a6f980dd9a9a7642da768a","5cd4ea7b3603b792aea9a00da9e18f44")
	authenticationClient.Host = "https://32l5hb-demo.authing.cn"
	authenticationClient.RedirectUri = "https://mvnrepository.com/"
	authenticationClient.Protocol = constant.OIDC
	authenticationClient.TokenEndPointAuthMethod = constant.None
	resp, err := authenticationClient.GetAccessTokenByCode("NukTfPbT2-N-AkWApQeYhxItgJzRmfECz_KmDivETE9")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}