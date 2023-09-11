package usecases_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	rulesgateway "correspondence-composer/gateways/rulesengine"
	"correspondence-composer/models"
	"correspondence-composer/usecases"
	"correspondence-composer/utils/log"
)

func TestRuleExecutor_ValidateAnniversaryData(t *testing.T) {
	// setup
	logger := log.New(log.Config{
		ServiceName: "correspondence-composer-test",
	})

	rulesEngineGateway := rulesgateway.NewMock()
	input := map[string]string{
		"key":   "planCode",
		"value": usecases.PolicyDatapoint + ".PlanCode",
	}
	ruleConfig := &models.RuleConfig{
		RuleName: rulesgateway.MockTestRuleName,
		Version:  1,
		Inputs:   []map[string]string{input},
	}
	rulesConfig := []*models.RuleConfig{ruleConfig}

	re := usecases.RuleExecutor{
		Logger:      logger,
		RulesEngine: rulesEngineGateway,
		RulesConfig: rulesConfig,
	}

	t.Run("When there is invalid policy data it returns failed validations", func(t *testing.T) {
		// execution
		invalidPolicy := &models.Policy{
			PlanCode: "not-valid",
		}
		data := &models.AnniversaryStatement{
			Policy: invalidPolicy,
		}
		validationFailures, err := re.ValidateAnniversaryData(context.Background(), data)

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, len(validationFailures), 1)
		assert.Equal(t, rulesgateway.MockTestRuleName, validationFailures[0].RuleName)
	})

	t.Run("When there is valid policy data it returns no failed validations", func(t *testing.T) {
		// execution
		validPolicy := &models.Policy{
			PlanCode: "SBFIXUL1",
		}
		data := &models.AnniversaryStatement{
			Policy: validPolicy,
		}
		validationFailures, err := re.ValidateAnniversaryData(context.Background(), data)

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, len(validationFailures), 0)
	})
}

func TestRuleExecutor_ValidatePolicyData(t *testing.T) {
	// setup
	logger := log.New(log.Config{
		ServiceName: "correspondence-composer-test",
	})

	rulesEngineGateway := rulesgateway.NewMock()
	input := map[string]string{
		"key":   "planCode",
		"value": usecases.PolicyDatapoint + ".PlanCode",
	}
	ruleConfig := &models.RuleConfig{
		RuleName: rulesgateway.MockTestRuleName,
		Version:  1,
		Inputs:   []map[string]string{input},
	}
	rulesConfig := []*models.RuleConfig{ruleConfig}

	re := usecases.RuleExecutor{
		Logger:      logger,
		RulesEngine: rulesEngineGateway,
		RulesConfig: rulesConfig,
	}

	t.Run("When there is invalid policy data it returns valid false", func(t *testing.T) {
		// execution
		policyData := &models.Policy{
			PlanCode: "not-valid",
		}
		ruleResult, err := re.ValidatePolicyData(context.Background(), policyData)
		ruleOutput := models.RuleOutput{Valid: false}
		expectedOutput := []models.RuleOutput{ruleOutput}

		// assertion
		assert.NoError(t, err)
		assert.NotNil(t, ruleResult)
		assert.Equal(t, expectedOutput, ruleResult[0].Output)
	})

	t.Run("When there is valid policy data it returns valid output", func(t *testing.T) {
		// execution
		policyData := &models.Policy{
			PlanCode: "SBFIXUL1",
		}
		ruleResult, err := re.ValidatePolicyData(context.Background(), policyData)
		ruleOutput := models.RuleOutput{Valid: true}
		expectedOutput := []models.RuleOutput{ruleOutput}

		// assertion
		assert.NoError(t, err)
		assert.NotNil(t, ruleResult)
		assert.Equal(t, expectedOutput, ruleResult[0].Output)
	})
}
