package usecases

import (
	"correspondence-composer/models"
)

type DataFetcher struct{}

func (df *DataFetcher) FetchAnniversaryData() (*models.AnniversaryStatement, error) {
	policy, err := df.fetchPolicyData()
	if err != nil {
		// handle
		return nil, err
	}

	// If policy data fetch is successful continue fetching other necessary data
	// roles, err := df.fetchRolesData()

	anniversaryStatementData := &models.AnniversaryStatement{
		Policy: policy,
		// Roles: roles
	}

	return anniversaryStatementData, nil
}

func (df *DataFetcher) fetchPolicyData() (*models.Policy, error) { //nolint:all
	println("fetching policy data!")

	// If we set up our models properly with json attributes we should be able to
	// hopefully easily unmarshal the API JSON response into the Policy type.
	return &models.Policy{PlanCode: "SBFIXUL1"}, nil
}
