package model

import "time"

type GroupModel struct {
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateGroupsRequest struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type UpdateGroupsRequest struct {
	Code        string  `json:"code"`
	NewCode     *string `json:"newCode,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type GetGroupUserResponse struct {
	Users struct {
		TotalCount int    `json:"totalCount"`
		List       []User `json:"list"`
	} `json:"users"`
}
