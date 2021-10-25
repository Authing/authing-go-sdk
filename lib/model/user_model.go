package model

import (
	"time"
)

type CreateUserInput struct {
	Username          *string  `json:"username,omitempty"`
	Email             *string  `json:"email,omitempty"`
	EmailVerified     *bool    `json:"emailVerified,omitempty"`
	Phone             *string  `json:"phone,omitempty"`
	PhoneVerified     *bool    `json:"phoneVerified,omitempty"`
	Unionid           *string  `json:"unionid,omitempty"`
	Openid            *string  `json:"openid,omitempty"`
	Nickname          *string  `json:"nickname,omitempty"`
	Photo             *string  `json:"photo,omitempty"`
	Password          *string  `json:"password,omitempty"`
	RegisterSource    []string `json:"registerSource,omitempty"`
	Browser           *string  `json:"browser,omitempty"`
	Oauth             *string  `json:"oauth,omitempty"`
	LoginsCount       *int64   `json:"loginsCount,omitempty"`
	LastLogin         *string  `json:"lastLogin,omitempty"`
	Company           *string  `json:"company,omitempty"`
	LastIP            *string  `json:"lastIP,omitempty"`
	SignedUp          *string  `json:"signedUp,omitempty"`
	Blocked           *bool    `json:"blocked,omitempty"`
	IsDeleted         *bool    `json:"isDeleted,omitempty"`
	Device            *string  `json:"device,omitempty"`
	Name              *string  `json:"name,omitempty"`
	GivenName         *string  `json:"givenName,omitempty"`
	FamilyName        *string  `json:"familyName,omitempty"`
	MiddleName        *string  `json:"middleName,omitempty"`
	Profile           *string  `json:"profile,omitempty"`
	PreferredUsername *string  `json:"preferredUsername,omitempty"`
	Website           *string  `json:"website,omitempty"`
	Gender            *string  `json:"gender,omitempty"`
	Birthdate         *string  `json:"birthdate,omitempty"`
	Zoneinfo          *string  `json:"zoneinfo,omitempty"`
	Locale            *string  `json:"locale,omitempty"`
	Address           *string  `json:"address,omitempty"`
	Formatted         *string  `json:"formatted,omitempty"`
	StreetAddress     *string  `json:"streetAddress,omitempty"`
	Locality          *string  `json:"locality,omitempty"`
	Region            *string  `json:"region,omitempty"`
	PostalCode        *string  `json:"postalCode,omitempty"`
	Country           *string  `json:"country,omitempty"`
	ExternalId        *string  `json:"externalId,omitempty"`
}

type UpdateUserInput struct {
	Email             *string `json:"email,omitempty"`
	Unionid           *string `json:"unionid,omitempty"`
	Openid            *string `json:"openid,omitempty"`
	EmailVerified     *bool   `json:"emailVerified,omitempty"`
	Phone             *string `json:"phone,omitempty"`
	PhoneVerified     *bool   `json:"phoneVerified,omitempty"`
	Username          *string `json:"username,omitempty"`
	Nickname          *string `json:"nickname,omitempty"`
	Password          *string `json:"password,omitempty"`
	Photo             *string `json:"photo,omitempty"`
	Company           *string `json:"company,omitempty"`
	Browser           *string `json:"browser,omitempty"`
	Device            *string `json:"device,omitempty"`
	Oauth             *string `json:"oauth,omitempty"`
	TokenExpiredAt    *string `json:"tokenExpiredAt,omitempty"`
	LoginsCount       *int64  `json:"loginsCount,omitempty"`
	LastLogin         *string `json:"lastLogin,omitempty"`
	LastIP            *string `json:"lastIP,omitempty"`
	Blocked           *bool   `json:"blocked,omitempty"`
	Name              *string `json:"name,omitempty"`
	GivenName         *string `json:"givenName,omitempty"`
	FamilyName        *string `json:"familyName,omitempty"`
	MiddleName        *string `json:"middleName,omitempty"`
	Profile           *string `json:"profile,omitempty"`
	PreferredUsername *string `json:"preferredUsername"`
	Website           *string `json:"website,omitempty"`
	Gender            *string `json:"gender,omitempty"`
	Birthdate         *string `json:"birthdate,omitempty"`
	Zoneinfo          *string `json:"zoneinfo,omitempty"`
	Locale            *string `json:"locale,omitempty"`
	Address           *string `json:"address,omitempty"`
	Formatted         *string `json:"formatted,omitempty"`
	StreetAddress     *string `json:"streetAddress,omitempty"`
	Locality          *string `json:"locality,omitempty"`
	Region            *string `json:"region,omitempty"`
	PostalCode        *string `json:"postalCode,omitempty"`
	City              *string `json:"city,omitempty"`
	Province          *string `json:"province,omitempty"`
	Country           *string `json:"country,omitempty"`
	ExternalId        *string `json:"externalId,omitempty"`
}

