package model

import (
	"github.com/Authing/authing-go-sdk/lib/constant"
	"time"
)

type MfaInput struct {
	MfaToken  *string
	MfaType   *string             `json:"type"`
	MfaSource *constant.MfaSource `json:"source"`
}

type GetMfaAuthenticatorsResponse struct {
	Id                string    `json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	UserId            string    `json:"userId"`
	Enable            bool      `json:"enable"`
	Secret            string    `json:"secret"`
	AuthenticatorType string    `json:"authenticatorType"`
	RecoveryCode      string    `json:"recoveryCode"`
	Source            string    `json:"source"`
}

type AssociateMfaAuthenticatorResponse struct {
	AuthenticatorType string `json:"authenticator_type"`
	Secret            string `json:"secret"`
	QrcodeUri         string `json:"qrcode_uri"`
	QrcodeDataUrl     string `json:"qrcode_data_url"`
	RecoveryCode      string `json:"recovery_code"`
}

type ConfirmAssociateMfaAuthenticatorRequest struct {
	Totp              string              `json:"totp"`
	AuthenticatorType *string             `json:"authenticatorType"`
	MfaSource         *constant.MfaSource `json:"source"`
	MfaToken          *string
}
