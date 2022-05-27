package model

type Namespace struct {
	UserPoolId     string `json:"userPoolId"`
	Name           string `json:"name"`
	Code           string `json:"code"`
	Description    string `json:"description"`
	Status         int    `json:"status"`
	ApplicationId  string `json:"applicationId"`
	IsIntegrateApp bool   `json:"isIntegrateApp"`
	IsDefaultApp   bool   `json:"isDefaultApp"`
	Id             int    `json:"id"`
}

type EditNamespaceRequest struct {
	Code        *string `json:"code,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ListGroupsAuthorizedResourcesRequest struct {
	Code         string            `json:"code"`
	Namespace    *string           `json:"namespace,omitempty"`
	ResourceType *EnumResourceType `json:"resourceType,omitempty"`
}
