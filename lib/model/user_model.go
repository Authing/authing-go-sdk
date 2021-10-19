package model

import (
	"github.com/Authing/authing-go-sdk/lib/constant"
	"time"
)

type CreateUserInput struct {
	Username          *string  `json:"username"`
	Email             *string  `json:"email"`
	EmailVerified     *bool    `json:"emailVerified"`
	Phone             *string  `json:"phone"`
	PhoneVerified     *bool    `json:"phoneVerified"`
	Unionid           *string  `json:"unionid"`
	Openid            *string  `json:"openid"`
	Nickname          *string  `json:"nickname"`
	Photo             *string  `json:"photo"`
	Password          *string  `json:"password"`
	RegisterSource    []string `json:"registerSource"`
	Browser           *string  `json:"browser"`
	Oauth             *string  `json:"oauth"`
	LoginsCount       *int64   `json:"loginsCount"`
	LastLogin         *string  `json:"lastLogin"`
	Company           *string  `json:"company"`
	LastIP            *string  `json:"lastIP"`
	SignedUp          *string  `json:"signedUp"`
	Blocked           *bool    `json:"blocked"`
	IsDeleted         *bool    `json:"isDeleted"`
	Device            *string  `json:"device"`
	Name              *string  `json:"name"`
	GivenName         *string  `json:"givenName"`
	FamilyName        *string  `json:"familyName"`
	MiddleName        *string  `json:"middleName"`
	Profile           *string  `json:"profile"`
	PreferredUsername *string  `json:"preferredUsername"`
	Website           *string  `json:"website"`
	Gender            *string  `json:"gender"`
	Birthdate         *string  `json:"birthdate"`
	Zoneinfo          *string  `json:"zoneinfo"`
	Locale            *string  `json:"locale"`
	Address           *string  `json:"address"`
	Formatted         *string  `json:"formatted"`
	StreetAddress     *string  `json:"streetAddress"`
	Locality          *string  `json:"locality"`
	Region            *string  `json:"region"`
	PostalCode        *string  `json:"postalCode"`
	Country           *string  `json:"country"`
	ExternalId        *string  `json:"externalId"`
}

type UpdateUserInput struct {
	Email             *string `json:"email"`
	Unionid           *string `json:"unionid"`
	Openid            *string `json:"openid"`
	EmailVerified     *bool   `json:"emailVerified"`
	Phone             *string `json:"phone"`
	PhoneVerified     *bool   `json:"phoneVerified"`
	Username          *string `json:"username"`
	Nickname          *string `json:"nickname"`
	Password          *string `json:"password"`
	Photo             *string `json:"photo"`
	Company           *string `json:"company"`
	Browser           *string `json:"browser"`
	Device            *string `json:"device"`
	Oauth             *string `json:"oauth"`
	TokenExpiredAt    *string `json:"tokenExpiredAt"`
	LoginsCount       *int64  `json:"loginsCount"`
	LastLogin         *string `json:"lastLogin"`
	LastIP            *string `json:"lastIP"`
	Blocked           *bool   `json:"blocked"`
	Name              *string `json:"name"`
	GivenName         *string `json:"givenName"`
	FamilyName        *string `json:"familyName"`
	MiddleName        *string `json:"middleName"`
	Profile           *string `json:"profile"`
	PreferredUsername *string `json:"preferredUsername"`
	Website           *string `json:"website"`
	Gender            *string `json:"gender"`
	Birthdate         *string `json:"birthdate"`
	Zoneinfo          *string `json:"zoneinfo"`
	Locale            *string `json:"locale"`
	Address           *string `json:"address"`
	Formatted         *string `json:"formatted"`
	StreetAddress     *string `json:"streetAddress"`
	Locality          *string `json:"locality"`
	Region            *string `json:"region"`
	PostalCode        *string `json:"postalCode"`
	City              *string `json:"city"`
	Province          *string `json:"province"`
	Country           *string `json:"country"`
	ExternalId        *string `json:"externalId"`
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
	Key      string          `json:"key"`
	Value    *string         `json:"value"`
	Label    *string         `json:"label"`
	DataType EnumUDFDataType `json:"dataType"`
}

type UserDdfInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UserDefinedData struct {
	Key      string          `json:"key"`
	DataType EnumUDFDataType `json:"dataType"`
	Value    string          `json:"value"`
	Label    *string         `json:"label"`
}

type UserDefinedDataInput struct {
	Key   string  `json:"key"`
	Value *string `json:"value"`
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
	UserInfo     CreateUserInput `json:"userInfo"`
	KeepPassword bool            `json:"keepPassword"`
	CustomData   []KeyValuePair  `json:"params"`
}

type CommonPageUsersResponse struct {
	TotalCount int    `json:"totalCount"`
	List       []User `json:"list"`
}

type FindUserRequest struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Phone          string `json:"phone"`
	ExternalId     string `json:"externalId"`
	WithCustomData bool   `json:"withCustomData"`
}

type SearchUserRequest struct {
	Query          string   `json:"query"`
	Page           int      `json:"page"`
	Limit          int      `json:"limit"`
	DepartmentOpts []string `json:"departmentOpts"`
	GroupOpts      []string `json:"groupOpts"`
	RoleOpts       []string `json:"roleOpts"`
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
	Namespace string `json:"namespace"`
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
	Id           string                    `json:"id"`
	Namespace    string                    `json:"namespace"`
	ResourceType constant.ResourceTypeEnum `json:"resourceType"`
}
