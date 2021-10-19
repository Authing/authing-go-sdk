package model

import "github.com/Authing/authing-go-sdk/lib/enum"

type Role struct {
	Id                  string                        `json:"id"`
	Namespace           string                        `json:"namespace"`
	Code                string                        `json:"code"`
	Arn                 string                        `json:"arn"`
	Description         *string                       `json:"description"`
	IsSystem            *bool                         `json:"isSystem"`
	CreatedAt           *string                       `json:"createdAt"`
	UpdatedAt           *string                       `json:"updatedAt"`
	Users               PaginatedUsers                `json:"users"`
	AuthorizedResources *PaginatedAuthorizedResources `json:"authorizedResources"`
	Parent              *Role                         `json:"parent"`
}

type RoleModel struct {
	Id          string  `json:"id"`
	Namespace   string  `json:"namespace"`
	Code        string  `json:"code"`
	Arn         string  `json:"arn"`
	Description *string `json:"description"`
	CreatedAt   *string `json:"createdAt"`
	UpdatedAt   *string `json:"updatedAt"`
	Parent      *struct {
		Id          string  `json:"id"`
		Namespace   string  `json:"namespace"`
		Code        string  `json:"code"`
		Arn         string  `json:"arn"`
		Description *string `json:"description"`
		CreatedAt   *string `json:"createdAt"`
		UpdatedAt   *string `json:"updatedAt"`
	} `json:"parent"`
}

type GetRoleListRequest struct {
	Page      int             `json:"page"`
	Limit     int             `json:"limit"`
	SortBy    enum.SortByEnum `json:"sortBy"`
	Namespace string          `json:"namespace"`
}

type Roles struct {
	Roles PaginatedRoles `json:"roles"`
}
type GetRoleListResponse struct {
	Data Roles `json:"data"`
}

type GetRoleUserListRequest struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Code      string `json:"code"`
	Namespace string `json:"namespace"`
}

type CreateRoleRequest struct {
	Code        string `json:"code"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
	ParentCode  string `json:"parent,omitempty"`
}

type DeleteRoleRequest struct {
	Code      string `json:"code"`
	Namespace string `json:"namespace,omitempty"`
}

type DeleteRole struct {
	DeleteRole Role `json:"createRole"`
}

type BatchDeleteRoleRequest struct {
	CodeList  []string `json:"codeList"`
	Namespace string   `json:"namespace,omitempty"`
}

type UpdateRoleRequest struct {
	Code        string `json:"code"`
	NewCode     string `json:"newCode,omitempty"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
	ParentCode  string `json:"parent,omitempty"`
}

type RoleDetailRequest struct {
	Code      string `json:"code"`
	Namespace string `json:"namespace,omitempty"`
}

type AssignAndRevokeRoleRequest struct {
	RoleCodes []string `json:"roleCodes"`
	Namespace string   `json:"namespace,omitempty"`
	UserIds   []string `json:"userIds"`
}

type AuthorizedResources struct {
	TotalCount int `json:"totalCount"`
	List       []struct {
		Code    string   `json:"code"`
		Type    string   `json:"type"`
		Actions []string `json:"actions"`
	} `json:"list"`
}

type BatchRoleUdv struct {
	TargetId string            `json:"targetId"`
	Data     []UserDefinedData `json:"data"`
}
