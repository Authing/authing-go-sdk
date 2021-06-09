package model

type ListMemberRequest struct {
	NodeId string `json:"nodeId"`
	Page int `json:"page"`
	Limit int  `json:"limit"`
	IncludeChildrenNodes bool `json:"includeChildrenNodes"`
}

type UserDetailData struct {
	ThirdPartyIdentity User `json:"thirdPartyIdentity"`
}

type UserDetailResponse struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
	Data User `json:"data"`
}

type ExportAllOrganizationResponse struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
	Data []Node `json:"data"`
}

type NodeByIdDetail struct {
	NodeById Node `json:"nodeById"`
}

type NodeByIdResponse struct {
	Data NodeByIdDetail `json:"data"`
}