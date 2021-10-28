// Code generated by go generate; DO NOT EDIT.
// This file was generated from GraphQL schema

package model

import "time"

type EnumEmailTemplateType string

const EnumEmailTemplateTypeRESET_PASSWORD EnumEmailTemplateType = "RESET_PASSWORD"
const EnumEmailTemplateTypePASSWORD_RESETED_NOTIFICATION EnumEmailTemplateType = "PASSWORD_RESETED_NOTIFICATION"
const EnumEmailTemplateTypeCHANGE_PASSWORD EnumEmailTemplateType = "CHANGE_PASSWORD"
const EnumEmailTemplateTypeWELCOME EnumEmailTemplateType = "WELCOME"
const EnumEmailTemplateTypeVERIFY_EMAIL EnumEmailTemplateType = "VERIFY_EMAIL"
const EnumEmailTemplateTypeCHANGE_EMAIL EnumEmailTemplateType = "CHANGE_EMAIL"

type EnumResourceType string

const EnumResourceTypeDATA EnumResourceType = "DATA"
const EnumResourceTypeAPI EnumResourceType = "API"
const EnumResourceTypeMENU EnumResourceType = "MENU"
const EnumResourceTypeUI EnumResourceType = "UI"
const EnumResourceTypeBUTTON EnumResourceType = "BUTTON"

type EnumSortByEnum string

const EnumSortByEnumCREATEDAT_DESC EnumSortByEnum = "CREATEDAT_DESC"
const EnumSortByEnumCREATEDAT_ASC EnumSortByEnum = "CREATEDAT_ASC"
const EnumSortByEnumUPDATEDAT_DESC EnumSortByEnum = "UPDATEDAT_DESC"
const EnumSortByEnumUPDATEDAT_ASC EnumSortByEnum = "UPDATEDAT_ASC"

type EnumUserStatus string

const EnumUserStatusSuspended EnumUserStatus = "Suspended"
const EnumUserStatusResigned EnumUserStatus = "Resigned"
const EnumUserStatusActivated EnumUserStatus = "Activated"
const EnumUserStatusArchived EnumUserStatus = "Archived"

type Enum__TypeKind string

const Enum__TypeKindSCALAR Enum__TypeKind = "SCALAR"
const Enum__TypeKindOBJECT Enum__TypeKind = "OBJECT"
const Enum__TypeKindINTERFACE Enum__TypeKind = "INTERFACE"
const Enum__TypeKindUNION Enum__TypeKind = "UNION"
const Enum__TypeKindENUM Enum__TypeKind = "ENUM"
const Enum__TypeKindINPUT_OBJECT Enum__TypeKind = "INPUT_OBJECT"
const Enum__TypeKindLIST Enum__TypeKind = "LIST"
const Enum__TypeKindNON_NULL Enum__TypeKind = "NON_NULL"

type EnumEmailScene string

const EnumEmailSceneRESET_PASSWORD EnumEmailScene = "RESET_PASSWORD"
const EnumEmailSceneVERIFY_EMAIL EnumEmailScene = "VERIFY_EMAIL"
const EnumEmailSceneCHANGE_EMAIL EnumEmailScene = "CHANGE_EMAIL"
const EnumEmailSceneMFA_VERIFY EnumEmailScene = "MFA_VERIFY"

type EnumOperator string

const EnumOperatorAND EnumOperator = "AND"
const EnumOperatorOR EnumOperator = "OR"

type EnumPolicyAssignmentTargetType string

const EnumPolicyAssignmentTargetTypeUSER EnumPolicyAssignmentTargetType = "USER"
const EnumPolicyAssignmentTargetTypeROLE EnumPolicyAssignmentTargetType = "ROLE"
const EnumPolicyAssignmentTargetTypeGROUP EnumPolicyAssignmentTargetType = "GROUP"
const EnumPolicyAssignmentTargetTypeORG EnumPolicyAssignmentTargetType = "ORG"
const EnumPolicyAssignmentTargetTypeAK_SK EnumPolicyAssignmentTargetType = "AK_SK"

type EnumPolicyEffect string

