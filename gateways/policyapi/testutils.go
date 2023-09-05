package policyapi

import (
	"context"
	"encoding/json"
	"net/http"

	"correspondence-composer/models"
)

type mockPolicyAPIClient struct{}

const (
	ValidToken        = "validtoken"
	InvalidToken      = "invalidtoken"
	ValidPolicyNumber = "AR38613036"
	ValidPlanCode     = "SBFIXUL1"
)

func (m *mockPolicyAPIClient) GetPolicy(_ context.Context, policyNumber string, token string) (*AppControllerGetPolicyResponse, error) {
	if token == InvalidToken {
		// returns a 401
		body := &models.GetPolicyResponseBody{
			Message: "Unauthorized",
		}
		resp := responseHelper(body, 401)

		return resp, nil
	}

	if policyNumber == ValidPolicyNumber {
		// return 200 and body with policy data
		body := &models.GetPolicyResponseBody{
			Message: "SUCCESS",
			Data: models.Policy{
				PolicyNumber: ValidPolicyNumber,
				PlanCode:     ValidPlanCode,
			},
		}
		resp := responseHelper(body, 200)

		return resp, nil
	}

	// default, return 404 not found
	body := &models.GetPolicyResponseBody{
		Message: "Request failed with status code 404",
	}
	resp := responseHelper(body, 404)

	return resp, nil
}

type MockGateway struct {
	gateway
}

// made available for other tests to import / inject
//nolint
func NewMock(token string) *MockGateway {
	return &MockGateway{
		gateway: gateway{
			authToken:       token,
			policyAPIclient: &mockPolicyAPIClient{},
		},
	}
}

func responseHelper(body *models.GetPolicyResponseBody, statusCode int) *AppControllerGetPolicyResponse {
	httpResp := &http.Response{StatusCode: statusCode}
	responseBody, _ := json.Marshal(body)
	resp := &AppControllerGetPolicyResponse{
		HTTPResponse: httpResp,
		Body:         responseBody,
	}

	return resp
}
