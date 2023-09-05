package models

type RuleValidation struct {
	Success  bool
	RuleName string
}

type RulesEngineRequest struct {
	Client string  `json:"client"`
	Source string  `json:"source"`
	Rules  []*Rule `json:"rules"`
}

type RulesEngineResponse struct {
	Client        string `json:"client"`
	Source        string `json:"source"`
	CorrelationID string `json:"correlationid"`
	Rules         []Rule `json:"rules"`
}

type Rule struct {
	RuleName string                 `json:"ruleName"`
	Version  int                    `json:"version"`
	Input    map[string]interface{} `json:"input"`
	Output   RuleOutput             `json:"output,omitempty"`
}

type RuleOutput []map[string]interface{}
