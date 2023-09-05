package usecases

import (
	"context"

	"correspondence-composer/models"
	"correspondence-composer/utils/log"
)

type DataFetcher struct {
	Logger    log.Logger
	PolicyAPI policyAPIGateway
}

type policyAPIGateway interface {
	FetchPolicyData(ctx context.Context, policyNumber string) (*models.GetPolicyResponseBody, error)
}

func (df *DataFetcher) FetchAnniversaryData(ctx context.Context, policyNumber string) (*models.AnniversaryStatement, error) {
	policy, err := df.fetchPolicyData(ctx, policyNumber)
	if err != nil {
		return nil, err
	}

	// If policy data fetch is successful continue fetching other necessary data
	// This may not be necessary, it seems like the policy API has most of the data we need.
	// roles, err := df.fetchRolesData()

	anniversaryStatementData := &models.AnniversaryStatement{
		Policy: policy,
		// Roles: roles
	}

	return anniversaryStatementData, nil
}

func (df *DataFetcher) fetchPolicyData(ctx context.Context, policyNumber string) (*models.Policy, error) {
	resp, err := df.PolicyAPI.FetchPolicyData(ctx, policyNumber)
	if err != nil {
		df.Logger.ErrorWithFields(err, log.Fields{
			"errorMessage": resp.Message,
			"policyNumber": policyNumber,
		})

		return nil, err
	}

	return &resp.Data, nil
}
