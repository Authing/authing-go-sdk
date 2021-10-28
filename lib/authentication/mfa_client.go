package authentication

import (
	"errors"
	"fmt"
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/model"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
)

// GetMfaAuthenticators
// 获取 MFA 认证器
func (c *Client) GetMfaAuthenticators(req *model.MfaInput) (*struct {
	Message string                               `json:"message"`
	Code    int64                                `json:"code"`
	Data    []model.GetMfaAuthenticatorsResponse `json:"data"`
}, error) {

	vars := make(map[string]interface{})
	if req.MfaType == nil {
		vars["type"] = "totp"
	} else {
		vars["type"] = req.MfaType
	}
	if req.MfaSource == nil {
		vars["source"] = constant.Self
	} else {
		vars["source"] = req.MfaSource
	}
	url := fmt.Sprintf("%s/api/v2/mfa/authenticator", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodGet, req.MfaToken, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                               `json:"message"`
		Code    int64                                `json:"code"`
		Data    []model.GetMfaAuthenticatorsResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// AssociateMfaAuthenticator
// 请求 MFA 二维码和密钥信息
func (c *Client) AssociateMfaAuthenticator(req *model.MfaInput) (*struct {
	Message string                                  `json:"message"`
	Code    int64                                   `json:"code"`
	Data    model.AssociateMfaAuthenticatorResponse `json:"data"`
}, error) {

	vars := make(map[string]interface{})
	if req.MfaType == nil {
		vars["authenticatorType"] = "totp"
	} else {
		vars["authenticatorType"] = req.MfaType
	}
	if req.MfaSource == nil {
		vars["source"] = constant.Self
	} else {
		vars["source"] = req.MfaSource
	}
	url := fmt.Sprintf("%s/api/v2/mfa/totp/associate", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, req.MfaToken, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string                                  `json:"message"`
		Code    int64                                   `json:"code"`
		Data    model.AssociateMfaAuthenticatorResponse `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// DeleteMfaAuthenticator
// 解绑 MFA
func (c *Client) DeleteMfaAuthenticator() (*model.CommonMessageAndCode, error) {

	url := fmt.Sprintf("%s/api/v2/mfa/totp/associate", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodDelete, nil, nil)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	var resp model.CommonMessageAndCode
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return &resp, nil
}

// ConfirmAssociateMfaAuthenticator
// 确认绑定 MFA
func (c *Client) ConfirmAssociateMfaAuthenticator(req *model.ConfirmAssociateMfaAuthenticatorRequest) (*struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
}, error) {

	vars := make(map[string]interface{})
	if req.AuthenticatorType == nil {
		vars["authenticatorType"] = "totp"
	} else {
		vars["authenticatorType"] = req.AuthenticatorType
	}
	if req.MfaSource == nil {
		vars["source"] = constant.Self
	} else {
		vars["source"] = req.MfaSource
	}
	vars["totp"] = req.Totp
	url := fmt.Sprintf("%s/api/v2/mfa/totp/associate/confirm", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, req.MfaToken, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string      `json:"message"`
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// VerifyTotpMfa
// 检验二次验证 MFA 口令
func (c *Client) VerifyTotpMfa(totp, token string) (*struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
}, error) {

	vars := make(map[string]interface{})

	vars["totp"] = totp
	url := fmt.Sprintf("%s/api/v2/mfa/totp/verify", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, &token, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string      `json:"message"`
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// VerifyAppSmsMfa
// 检验二次验证 MFA 短信验证码
func (c *Client) VerifyAppSmsMfa(phone, code, token string) (*struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
}, error) {

	vars := map[string]interface{}{
		"code":  code,
		"phone": phone,
	}

	url := fmt.Sprintf("%s/api/v2/applications/mfa/sms/verify", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, &token, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string      `json:"message"`
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// VerifyAppEmailMfa
// 检验二次验证 MFA 邮箱验证码
func (c *Client) VerifyAppEmailMfa(email, code, token string) (*struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
}, error) {

	vars := map[string]interface{}{
		"code":  code,
		"email": email,
	}

	url := fmt.Sprintf("%s/api/v2/applications/mfa/email/verify", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, &token, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string      `json:"message"`
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// PhoneOrEmailBindable
// 检测手机号或邮箱是否已被绑定
func (c *Client) PhoneOrEmailBindable(email, phone *string, token string) (*struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
}, error) {

	vars := make(map[string]interface{})
	if email != nil {
		vars["email"] = email
	}
	if phone != nil {
		vars["phone"] = phone
	}

	url := fmt.Sprintf("%s/api/v2/applications/mfa/check", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, &token, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string      `json:"message"`
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// VerifyTotpRecoveryCode
// 检验二次验证 MFA 恢复代码
func (c *Client) VerifyTotpRecoveryCode(code, token string) (*struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
}, error) {

	vars := make(map[string]interface{})

	vars["recoveryCode"] = code
	url := fmt.Sprintf("%s/api/v2/mfa/totp/recovery", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, &token, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string      `json:"message"`
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// AssociateFaceByUrl
// 通过图片 URL 绑定人脸
func (c *Client) AssociateFaceByUrl(baseFaceUrl, CompareFaceUrl, token string) (*struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
}, error) {

	vars := map[string]interface{}{
		"photoA":     baseFaceUrl,
		"photoB":     CompareFaceUrl,
		"isExternal": true,
	}
	url := fmt.Sprintf("%s/api/v2/mfa/face/associate", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, &token, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string      `json:"message"`
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}

// VerifyFaceMfa
// 人脸二次认证
func (c *Client) VerifyFaceMfa(faceUrl, token string) (*struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
}, error) {

	vars := map[string]interface{}{
		"photo": faceUrl,
	}
	url := fmt.Sprintf("%s/api/v2/mfa/face/associate", c.Host)
	b, err := c.SendHttpRestRequest(url, http.MethodPost, &token, vars)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	resp := &struct {
		Message string      `json:"message"`
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
	}{}
	jsoniter.Unmarshal(b, &resp)
	if resp.Code != 200 {
		return nil, errors.New(resp.Message)
	}
	return resp, nil
}
