package model

type CreateOrgRequest struct {
	Name        string  `json:"name"`
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

type OrgNode struct {
	Id              string     `json:"id"`
	OrgId           *string    `json:"orgId"`
	CreatedAt       *string    `json:"createdAt"`
	UpdatedAt       *string    `json:"updatedAt"`
	UserPoolId      *string    `json:"userPoolId"`
	Name            string     `json:"name"`
	Description     *string    `json:"description"`
	DescriptionI18n *string    `json:"descriptionI18n"`
	Order           *int64     `json:"order"`
	Code            *string    `json:"code"`
	Members         *[]User    `json:"members,omitempty"`
	Children        *[]OrgNode `json:"children,omitempty"`
}

type OrgResponse struct {
	Id       string     `json:"id"`
	RootNode *OrgNode   `json:"rootNode,omitempty"`
	Nodes    *[]OrgNode `json:"nodes,omitempty"`
}

type PaginatedOrgs struct {
	TotalCount int64 `json:"totalCount"`
	List       []Org `json:"list"`
}

type Node struct {
	Id                  string                        `json:"id"`
	OrgId               *string                       `json:"orgId"`
	Name                string                        `json:"name"`
	NameI18n            *string                       `json:"nameI18n"`
	Description         *string                       `json:"description"`
	DescriptionI18n     *string                       `json:"descriptionI18n"`
	Order               *int64                        `json:"order"`
	Code                *string                       `json:"code"`
	Root                *bool                         `json:"root"`
	Depth               *int64                        `json:"depth"`
	Path                []string                      `json:"path"`
	CodePath            []*string                     `json:"codePath"`
	NamePath            []string                      `json:"namePath"`
	CreatedAt           *string                       `json:"createdAt"`
	UpdatedAt           *string                       `json:"updatedAt"`
	Children            []string                      `json:"children"`
	Users               PaginatedUsers                `json:"users"`
	AuthorizedResources *PaginatedAuthorizedResources `json:"authorizedResources"`
}

type Org struct {
	Id       string `json:"id"`
	RootNode Node   `json:"rootNode"`
	Nodes    []Node `json:"nodes"`
}

type AddNodeOrg struct {
	Id       string            `json:"id"`
	RootNode OrgNodeChildStr   `json:"rootNode"`
	Nodes    []OrgNodeChildStr `json:"nodes"`
}
type AddOrgNodeRequest struct {
	OrgId           string  `json:"orgId"`
	ParentNodeId    string  `json:"parentNodeId"`
	Name            string  `json:"name"`
	Code            *string `json:"code,omitempty"`
	Description     *string `json:"description,omitempty"`
	Order           *int    `json:"order,omitempty"`
	NameI18N        *string `json:"nameI18n,omitempty"`
	DescriptionI18N *string `json:"descriptionI18n,omitempty"`
}

type OrgNodeChildStr struct {
	Id              string    `json:"id"`
	OrgId           *string   `json:"orgId"`
	Name            string    `json:"name"`
	NameI18n        *string   `json:"nameI18n"`
	Description     *string   `json:"description"`
	DescriptionI18n *string   `json:"descriptionI18n"`
	Order           *int64    `json:"order"`
	Code            *string   `json:"code"`
	Root            *bool     `json:"root"`
	Depth           *int64    `json:"depth"`
	Path            []string  `json:"path"`
	CodePath        []*string `json:"codePath"`
	NamePath        []string  `json:"namePath"`
	CreatedAt       *string   `json:"createdAt"`
	UpdatedAt       *string   `json:"updatedAt"`
	Children        []string  `json:"children"`
}

type UpdateOrgNodeRequest struct {
	Id          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ListAuthorizedResourcesByNodeCodeRequest struct {
	Id           string  `json:"id"`
	Code         string  `json:"code"`
	Namespace    *string `json:"namespace,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}