type User struct {
	Id                  string                        `json:"id"`
	Arn                 string                        `json:"arn"`
	Status              *EnumUserStatus               `json:"status"`
	UserPoolId          string                        `json:"userPoolId"`
	Username            *string                       `json:"username"`
	Email               *string                       `json:"email"`
	EmailVerified       *bool                         `json:"emailVerified"`
	Phone               *string                       `json:"phone"`
	PhoneVerified       *bool                         `json:"phoneVerified"`
	Unionid             *string                       `json:"unionid"`
	Openid              *string                       `json:"openid"`
	Identities          []*Identity                   `json:"identities"`
	Nickname            *string                       `json:"nickname"`
	RegisterSource      []string                      `json:"registerSource"`
	Photo               *string                       `json:"photo"`
	Password            *string                       `json:"password"`
	Oauth               *string                       `json:"oauth"`
	Token               *string                       `json:"token"`
	TokenExpiredAt      *string                       `json:"tokenExpiredAt"`
	LoginsCount         *int64                        `json:"loginsCount"`
	LastLogin           *string                       `json:"lastLogin"`
	LastIP              *string                       `json:"lastIP"`
	SignedUp            *string                       `json:"signedUp"`
	Blocked             *bool                         `json:"blocked"`
	IsDeleted           *bool                         `json:"isDeleted"`
	Device              *string                       `json:"device"`
	Browser             *string                       `json:"browser"`
	Company             *string                       `json:"company"`
	Name                *string                       `json:"name"`
	GivenName           *string                       `json:"givenName"`
	FamilyName          *string                       `json:"familyName"`
	MiddleName          *string                       `json:"middleName"`
	Profile             *string                       `json:"profile"`
	PreferredUsername   *string                       `json:"preferredUsername"`
	Website             *string                       `json:"website"`
	Gender              *string                       `json:"gender"`
	Birthdate           *string                       `json:"birthdate"`
	Zoneinfo            *string                       `json:"zoneinfo"`
	Locale              *string                       `json:"locale"`
	Address             *string                       `json:"address"`
	Formatted           *string                       `json:"formatted"`
	StreetAddress       *string                       `json:"streetAddress"`
	Locality            *string                       `json:"locality"`
	Region              *string                       `json:"region"`
	PostalCode          *string                       `json:"postalCode"`
	City                *string                       `json:"city"`
	Province            *string                       `json:"province"`
	Country             *string                       `json:"country"`
	CreatedAt           *string                       `json:"createdAt"`
	UpdatedAt           *string                       `json:"updatedAt"`
	Roles               *PaginatedRoles               `json:"roles"`
	Groups              *PaginatedGroups              `json:"groups"`
	Departments         *PaginatedDepartments         `json:"departments"`
	AuthorizedResources *PaginatedAuthorizedResources `json:"authorizedResources"`
	ExternalId          *string                       `json:"externalId"`
	CustomData          []*UserCustomData             `json:"customData"`
}

