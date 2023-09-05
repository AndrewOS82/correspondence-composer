package models

type Policy struct {
	PlanCode      string
	PolicyNumber  string
	PolicyStatus  string
	PlanName      string
	MarketingName string
}

type GetPolicyResponseBody struct {
	Message string `json:"message"`
	Data    Policy `json:"data,omitempty"`
}
