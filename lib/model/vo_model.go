package model

import (
	"github.com/Authing/authing-go-sdk/lib/enum"
)

type ListMemberRequest struct {
	NodeId               string          `json:"nodeId"`
	Page                 int             `json:"page"`
	Limit                int             `json:"limit"`
	SortBy               enum.SortByEnum `json:"sortBy"`
	IncludeChildrenNodes bool            `json:"includeChildrenNodes"`
}

type UserDetailData struct {
	ThirdPartyIdentity User `json:"thirdPartyIdentity"`
}

type UserDetailResponse struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
	Data    User   `json:"data"`
}

type ExportAllOrganizationResponse struct {
	Message string    `json:"message"`
	Code    int64     `json:"code"`
	Data    []OrgNode `json:"data"`
}

type NodeByIdDetail struct {
	NodeById Node `json:"nodeById"`
}

type NodeByIdResponse struct {
	Data NodeByIdDetail `json:"data"`
}

type QueryListRequest struct {
	Page           int             `json:"page"`
	Limit          int             `json:"limit"`
	SortBy         enum.SortByEnum `json:"sortBy"`
	WithCustomData *bool
}

type Users struct {
	Users PaginatedUsers `json:"users"`
}
type ListUserResponse struct {
	Data Users `json:"data"`
}

/*type OrganizationChildren struct {
	Id	string	`json:"id"`
	CreatedAt *string `json:"createdAt"`
	UpdateAt *string `json:"updateAt"`
	UserPoolId *string `json:"userPoolId"`
	OrgId *string `json:"orgId"`
	Name string `json:"name"`
	Description *string `json:"description"`
	Order *int64 `json:"order"`
	Code *string `json:"code"`
}*/

type ListOrganizationResponse struct {
	Message string        `json:"message"`
	Code    int64         `json:"code"`
	Data    PaginatedOrgs `json:"data"`
}

type GetOrganizationChildrenResponse struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
	Data    []Node `json:"data"`
}

type GetOrganizationByIdData struct {
	Org Org `json:"org"`
}

type GetOrganizationByIdResponse struct {
	Data GetOrganizationByIdData `json:"data"`
}

type ValidateTokenRequest struct {
	AccessToken string `json:"accessToken"`
	IdToken     string `json:"idToken"`
}

type ClientCredentialInput struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type GetAccessTokenByClientCredentialsRequest struct {
	Scope                 string                 `json:"scope"`
	ClientCredentialInput *ClientCredentialInput `json:"client_credential_input"`
}

type OidcParams struct {
	AppId               string
	RedirectUri         string
	ResponseType        string
	ResponseMode        string
	State               string
	Nonce               string
	Scope               string
	CodeChallengeMethod string
	CodeChallenge       string
}

type OrgNode struct {
	Id              string    `json:"id"`
	OrgId           *string   `json:"orgId"`
	CreatedAt       *string   `json:"createdAt"`
	UpdatedAt       *string   `json:"updatedAt"`
	UserPoolId      *string   `json:"userPoolId"`
	Name            string    `json:"name"`
	Description     *string   `json:"description"`
	DescriptionI18n *string   `json:"descriptionI18n"`
	Order           *int64    `json:"order"`
	Code            *string   `json:"code"`
	Members         []User    `json:"members"`
	Children        []OrgNode `json:"children"`
}

type GetUserDepartmentsRequest struct {
	Id    string  `json:"id"`
	OrgId *string `json:"orgId"`
}

type CheckUserExistsRequest struct {
	Email      *string `json:"email"`
	Phone      *string `json:"phone"`
	Username   *string `json:"username"`
	ExternalId *string `json:"externalId"`
}

type CheckUserExistsResponse struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
	Data    bool   `json:"data"`
}

type UserDepartments struct {
	Departments *PaginatedDepartments `json:"departments"`
}

type UserDepartmentsData struct {
	User UserDepartments `json:"user"`
}
type GetUserDepartmentsResponse struct {
	Data UserDepartmentsData `json:"data"`
}

type CommonPageRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ListPoliciesResponse struct {
	TotalCount int `json:"totalCount"`
	List       []struct {
		Code             string `json:"code"`
		TargetType       string `json:"targetType"`
		TargetIdentifier string `json:"targetIdentifier"`
	} `json:"list"`
}

type ListPoliciesRequest struct {
	Code  string `json:"targetIdentifier"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

type ListPoliciesOnIdRequest struct {
	Id    string `json:"targetIdentifier"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