const EnumPolicyEffectALLOW EnumPolicyEffect = "ALLOW"
const EnumPolicyEffectDENY EnumPolicyEffect = "DENY"

type EnumUDFDataType string

const EnumUDFDataTypeSTRING EnumUDFDataType = "STRING"
const EnumUDFDataTypeNUMBER EnumUDFDataType = "NUMBER"
const EnumUDFDataTypeDATETIME EnumUDFDataType = "DATETIME"
const EnumUDFDataTypeBOOLEAN EnumUDFDataType = "BOOLEAN"
const EnumUDFDataTypeOBJECT EnumUDFDataType = "OBJECT"

type EnumUDFTargetType string

const EnumUDFTargetTypeNODE EnumUDFTargetType = "NODE"
const EnumUDFTargetTypeORG EnumUDFTargetType = "ORG"
const EnumUDFTargetTypeUSER EnumUDFTargetType = "USER"
const EnumUDFTargetTypeUSERPOOL EnumUDFTargetType = "USERPOOL"
const EnumUDFTargetTypeROLE EnumUDFTargetType = "ROLE"
const EnumUDFTargetTypePERMISSION EnumUDFTargetType = "PERMISSION"
const EnumUDFTargetTypeAPPLICATION EnumUDFTargetType = "APPLICATION"

type EnumWhitelistType string

const EnumWhitelistTypeUSERNAME EnumWhitelistType = "USERNAME"
const EnumWhitelistTypeEMAIL EnumWhitelistType = "EMAIL"
const EnumWhitelistTypePHONE EnumWhitelistType = "PHONE"

type Enum__DirectiveLocation string

const Enum__DirectiveLocationQUERY Enum__DirectiveLocation = "QUERY"
const Enum__DirectiveLocationMUTATION Enum__DirectiveLocation = "MUTATION"
const Enum__DirectiveLocationSUBSCRIPTION Enum__DirectiveLocation = "SUBSCRIPTION"
const Enum__DirectiveLocationFIELD Enum__DirectiveLocation = "FIELD"
const Enum__DirectiveLocationFRAGMENT_DEFINITION Enum__DirectiveLocation = "FRAGMENT_DEFINITION"
const Enum__DirectiveLocationFRAGMENT_SPREAD Enum__DirectiveLocation = "FRAGMENT_SPREAD"
const Enum__DirectiveLocationINLINE_FRAGMENT Enum__DirectiveLocation = "INLINE_FRAGMENT"
const Enum__DirectiveLocationVARIABLE_DEFINITION Enum__DirectiveLocation = "VARIABLE_DEFINITION"
const Enum__DirectiveLocationSCHEMA Enum__DirectiveLocation = "SCHEMA"
const Enum__DirectiveLocationSCALAR Enum__DirectiveLocation = "SCALAR"
const Enum__DirectiveLocationOBJECT Enum__DirectiveLocation = "OBJECT"
const Enum__DirectiveLocationFIELD_DEFINITION Enum__DirectiveLocation = "FIELD_DEFINITION"
const Enum__DirectiveLocationARGUMENT_DEFINITION Enum__DirectiveLocation = "ARGUMENT_DEFINITION"
const Enum__DirectiveLocationINTERFACE Enum__DirectiveLocation = "INTERFACE"
const Enum__DirectiveLocationUNION Enum__DirectiveLocation = "UNION"
const Enum__DirectiveLocationENUM Enum__DirectiveLocation = "ENUM"
const Enum__DirectiveLocationENUM_VALUE Enum__DirectiveLocation = "ENUM_VALUE"
const Enum__DirectiveLocationINPUT_OBJECT Enum__DirectiveLocation = "INPUT_OBJECT"
const Enum__DirectiveLocationINPUT_FIELD_DEFINITION Enum__DirectiveLocation = "INPUT_FIELD_DEFINITION"

type __Schema struct {
	Types            []__Type      `json:"types"`
	QueryType        __Type        `json:"queryType"`
	MutationType     *__Type       `json:"mutationType"`
	SubscriptionType *__Type       `json:"subscriptionType"`
	Directives       []__Directive `json:"directives"`
}

