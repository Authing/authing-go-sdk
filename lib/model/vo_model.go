package model

import (
	"github.com/Authing/authing-go-sdk/lib/constant"
	"github.com/Authing/authing-go-sdk/lib/enum"
	"time"
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

type ListAuthorizedResourcesByIdRequest struct {
	Id           string  `json:"id"`
	Namespace    string  `json:"namespace,omitempty"`
	ResourceType *string `json:"resourceType"`
}

type ListUserAuthorizedResourcesResponse struct {
	User User `json:"user"`
}

type IsAllowedRequest struct {
	Resource  string  `json:"resource"`
	Action    string  `json:"action"`
	UserId    string  `json:"userId"`
	Namespace *string `json:"namespace"`
}

type AllowRequest struct {
	Resource  string `json:"resource"`
	Action    string `json:"action"`
	UserId    string `json:"userId"`
	Namespace string `json:"namespace"`
}

type AuthorizeResourceRequest struct {
	Namespace    string                 `json:"namespace"`
	Resource     string                 `json:"resource"`
	ResourceType EnumResourceType       `json:"resourceType"`
	Opts         []AuthorizeResourceOpt `json:"opts"`
}

type RevokeResourceRequest struct {
	Namespace    string                 `json:"namespace"`
	Resource     string                 `json:"resource"`
	ResourceType EnumResourceType       `json:"resourceType"`
	Opts         []AuthorizeResourceOpt `json:"opts"`
}

type GetUserRoleListRequest struct {
	UserId    string  `json:"userId"`
	Namespace *string `json:"namespace"`
}

type CheckResourcePermissionBatchRequest struct {
	UserId    string   `json:"userId"`
	Namespace string   `json:"namespace"`
	Resources []string `json:"resources"`
}

type GetAuthorizedResourcesOfResourceKindRequest struct {
	UserId    string `json:"userId"`
	Namespace string `json:"namespace"`
	Resource  string `json:"resource"`
}

type ListAuthorizedResourcesRequest struct {
	TargetIdentifier string                          `json:"targetIdentifier"`
	Namespace        string                          `json:"namespace"`
	TargetType       constant.ResourceTargetTypeEnum `json:"targetType"`
	ResourceType     *EnumResourceType               `json:"resourceType"`
}

type ProgrammaticAccessAccount struct {
	AppId         string    `json:"appId"`
	Secret        string    `json:"secret"`
	TokenLifetime int       `json:"tokenLifetime"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Id            string    `json:"id"`
	Remarks       string    `json:"remarks"`
	UserId        string    `json:"userId"`
	Enabled       bool      `json:"enabled"`
}

type ListResourceRequest struct {
	Namespace    string           `json:"namespace"`
	ResourceType EnumResourceType `json:"resourceType,omitempty"`
	Page         int              `json:"page"`
	Limit        int              `json:"limit"`
}
type ActionsModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Resource struct {
	Id            string         `json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	UserPoolId    string         `json:"userPoolId"`
	Code          string         `json:"code"`
	Actions       []ActionsModel `json:"actions"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	NamespaceId   int            `json:"namespaceId"`
	ApiIdentifier *string        `json:"apiIdentifier"`
	Namespace     string         `json:"namespace,omitempty"`
}
type ResourceResponse struct {
	Id            string         `json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	UserPoolId    string         `json:"userPoolId"`
	Code          string         `json:"code"`
	Actions       []ActionsModel `json:"actions"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	NamespaceId   int            `json:"namespaceId"`
	ApiIdentifier *string        `json:"apiIdentifier"`
}

type ListNamespaceResourceResponse struct {
	List       []Resource `json:"list"`
	TotalCount int        `json:"totalCount"`
}

type CreateResourceRequest struct {
	Code          string         `json:"code"`
	Actions       []ActionsModel `json:"actions,omitempty"`
	Type          string         `json:"type,omitempty"`
	Description   *string        `json:"description,omitempty"`
	ApiIdentifier *string        `json:"apiIdentifier,omitempty"`
	Namespace     string         `json:"namespace,omitempty"`
}

type UpdateResourceRequest struct {
	Actions       []ActionsModel `json:"actions,omitempty"`
	Type          string         `json:"type,omitempty"`
	Description   *string        `json:"description,omitempty"`
	ApiIdentifier *string        `json:"apiIdentifier,omitempty"`
	Namespace     string         `json:"namespace,omitempty"`
}

type ApplicationAccessPolicies struct {
	AssignedAt        time.Time   `json:"assignedAt"`
	InheritByChildren interface{} `json:"inheritByChildren"`
	Enabled           bool        `json:"enabled"`
	PolicyId          string      `json:"policyId"`
	Code              string      `json:"code"`
	Policy            struct {
		Id          string    `json:"id"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		UserPoolId  string    `json:"userPoolId"`
		IsDefault   bool      `json:"isDefault"`
		IsAuto      bool      `json:"isAuto"`
		Hidden      bool      `json:"hidden"`
		Code        string    `json:"code"`
		Description string    `json:"description"`
		Statements  []struct {
			Resource     string           `json:"resource"`
			Actions      []string         `json:"actions"`
			Effect       string           `json:"effect"`
			Condition    []interface{}    `json:"condition"`
			ResourceType EnumResourceType `json:"resourceType"`
		} `json:"statements"`
		NamespaceId int `json:"namespaceId"`
	} `json:"policy"`
	TargetNamespace  string `json:"targetNamespace"`
	TargetType       string `json:"targetType"`
	TargetIdentifier string `json:"targetIdentifier"`
	Target           struct {
		Id          string    `json:"id"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		UserPoolId  string    `json:"userPoolId"`
		Code        string    `json:"code"`
		Description string    `json:"description"`
		ParentCode  string    `json:"parentCode"`
		IsSystem    bool      `json:"isSystem"`
		NamespaceId int       `json:"namespaceId"`
	} `json:"target"`
	Namespace string `json:"namespace"`
}

type GetApplicationAccessPoliciesResponse struct {
	List       []ApplicationAccessPolicies `json:"list"`
	TotalCount int                         `json:"totalCount"`
}

type ApplicationAccessPoliciesRequest struct {
	TargetIdentifiers []string                        `json:"targetIdentifiers,omitempty"`
	TargetType        constant.ResourceTargetTypeEnum `json:"targetType,omitempty"`
	Namespace         string                          `json:"namespace,omitempty"`
	InheritByChildren bool                            `json:"inheritByChildren,omitempty"`
}

type GetAuthorizedTargetsRequest struct {
	TargetType   constant.ResourceTargetTypeEnum `json:"targetType"`
	Namespace    string                          `json:"namespace"`
	Resource     string                          `json:"resource"`
	ResourceType EnumResourceType                `json:"resourceType"`
	Actions      *struct {
		Op   constant.GetAuthorizedTargetsOpt `json:"op,omitempty"`
		List []string                         `json:"list,omitempty"`
	} `json:"actions,omitempty"`
}

type ListAuditLogsRequest struct {
	ClientIp       *string   `json:"clientip,omitempty"`
	OperationNames *[]string `json:"operation_name,omitempty"`
	UserIds        *[]string `json:"operator_arn,omitempty"`
	AppIds         *[]string `json:"app_id,omitempty"`
	Page           *int      `json:"page,omitempty"`
	Limit          *int      `json:"limit,omitempty"`
}

type ListUserActionRequest struct {
	ClientIp       *string   `json:"clientip,omitempty"`
	OperationNames *[]string `json:"operation_name,omitempty"`
	UserIds        *[]string `json:"operator_arn,omitempty"`
	Page           *int      `json:"page,omitempty"`
	Limit          *int      `json:"limit,omitempty"`
}

type CheckLoginStatusResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Exp     int    `json:"exp"`
	Iat     int    `json:"iat"`
	Data    struct {
		Id         string `json:"id"`
		UserPoolId string `json:"userPoolId"`
		Arn        string `json:"arn"`
	} `json:"data"`
}

type SetUdfInput struct {
	TargetType EnumUDFTargetType `json:"targetType"`
	Key        string            `json:"key"`
	DataType   EnumUDFDataType   `json:"dataType"`
	Label      string            `json:"label"`
}

type PrincipalAuthenticateRequest struct {
	Type   constant.PrincipalAuthenticateType `json:"type"`
	Name   string                             `json:"name"`
	IdCard string                             `json:"idCard"`
	Ext    string                             `json:"ext"`
}
