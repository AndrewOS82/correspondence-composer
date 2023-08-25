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

func (m *MockGateway) ExecuteRules(rules []*models.Rule) (*models.RulesAdminResponse, error) {
	if len(rules) < 1 {
		return m.gateway.ExecuteRules(rules)
	}

	rule := models.Rule{
		RuleName: MockTestRuleName,
		Version:  1,
	}
	// Default to invalid result
	invalidOutput := models.RuleOutput{
		map[string]interface{}{},
	}
	rule.Output = invalidOutput

	// Mock valid response for valid scenario
	if rules[0].Input["planCode"] == "SBFIXUL1" {
		validOutput := models.RuleOutput{
			map[string]interface{}{"valid": "true"},
		}
		rule.Output = validOutput
	}

	rulesOutput := []models.Rule{rule}
	return &models.RulesAdminResponse{
		Client:        "Corro",
		Source:        "Camunda",
		CorrelationID: "abc-123",
		Rules:         rulesOutput,
	}, nil
}
