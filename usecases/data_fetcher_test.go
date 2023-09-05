package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"correspondence-composer/gateways/policyapi"
	"correspondence-composer/usecases"
	"correspondence-composer/utils/log"
)

func TestDataFetcher_FetchAnniversaryData(t *testing.T) {
	// setup
	policyAPIGateway := policyapi.NewMock("validtoken")
	logger := log.New(log.Config{
		ServiceName: "correspondence-composer-test",
	})
	df := usecases.DataFetcher{
		Logger:    logger,
		PolicyAPI: policyAPIGateway,
	}

	t.Run("When the API fails to fetch data it returns an error", func(t *testing.T) {
		// execution
		_, err := df.FetchAnniversaryData(context.Background(), "notapolicy")

		// assertion
		assert.Equal(t, err, errors.New(policyapi.NotFound))
	})

	t.Run("When the API is successful it returns data", func(t *testing.T) {
		// execution
		resp, err := df.FetchAnniversaryData(context.Background(), policyapi.ValidPolicyNumber)

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, resp.Policy.PlanCode, policyapi.ValidPlanCode)
	})
}