type __Type struct {
	Kind          Enum__TypeKind `json:"kind"`
	Name          *string        `json:"name"`
	Description   *string        `json:"description"`
	Fields        []__Field      `json:"fields"`
	Interfaces    []__Type       `json:"interfaces"`
	PossibleTypes []__Type       `json:"possibleTypes"`
	EnumValues    []__EnumValue  `json:"enumValues"`
	InputFields   []__InputValue `json:"inputFields"`
	OfType        *__Type        `json:"ofType"`
}

type __Field struct {
	Name              string         `json:"name"`
	Description       *string        `json:"description"`
	Args              []__InputValue `json:"args"`
	Type              __Type         `json:"type"`
	IsDeprecated      bool           `json:"isDeprecated"`
	DeprecationReason *string        `json:"deprecationReason"`
}

type __InputValue struct {
	Name         string  `json:"name"`
	Description  *string `json:"description"`
	Type         __Type  `json:"type"`
	DefaultValue *string `json:"defaultValue"`
}

type __EnumValue struct {
	Name              string  `json:"name"`
	Description       *string `json:"description"`
	IsDeprecated      bool    `json:"isDeprecated"`
	DeprecationReason *string `json:"deprecationReason"`
}

type __Directive struct {
	Name         string                    `json:"name"`
	Description  *string                   `json:"description"`
	Locations    []Enum__DirectiveLocation `json:"locations"`
	Args         []__InputValue            `json:"args"`
	IsRepeatable bool                      `json:"isRepeatable"`
}

type AccessTokenRes struct {
	AccessToken *string `json:"accessToken"`
	Exp         *int64  `json:"exp"`
	Iat         *int64  `json:"iat"`
}

type App2WxappLoginStrategy struct {
	TicketExpriresAfter              *int64 `json:"ticketExpriresAfter"`
	TicketExchangeUserInfoNeedSecret *bool  `json:"ticketExchangeUserInfoNeedSecret"`
}

type App2WxappLoginStrategyInput struct {
	TicketExpriresAfter              *int64 `json:"ticketExpriresAfter,omitempty"`
	TicketExchangeUserInfoNeedSecret *bool  `json:"ticketExchangeUserInfoNeedSecret,omitempty"`
}

type AuthorizedResource struct {
	Code    string            `json:"code"`
	Type    *EnumResourceType `json:"type"`
	Actions []string          `json:"actions"`
}

type AuthorizedTargetsActionsInput struct {
	Op   EnumOperator `json:"op"`
	List []*string    `json:"list"`
}

type AuthorizeResourceOpt struct {
	TargetType       EnumPolicyAssignmentTargetType `json:"targetType"`
	TargetIdentifier string                         `json:"targetIdentifier"`
	Actions          []string                       `json:"actions"`
}

type BatchOperationResult struct {
	SucceedCount int64    `json:"succeedCount"`
	FailedCount  int64    `json:"failedCount"`
	Message      *string  `json:"message"`
	Errors       []string `json:"errors"`
}

type ChangeEmailStrategy struct {
	VerifyOldEmail *bool `json:"verifyOldEmail,omitempty"`
}

type ChangeEmailStrategyInput struct {
	VerifyOldEmail *bool `json:"verifyOldEmail,omitempty"`
}

type ChangePhoneStrategy struct {
	VerifyOldPhone *bool `json:"verifyOldPhone,omitempty"`
}

type ChangePhoneStrategyInput struct {
	VerifyOldPhone *bool `json:"verifyOldPhone,omitempty"`
}

type CheckPasswordStrengthResult struct {
	Valid   bool    `json:"valid"`
	Message *string `json:"message"`
}

type CommonMessage struct {
	Message *string `json:"message"`
	Code    *int64  `json:"code"`
}

type ConfigEmailTemplateInput struct {
	Type       EnumEmailTemplateType `json:"type"`
	Name       string                `json:"name"`
	Subject    string                `json:"subject"`
	Sender     string                `json:"sender"`
	Content    string                `json:"content"`
	RedirectTo *string               `json:"redirectTo"`
	HasURL     *bool                 `json:"hasURL"`
	ExpiresIn  *int64                `json:"expiresIn"`
}

