package models

type Policy struct {
	PlanCode          string `json:"planCode,omitempty"`
	PolicyNumber      string `json:"policyNumber,omitempty"`
	PolicyStatus      string `json:"policyStatus,omitempty"`
	PlanName          string `json:"planName,omitempty"`
	MarketingName     string `json:"marketingName,omitempty"`
	QualificationType string `json:"qualificationType,omitempty"`
	LineOfBusiness    string `json:"lineOfBusiness,omitempty"`
	IssueDate         string `json:"issueDate,omitempty"`
	IssueState        string `json:"issueState,omitempty"`
	PolicyYear        int    `json:"policyYear,omitempty"`
	ProductType       string `json:"productType,omitempty"`
}

type GetPolicyResponseBody struct {
	Message string `json:"message"`
	Data    Policy `json:"data,omitempty"`
}
