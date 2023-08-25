package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	rulesgateway "correspondence-composer/gateways/rulesengine"
	"correspondence-composer/models"
	"correspondence-composer/usecases"
)

func TestRuleExecutor_ValidateAnniversaryData(t *testing.T) {
	// setup
	rulesEngineGateway := rulesgateway.NewMock()
	re := usecases.RuleExecutor{
		RulesEngine: rulesEngineGateway,
	}

	t.Run("When there is invalid policy data it returns failed validations", func(t *testing.T) {
		// execution
		invalidPolicy := &models.Policy{
			PlanCode: "not-valid",
		}
		data := &models.AnniversaryStatement{
			Policy: invalidPolicy,
		}
		validationFailures, err := re.ValidateAnniversaryData(data)

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, len(validationFailures), 1)
		assert.Equal(t, validationFailures[0].RuleName, rulesgateway.MockTestRuleName)
	})

	t.Run("When there is valid policy data it returns no failed validations", func(t *testing.T) {
		// execution
		validPolicy := &models.Policy{
			PlanCode: "SBFIXUL1",
		}
		data := &models.AnniversaryStatement{
			Policy: validPolicy,
		}
		validationFailures, err := re.ValidateAnniversaryData(data)

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, len(validationFailures), 0)
	})
}

func TestRuleExecutor_ValidatePolicyData(t *testing.T) {
	// setup
	rulesEngineGateway := rulesgateway.NewMock()
	re := usecases.RuleExecutor{
		RulesEngine: rulesEngineGateway,
	}

	t.Run("When there is invalid policy data it returns empty output", func(t *testing.T) {
		// execution
		invalidPolicyData := &models.Policy{
			PlanCode: "not-valid",
		}
		ruleResult, err := re.ValidatePolicyData(invalidPolicyData)
		emptyOutput := models.RuleOutput{map[string]interface{}{}}

		// assertion
		assert.NoError(t, err)
		assert.NotNil(t, ruleResult)
		assert.Equal(t, ruleResult[0].Output, emptyOutput)
	})

	t.Run("When there is valid policy data it returns valid output", func(t *testing.T) {
		// execution
		invalidPolicyData := &models.Policy{
			PlanCode: "SBFIXUL1",
		}
		ruleResult, err := re.ValidatePolicyData(invalidPolicyData)
		validOutput := models.RuleOutput{map[string]interface{}{"valid": "true"}}

		// assertion
		assert.NoError(t, err)
		assert.NotNil(t, ruleResult)
		assert.Equal(t, ruleResult[0].Output, validOutput)
	})
}