type CreateFunctionInput struct {
	Name        string  `json:"name"`
	SourceCode  string  `json:"sourceCode"`
	Description *string `json:"description"`
	Url         *string `json:"url"`
}

type CreateSocialConnectionInput struct {
	Provider    string                       `json:"provider"`
	Name        string                       `json:"name"`
	Logo        string                       `json:"logo"`
	Description *string                      `json:"description"`
	Fields      []SocialConnectionFieldInput `json:"fields"`
}

type CreateSocialConnectionInstanceFieldInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateSocialConnectionInstanceInput struct {
	Provider string                                      `json:"provider"`
	Fields   []*CreateSocialConnectionInstanceFieldInput `json:"fields"`
}

type CustomSMSProvider struct {
	Enabled  *bool   `json:"enabled"`
	Provider *string `json:"provider"`
	Config   *string `json:"config"`
}

type CustomSMSProviderInput struct {
	Enabled  *bool   `json:"enabled,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Config   *string `json:"config,omitempty"`
}

type EmailTemplate struct {
	Type       EnumEmailTemplateType `json:"type"`
	Name       string                `json:"name"`
	Subject    string                `json:"subject"`
	Sender     string                `json:"sender"`
	Content    string                `json:"content"`
	RedirectTo *string               `json:"redirectTo"`
	HasURL     *bool                 `json:"hasURL"`
	ExpiresIn  *int64                `json:"expiresIn"`
	Enabled    *bool                 `json:"enabled"`
	IsSystem   *bool                 `json:"isSystem"`
}

type FrequentRegisterCheckConfig struct {
	TimeInterval *int64 `json:"timeInterval"`
	Limit        *int64 `json:"limit"`
	Enabled      *bool  `json:"enabled"`
}

type FrequentRegisterCheckConfigInput struct {
	TimeInterval *int64 `json:"timeInterval,omitempty"`
	Limit        *int64 `json:"limit,omitempty"`
	Enabled      *bool  `json:"enabled,omitempty"`
}

type Function struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	SourceCode  string  `json:"sourceCode"`
	Description *string `json:"description"`
	Url         *string `json:"url"`
}

type Group struct {
	Code                string                        `json:"code"`
	Name                string                        `json:"name"`
	Description         *string                       `json:"description"`
	CreatedAt           *string                       `json:"createdAt"`
	UpdatedAt           *string                       `json:"updatedAt"`
	Users               PaginatedUsers                `json:"users"`
	AuthorizedResources *PaginatedAuthorizedResources `json:"authorizedResources"`
}

type Identity struct {
	Openid       *string `json:"openid"`
	UserIdInIdp  *string `json:"userIdInIdp"`
	UserId       *string `json:"userId"`
	ConnectionId *string `json:"connectionId"`
	IsSocial     *bool   `json:"isSocial"`
	Provider     *string `json:"provider"`
	UserPoolId   *string `json:"userPoolId"`
	RefreshToken *string `json:"refreshToken"`
	AccessToken  *string `json:"accessToken"`
}

type JWTTokenStatus struct {
	Code    *int64                `json:"code"`
	Message *string               `json:"message"`
	Status  *bool                 `json:"status"`
	Exp     *int64                `json:"exp"`
	Iat     *int64                `json:"iat"`
	Data    *JWTTokenStatusDetail `json:"data"`
}

type JWTTokenStatusDetail struct {
	Id         *string `json:"id"`
	UserPoolId *string `json:"userPoolId"`
	Arn        *string `json:"arn"`
}

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type LoginByEmailInput struct {
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	CaptchaCode  *string `json:"captchaCode"`
	AutoRegister *bool   `json:"autoRegister"`
	ClientIp     *string `json:"clientIp"`
	Params       *string `json:"params"`
	Context      *string `json:"context"`
}

type LoginByPhoneCodeInput struct {
	Phone        string  `json:"phone"`
	Code         string  `json:"code"`
	AutoRegister *bool   `json:"autoRegister"`
	ClientIp     *string `json:"clientIp"`
	Params       *string `json:"params"`
	Context      *string `json:"context"`
}

type LoginByPhonePasswordInput struct {
	Phone        string  `json:"phone"`
	Password     string  `json:"password"`
	CaptchaCode  *string `json:"captchaCode"`
	AutoRegister *bool   `json:"autoRegister"`
	ClientIp     *string `json:"clientIp"`
	Params       *string `json:"params"`
	Context      *string `json:"context"`
}

type LoginByUsernameInput struct {
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	CaptchaCode  *string `json:"captchaCode"`
	AutoRegister *bool   `json:"autoRegister"`
	ClientIp     *string `json:"clientIp"`
	Params       *string `json:"params"`
	Context      *string `json:"context"`
}

type LoginFailCheckConfig struct {
	TimeInterval *int64 `json:"timeInterval"`
	Limit        *int64 `json:"limit"`
	Enabled      *bool  `json:"enabled"`
}

type LoginFailCheckConfigInput struct {
	TimeInterval *int64 `json:"timeInterval,omitempty"`
	Limit        *int64 `json:"limit,omitempty"`
	Enabled      *bool  `json:"enabled,omitempty"`
}

type LoginPasswordFailCheckConfig struct {
	TimeInterval *int64 `json:"timeInterval,omitempty"`
	Limit        *int64 `json:"limit,omitempty"`
	Enabled      *bool  `json:"enabled,omitempty"`
}

type LoginPasswordFailCheckConfigInput struct {
	TimeInterval *int64 `json:"timeInterval,omitempty"`
	Limit        *int64 `json:"limit,omitempty"`
	Enabled      *bool  `json:"enabled,omitempty"`
}

type Mfa struct {
	Id         string  `json:"id"`
	UserId     string  `json:"userId"`
	UserPoolId string  `json:"userPoolId"`
	Enable     bool    `json:"enable"`
	Secret     *string `json:"secret"`
}

type PaginatedAuthorizedResources struct {
	TotalCount int64                `json:"totalCount"`
	List       []AuthorizedResource `json:"list"`
}

type PaginatedAuthorizedTargets struct {
	List       []*ResourcePermissionAssignment `json:"list"`
	TotalCount *int64                          `json:"totalCount"`
}

type PaginatedDepartments struct {
	List       []UserDepartment `json:"list"`
	TotalCount int64            `json:"totalCount"`
}

type PaginatedFunctions struct {
	List       []Function `json:"list"`
	TotalCount int64      `json:"totalCount"`
}

type PaginatedGroups struct {
	TotalCount int64   `json:"totalCount"`
	List       []Group `json:"list"`
}

type PaginatedRoles struct {
	TotalCount int64  `json:"totalCount"`
	List       []Role `json:"list"`
}

type PaginatedUserpool struct {
	TotalCount int64      `json:"totalCount"`
	List       []UserPool `json:"list"`
}

type PaginatedUsers struct {
	TotalCount int64  `json:"totalCount"`
	List       []User `json:"list"`
}

type QrcodeLoginStrategy struct {
	QrcodeExpiresAfter               *int64 `json:"qrcodeExpiresAfter"`
	ReturnFullUserInfo               *bool  `json:"returnFullUserInfo"`
	AllowExchangeUserInfoFromBrowser *bool  `json:"allowExchangeUserInfoFromBrowser"`
	TicketExpiresAfter               *int64 `json:"ticketExpiresAfter"`
}

type QrcodeLoginStrategyInput struct {
	QrcodeExpiresAfter               *int64 `json:"qrcodeExpiresAfter,omitempty"`
	ReturnFullUserInfo               *bool  `json:"returnFullUserInfo,omitempty"`
	AllowExchangeUserInfoFromBrowser *bool  `json:"allowExchangeUserInfoFromBrowser,omitempty"`
	TicketExpiresAfter               *int64 `json:"ticketExpiresAfter,omitempty"`
}

type RefreshAccessTokenRes struct {
	AccessToken *string `json:"accessToken"`
	Exp         *int64  `json:"exp"`
	Iat         *int64  `json:"iat"`
}

type RefreshToken struct {
	Token *string `json:"token"`
	Iat   *int64  `json:"iat"`
	Exp   *int64  `json:"exp"`
}

type RegisterByEmailInput struct {
	Email         string           `json:"email"`
	Password      string           `json:"password"`
	Profile       *RegisterProfile `json:"profile,omitempty"`
	ForceLogin    *bool            `json:"forceLogin,omitempty"`
	GenerateToken *bool            `json:"generateToken,omitempty"`
	ClientIp      *string          `json:"clientIp,omitempty"`
	Params        *string          `json:"params,omitempty"`
	Context       *string          `json:"context,omitempty"`
}

type RegisterByPhoneCodeInput struct {
	Phone         string           `json:"phone"`
	Code          string           `json:"code"`
	Password      *string          `json:"password"`
	Profile       *RegisterProfile `json:"profile"`
	ForceLogin    *bool            `json:"forceLogin"`
	GenerateToken *bool            `json:"generateToken"`
	ClientIp      *string          `json:"clientIp"`
	Params        *string          `json:"params"`
	Context       *string          `json:"context"`
}

type RegisterByUsernameInput struct {
	Username      string           `json:"username"`
	Password      string           `json:"password"`
	Profile       *RegisterProfile `json:"profile"`
	ForceLogin    *bool            `json:"forceLogin"`
	GenerateToken *bool            `json:"generateToken"`
	ClientIp      *string          `json:"clientIp"`
	Params        *string          `json:"params"`
	Context       *string          `json:"context"`
}

type RegisterProfile struct {
	Ip                *string        `json:"ip"`
	Oauth             *string        `json:"oauth"`
	Username          *string        `json:"username"`
	Nickname          *string        `json:"nickname"`
	Company           *string        `json:"company"`
	Photo             *string        `json:"photo"`
	Device            *string        `json:"device"`
	Browser           *string        `json:"browser"`
	Name              *string        `json:"name"`
	GivenName         *string        `json:"givenName"`
	FamilyName        *string        `json:"familyName"`
	MiddleName        *string        `json:"middleName"`
	Profile           *string        `json:"profile"`
	PreferredUsername *string        `json:"preferredUsername"`
	Website           *string        `json:"website"`
	Gender            *string        `json:"gender"`
	Birthdate         *string        `json:"birthdate"`
	Zoneinfo          *string        `json:"zoneinfo"`
	Locale            *string        `json:"locale"`
	Address           *string        `json:"address"`
	Formatted         *string        `json:"formatted"`
	StreetAddress     *string        `json:"streetAddress"`
	Locality          *string        `json:"locality"`
	Region            *string        `json:"region"`
	PostalCode        *string        `json:"postalCode"`
	Country           *string        `json:"country"`
	Udf               []UserDdfInput `json:"udf"`
}

type RegisterWhiteListConfig struct {
	PhoneEnabled    *bool `json:"phoneEnabled"`
	EmailEnabled    *bool `json:"emailEnabled"`
	UsernameEnabled *bool `json:"usernameEnabled"`
}

type RegisterWhiteListConfigInput struct {
	PhoneEnabled    *bool `json:"phoneEnabled,omitempty"`
	EmailEnabled    *bool `json:"emailEnabled,omitempty"`
	UsernameEnabled *bool `json:"usernameEnabled,omitempty"`
}

type ResourcePermissionAssignment struct {
	TargetType       *EnumPolicyAssignmentTargetType `json:"targetType"`
	TargetIdentifier *string                         `json:"targetIdentifier"`
	Actions          []string                        `json:"actions"`
}

type SearchUserDepartmentOpt struct {
	DepartmentId               *string `json:"departmentId"`
	IncludeChildrenDepartments *bool   `json:"includeChildrenDepartments"`
}

type SearchUserGroupOpt struct {
	Code *string `json:"code"`
}

type SearchUserRoleOpt struct {
	Namespace *string `json:"namespace"`
	Code      string  `json:"code"`
}

type SetUdfValueBatchInput struct {
	TargetId string `json:"targetId"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type SocialConnection struct {
	Provider    string                  `json:"provider"`
	Name        string                  `json:"name"`
	Logo        string                  `json:"logo"`
	Description *string                 `json:"description"`
	Fields      []SocialConnectionField `json:"fields"`
}

type SocialConnectionField struct {
	Key         *string                  `json:"key"`
	Label       *string                  `json:"label"`
	Type        *string                  `json:"type"`
	Placeholder *string                  `json:"placeholder"`
	Children    []*SocialConnectionField `json:"children"`
}

type SocialConnectionFieldInput struct {
	Key         *string                       `json:"key"`
	Label       *string                       `json:"label"`
	Type        *string                       `json:"type"`
	Placeholder *string                       `json:"placeholder"`
	Children    []*SocialConnectionFieldInput `json:"children"`
}

type SocialConnectionInstance struct {
	Provider string                           `json:"provider"`
	Enabled  bool                             `json:"enabled"`
	Fields   []*SocialConnectionInstanceField `json:"fields"`
}

type SocialConnectionInstanceField struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UpdateFunctionInput struct {
	Id          string  `json:"id"`
	Name        *string `json:"name"`
	SourceCode  *string `json:"sourceCode"`
	Description *string `json:"description"`
	Url         *string `json:"url"`
}

type UpdateUserpoolInput struct {
	Name                      *string                            `json:"name,omitempty"`
	Logo                      *string                            `json:"logo,omitempty"`
	Domain                    *string                            `json:"domain,omitempty"`
	Description               *string                            `json:"description,omitempty"`
	UserpoolTypes             []string                           `json:"userpoolTypes,omitempty"`
	EmailVerifiedDefault      *bool                              `json:"emailVerifiedDefault,omitempty"`
	SendWelcomeEmail          *bool                              `json:"sendWelcomeEmail,omitempty"`
	RegisterDisabled          *bool                              `json:"registerDisabled,omitempty"`
	AppSsoEnabled             *bool                              `json:"appSsoEnabled,omitempty"`
	AllowedOrigins            *string                            `json:"allowedOrigins,omitempty"`
	TokenExpiresAfter         *int64                             `json:"tokenExpiresAfter,omitempty"`
	FrequentRegisterCheck     *FrequentRegisterCheckConfigInput  `json:"frequentRegisterCheck,omitempty"`
	LoginFailCheck            *LoginFailCheckConfigInput         `json:"loginFailCheck,omitempty"`
	LoginFailStrategy         *string                            `json:"loginFailStrategy,omitempty"`
	LoginPasswordFailCheck    *LoginPasswordFailCheckConfigInput `json:"loginPasswordFailCheck,omitempty"`
	ChangePhoneStrategy       *ChangePhoneStrategyInput          `json:"changePhoneStrategy,omitempty"`
	ChangeEmailStrategy       *ChangeEmailStrategyInput          `json:"changeEmailStrategy,omitempty"`
	QrcodeLoginStrategy       *QrcodeLoginStrategyInput          `json:"qrcodeLoginStrategy,omitempty"`
	App2WxappLoginStrategy    *App2WxappLoginStrategyInput       `json:"app2WxappLoginStrategy,omitempty"`
	Whitelist                 *RegisterWhiteListConfigInput      `json:"whitelist,omitempty"`
	CustomSMSProvider         *CustomSMSProviderInput            `json:"customSMSProvider,omitempty"`
	LoginRequireEmailVerified *bool                              `json:"loginRequireEmailVerified,omitempty"`
	VerifyCodeLength          *int64                             `json:"verifyCodeLength,omitempty"`
}

type UserDepartment struct {
	Department       Node    `json:"department"`
	IsMainDepartment bool    `json:"isMainDepartment"`
	JoinedAt         *string `json:"joinedAt"`
}

type UserPool struct {
	Id                               string                        `json:"id"`
	Name                             string                        `json:"name"`
	Domain                           string                        `json:"domain"`
	Description                      *string                       `json:"description"`
	Secret                           string                        `json:"secret"`
	JwtSecret                        string                        `json:"jwtSecret"`
	OwnerId                          *string                       `json:"ownerId"`
	UserpoolTypes                    []UserPoolType                `json:"userpoolTypes"`
	Logo                             string                        `json:"logo"`
	CreatedAt                        *string                       `json:"createdAt"`
	UpdatedAt                        *string                       `json:"updatedAt"`
	EmailVerifiedDefault             bool                          `json:"emailVerifiedDefault"`
	SendWelcomeEmail                 bool                          `json:"sendWelcomeEmail"`
	RegisterDisabled                 bool                          `json:"registerDisabled"`
	AppSsoEnabled                    bool                          `json:"appSsoEnabled"`
	ShowWxQRCodeWhenRegisterDisabled *bool                         `json:"showWxQRCodeWhenRegisterDisabled"`
	AllowedOrigins                   *string                       `json:"allowedOrigins"`
	TokenExpiresAfter                *int64                        `json:"tokenExpiresAfter"`
	IsDeleted                        *bool                         `json:"isDeleted"`
	FrequentRegisterCheck            *FrequentRegisterCheckConfig  `json:"frequentRegisterCheck"`
	LoginFailCheck                   *LoginFailCheckConfig         `json:"loginFailCheck"`
	LoginPasswordFailCheck           *LoginPasswordFailCheckConfig `json:"loginPasswordFailCheck"`
	LoginFailStrategy                *string                       `json:"loginFailStrategy"`
	ChangePhoneStrategy              *ChangePhoneStrategy          `json:"changePhoneStrategy"`
	ChangeEmailStrategy              *ChangeEmailStrategy          `json:"changeEmailStrategy"`
	QrcodeLoginStrategy              *QrcodeLoginStrategy          `json:"qrcodeLoginStrategy"`
	App2WxappLoginStrategy           *App2WxappLoginStrategy       `json:"app2WxappLoginStrategy"`
	Whitelist                        *RegisterWhiteListConfig      `json:"whitelist"`
	CustomSMSProvider                *CustomSMSProvider            `json:"customSMSProvider"`
	PackageType                      *int64                        `json:"packageType"`
	UseCustomUserStore               *bool                         `json:"useCustomUserStore"`
	LoginRequireEmailVerified        *bool                         `json:"loginRequireEmailVerified"`
	VerifyCodeLength                 *int64                        `json:"verifyCodeLength"`
}

type UserPoolType struct {
	Code        *string   `json:"code"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Image       *string   `json:"image"`
	Sdks        []*string `json:"sdks"`
}

type WhiteList struct {
	CreatedAt *string `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt"`
	Value     string  `json:"value"`
}

type GqlCommonErrors struct {
	Message   CommonMessageAndCode `json:"message"`
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations"`
	Extensions struct {
		Code      string `json:"code"`
		Extension struct {
			Name string `json:"name"`
		} `json:"extension"`
	}
}

type CommonMessageAndCode struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

type UserPoolEnv struct {
	UserPoolId string    `json:"userPoolId"`
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Id         string    `json:"id"`
}

type UserOrgs []struct {
	Type            string        `json:"type"`
	Id              string        `json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	UpdatedAt       time.Time     `json:"updatedAt"`
	UserPoolId      string        `json:"userPoolId"`
	RootNodeId      string        `json:"rootNodeId,omitempty"`
	Logo            string        `json:"logo"`
	OrgId           string        `json:"orgId,omitempty"`
	Name            string        `json:"name,omitempty"`
	NameI18N        string        `json:"nameI18n"`
	Description     *string       `json:"description,omitempty"`
	DescriptionI18N string        `json:"descriptionI18n"`
	Order           string        `json:"order"`
	Code            *string       `json:"code,omitempty"`
	LeaderUserId    string        `json:"leaderUserId"`
	Source          []interface{} `json:"source,omitempty"`
	DataVersion     interface{}   `json:"dataVersion"`
	SourceData      interface{}   `json:"sourceData"`
}

type GetSecurityLevelResponse struct {
	Score                 int  `json:"score"`
	Email                 bool `json:"email"`
	Phone                 bool `json:"phone"`
	Password              bool `json:"password"`
	PasswordSecurityLevel int  `json:"passwordSecurityLevel"`
	Mfa                   bool `json:"mfa"`
}

type LoginBySubAccountRequest struct {
	Account     string `json:"account"`
	Password    string `json:"password"`
	CaptchaCode string `json:"captchaCode,omitempty"`
	ClientIp    string `json:"clientIp,omitempty"`
}

type IsUserExistsRequest struct {
	Username   *string `json:"username,omitempty"`
	Email      *string `json:"email,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
}
