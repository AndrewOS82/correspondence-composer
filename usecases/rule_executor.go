package usecases

import (
	"errors"
	"fmt"

	"correspondence-composer/models"
)

type RuleExecutor struct {
	RulesEngine rulesEngineGateway
}

type rulesEngineGateway interface {
	ExecuteRules(rules []*models.Rule) (*models.RulesEngineResponse, error)
}

func (re *RuleExecutor) ValidateAnniversaryData(data *models.AnniversaryStatement) ([]*models.RuleValidation, error) {
	validationResults, err := re.ValidatePolicyData(data.Policy)
	if err != nil {
		return nil, err
	}

	// Iterate through the results and track which validations failed so we can
	// send out detailed exceptions.
	var failedValidations []*models.RuleValidation
	for _, result := range validationResults {
		// There is a lot of flexibility in the rules engine around what output is returned when a rule is run.
		// It seems like the pattern for DMN is: if match, return something, otherwise return nothing.
		// This is a placeholder for now based on that assumption.
		if len(result.Output[0]) < 1 {
			validation := &models.RuleValidation{
				RuleName: result.RuleName,
				Success:  false,
			}
			failedValidations = append(failedValidations, validation)
		}
	}
	// Question for product, given new auditing requirements: If there are failed validations
	// can we stop validating the other fetched data and early return?

	// If policy data is validated successfully continue validating other necessary data.
	// roles, err := df.validateRolesData()

	return failedValidations, nil
}

func (re *RuleExecutor) ValidatePolicyData(data *models.Policy) ([]models.Rule, error) {
	fmt.Printf("validating policy data! %v\n", data)

	// This is a random rule and input for demonstration purposes. We'll need to decide on a strategy
	// for selecting the appropriate rules to validate policy data once they have been created.
	rule := &models.Rule{
		RuleName: "testRule",
		Version:  1,
		Input:    map[string]interface{}{"planCode": data.PlanCode},
	}
	rules := []*models.Rule{rule}

	resp, err := re.RulesEngine.ExecuteRules(rules)
	if err != nil {
		return nil, err
	}

	if len(resp.Rules) < 1 {
		return nil, errors.New("no results returned for rules")
	}
	ruleResults := resp.Rules
	fmt.Printf("Rule result: %v\n", ruleResults)

	return ruleResults, nil
}
