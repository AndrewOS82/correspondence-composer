package rulesengine

import "correspondence-composer/models"

type MockGateway struct {
	gateway
}

// made available for other tests to import / inject
//nolint
func NewMock() *MockGateway {
	return &MockGateway{
		gateway: gateway{},
	}
}

var MockTestRuleName = "testRule"

func (m *MockGateway) ExecuteRules(rules []*models.Rule) (*models.RulesEngineResponse, error) {
	if len(rules) < 1 {
		return m.gateway.ExecuteRules(rules)
	}

	rule := models.Rule{
		RuleName: MockTestRuleName,
		Version:  1,
	}
	// Default to invalid result
	invalidOutput := models.RuleOutput{}
	rule.Output = []models.RuleOutput{invalidOutput}

	// Mock valid response for valid scenario
	if rules[0].Input["planCode"] == "SBFIXUL1" {
		output := models.RuleOutput{Valid: true}
		rule.Output = []models.RuleOutput{output}
	}

	rulesOutput := []models.Rule{rule}
	return &models.RulesEngineResponse{
		Client:        "Corro",
		Source:        "Camunda",
		CorrelationID: "abc-123",
		Rules:         rulesOutput,
	}, nil
}
