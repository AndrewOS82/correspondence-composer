package policyapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolicyAPIClient_ReturnsErrorWhenAuthTokenInvalid(t *testing.T) {
	token := InvalidToken
	policyAPIGateway, ctx := setup(token)
	policyNumber := ValidPolicyNumber

	resp, err := policyAPIGateway.FetchPolicyData(ctx, policyNumber)

	assert.Equal(t, err, ErrUnauthorized)
	assert.Equal(t, resp.Message, "Unauthorized")
}

func TestPolicyAPIClient_ReturnsErrorWhenPolicyNotFound(t *testing.T) {
	token := ValidToken
	policyAPIGateway, ctx := setup(token)
	policyNumber := "NOTAPOLICY"

	resp, err := policyAPIGateway.FetchPolicyData(ctx, policyNumber)

	assert.Equal(t, err, ErrNotFound)
	assert.Equal(t, resp.Message, "Request failed with status code 404")
}

func TestPolicyAPIClient_ReturnsPolicyDataOnSuccess(t *testing.T) {
	token := ValidToken
	policyAPIGateway, ctx := setup(token)
	policyNumber := ValidPolicyNumber

	resp, err := policyAPIGateway.FetchPolicyData(ctx, policyNumber)

	assert.NoError(t, err)
	assert.Equal(t, resp.Data.PlanCode, ValidPlanCode)
}

func setup(token string) (*gateway, context.Context) {
	mockPolicyAPIClient := &mockPolicyAPIClient{}

	return &gateway{
		authToken:       token,
		policyAPIclient: mockPolicyAPIClient,
	}, context.Background()
}
