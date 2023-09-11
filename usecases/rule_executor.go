package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"strings"

	"correspondence-composer/models"
	"correspondence-composer/utils/log"
)

const PolicyDatapoint = "policyDetails"

var (
	ErrNoResults    = errors.New("no results returned for rules")
	ErrBadDatapoint = errors.New("this is not a policy data point")
)

type RuleExecutor struct {
	Logger      log.Logger
	RulesEngine rulesEngineGateway
	RulesConfig []*models.RuleConfig
}

type rulesEngineGateway interface {
	ExecuteRules(rules []*models.Rule) (*models.RulesEngineResponse, error)
}

func (re *RuleExecutor) ValidateAnniversaryData(ctx context.Context, data *models.AnniversaryStatement) ([]*models.RuleValidation, error) {
	validationResults, err := re.ValidatePolicyData(ctx, data.Policy)
	if err != nil {
		return nil, err
	}

	// Iterate through the results and track which validations failed so we can
	// send out detailed exceptions.
	var failedValidations []*models.RuleValidation
	for _, result := range validationResults {
		re.Logger.InfoWithFields("Rule result", log.Fields{
			"_ruleName": result.RuleName,
			"version":   result.Version,
			"input":     result.Input,
			"output":    result.Output,
		})

		if len(result.Output) > 0 && !result.Output[0].Valid {
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

	return failedValidations, nil
}

func (re *RuleExecutor) ValidatePolicyData(_ context.Context, data *models.Policy) ([]models.Rule, error) {
	var rules []*models.Rule
	for _, ruleConfig := range re.RulesConfig {
		rule := re.populatePolicyDataInput(data, ruleConfig)
		rules = append(rules, rule)
	}

	resp, err := re.RulesEngine.ExecuteRules(rules)
	if err != nil {
		return nil, err
	}

	if len(resp.Rules) < 1 {
		return nil, ErrNoResults
	}
	ruleResults := resp.Rules

	return ruleResults, nil
}

func (re *RuleExecutor) populatePolicyDataInput(data *models.Policy, ruleConfig *models.RuleConfig) *models.Rule {
	inputMap := make(map[string]interface{})

	for _, input := range ruleConfig.Inputs {
		key := input["key"]
		value, _ := re.getPolicyDatapoint(data, input["value"])
		inputMap[key] = value
	}

	rule := &models.Rule{
		RuleName: ruleConfig.RuleName,
		Version:  ruleConfig.Version,
		Input:    inputMap,
	}

	return rule
}

func (re *RuleExecutor) getPolicyDatapoint(data *models.Policy, datapoint string) (interface{}, error) {
	split := strings.Split(datapoint, ".")
	if !strings.EqualFold(split[0], PolicyDatapoint) {
		err := ErrBadDatapoint
		re.Logger.ErrorWithFields(err, log.Fields{
			"datapoint": datapoint,
		})
		return nil, err
	}

	policyAttribute := split[1]
	policy := reflect.ValueOf(*data)
	value := policy.FieldByName(policyAttribute)

	return value.Interface(), nil
}

func GetRulesConfig(configFile string) ([]*models.RuleConfig, error) {
	var rulesConfig []*models.RuleConfig

	content, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &rulesConfig)
	if err != nil {
		return nil, err
	}

	return rulesConfig, nil
}
