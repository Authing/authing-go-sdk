package model

import (
	"time"
)

type PolicyRequest struct {
	Code        string            `json:"code"`
	Description *string           `json:"description,omitempty"`
	Statements  []PolicyStatement `json:"statements,omitempty"`
}

type CreatePolicyResponse struct {
	Namespace        string            `json:"namespace"`
	Code             string            `json:"code"`
	IsDefault        bool              `json:"isDefault"`
	Description      string            `json:"description"`
	Statements       []PolicyStatement `json:"statements"`
	CreatedAt        time.Time         `json:"createdAt"`
	UpdatedAt        time.Time         `json:"updatedAt"`
	AssignmentsCount int               `json:"assignmentsCount"`
}

type UpdatePolicyResponse struct {
	Namespace   string            `json:"namespace"`
	Code        string            `json:"code"`
	IsDefault   bool              `json:"isDefault"`
	Description string            `json:"description"`
	Statements  []PolicyStatement `json:"statements"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

type PaginatedPolicies struct {
	TotalCount int64    `json:"totalCount"`
	List       []Policy `json:"list"`
}

type PaginatedPolicyAssignments struct {
	TotalCount int64              `json:"totalCount"`
	List       []PolicyAssignment `json:"list"`
}

type Policy struct {
	Namespace        string             `json:"namespace"`
	Code             string             `json:"code"`
	IsDefault        bool               `json:"isDefault"`
	Description      *string            `json:"description"`
	Statements       []PolicyStatement  `json:"statements"`
	CreatedAt        *string            `json:"createdAt"`
	UpdatedAt        *string            `json:"updatedAt"`
	AssignmentsCount int64              `json:"assignmentsCount"`
	Assignments      []PolicyAssignment `json:"assignments"`
}

type PolicyAssignment struct {
	Code             string                         `json:"code"`
	TargetType       EnumPolicyAssignmentTargetType `json:"targetType"`
	TargetIdentifier string                         `json:"targetIdentifier"`
}

type PolicyStatement struct {
	Resource  string                     `json:"resource"`
	Actions   []string                   `json:"actions"`
	Effect    *EnumPolicyEffect          `json:"effect"`
	Condition []PolicyStatementCondition `json:"condition,omitempty"`
}

type PolicyStatementCondition struct {
	Param    string `json:"param"`
	Operator string `json:"operator"`
	//Value    Object `json:"value"`
}

type PolicyStatementConditionInput struct {
	Param    string `json:"param"`
	Operator string `json:"operator"`
	//Value    Object `json:"value"`
}

type PolicyStatementInput struct {
	Resource  string                          `json:"resource"`
	Actions   []string                        `json:"actions"`
	Effect    *EnumPolicyEffect               `json:"effect"`
	Condition []PolicyStatementConditionInput `json:"condition"`
}

type PolicyAssignmentsRequest struct {
	Policies          []string                       `json:"policies"`
	TargetType        EnumPolicyAssignmentTargetType `json:"targetType"`
	TargetIdentifiers []string                       `json:"targetIdentifiers"`
}

type SwitchPolicyAssignmentsRequest struct {
	Policy           string                         `json:"policy"`
	TargetType       EnumPolicyAssignmentTargetType `json:"targetType"`
	TargetIdentifier string                         `json:"targetIdentifier"`
	Namespace        *string                        `json:"namespace,omitempty"`
}