type UserCustomData struct {
	Key      string          `json:"key,omitempty"`
	Value    *string         `json:"value,omitempty"`
	Label    *string         `json:"label,omitempty"`
	DataType EnumUDFDataType `json:"dataType,omitempty"`
}

type UserDdfInput struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type UserDefinedData struct {
	Key      string          `json:"key,omitempty"`
	DataType EnumUDFDataType `json:"dataType,omitempty"`
	Value    string          `json:"value,omitempty"`
	Label    *string         `json:"label,omitempty"`
}

type UserDefinedDataInput struct {
	Key   string  `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type UserDefinedDataMap struct {
	TargetId string            `json:"targetId"`
	Data     []UserDefinedData `json:"data"`
}

type UserDefinedField struct {
	TargetType EnumUDFTargetType `json:"targetType"`
	DataType   EnumUDFDataType   `json:"dataType"`
	Key        string            `json:"key"`
	Label      *string           `json:"label"`
	Options    *string           `json:"options"`
}

type CreateUserRequest struct {
	UserInfo     CreateUserInput `json:"userInfo,omitempty"`
	KeepPassword bool            `json:"keepPassword,omitempty"`
	CustomData   []KeyValuePair  `json:"params,omitempty"`
}

type CommonPageUsersResponse struct {
	TotalCount int    `json:"totalCount"`
	List       []User `json:"list"`
}

type FindUserRequest struct {
	Email          *string `json:"email,omitempty"`
	Username       *string `json:"username,omitempty"`
	Phone          *string `json:"phone,omitempty"`
	ExternalId     *string `json:"externalId,omitempty"`
	WithCustomData bool    `json:"withCustomData,omitempty"`
}

type SearchUserRequest struct {
	Query          string    `json:"query"`
	Page           int       `json:"page"`
	Limit          int       `json:"limit"`
	DepartmentOpts *[]string `json:"departmentOpts,omitempty"`
	GroupOpts      *[]string `json:"groupOpts,omitempty"`
	RoleOpts       *[]string `json:"roleOpts,omitempty"`
	WithCustomData bool
}

type GetUserGroupsResponse struct {
	Groups struct {
		TotalCount int          `json:"totalCount"`
		List       []GroupModel `json:"list"`
	} `json:"groups"`
}

type GetUserRolesRequest struct {
	Id        string `json:"id"`
	Namespace string `json:"namespace,omitempty"`
}

type GetUserRolesResponse struct {
	Roles struct {
		TotalCount int         `json:"totalCount"`
		List       []RoleModel `json:"list"`
	} `json:"roles"`
}

type UserRoleOptRequest struct {
	UserIds   []string `json:"userIds"`
	RoleCodes []string `json:"roleCodes"`
	Namespace *string  `json:"namespace"`
}

type OrgModel struct {
	RootNodeId      string    `json:"rootNodeId"`
	Logo            string    `json:"logo"`
	Type            string    `json:"type"`
	Id              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UserPoolId      string    `json:"userPoolId"`
	OrgId           string    `json:"orgId"`
	Name            string    `json:"name"`
	NameI18N        string    `json:"nameI18n"`
	Description     string    `json:"description"`
	DescriptionI18N string    `json:"descriptionI18n"`
	Order           string    `json:"order"`
	Code            string    `json:"code"`
	LeaderUserId    string    `json:"leaderUserId"`
	Source          []string  `json:"source"`
	DataVersion     string    `json:"dataVersion"`
	SourceData      string    `json:"sourceData"`
}

type ListUserOrgResponse struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Data    [][]OrgModel `json:"data"`
}

type ListUserAuthResourceRequest struct {
	Id           string           `json:"id"`
	Namespace    string           `json:"namespace"`
	ResourceType EnumResourceType `json:"resourceType"`
}
