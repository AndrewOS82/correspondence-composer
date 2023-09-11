package policyapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"correspondence-composer/models"
)

const (
	NotFound     = "policy not found"
	Unauthorized = "unauthorized"
)

// a simplified wrapper interface is provided for convenience of testing/mocking
type ClientWrapper interface {
	GetPolicy(ctx context.Context, policyNumber string, token string) (*AppControllerGetPolicyResponse, error)
}

type gateway struct {
	authToken       string
	policyAPIclient ClientWrapper
	mockPolicyData  *models.GetPolicyResponseBody
}

// nolint
func New(client *ClientWithResponses, authToken string, mockPolicyData *models.GetPolicyResponseBody) *gateway {
	return &gateway{
		authToken: authToken,
		mockPolicyData: mockPolicyData,
		policyAPIclient: &clientWrapper{
			client: client,
		},
	}
}

func (g *gateway) FetchPolicyData(ctx context.Context, policyNumber string) (*models.GetPolicyResponseBody, error) {
	if g.authToken == "fake" {
		return g.returnMockSuccess(), nil
	}

	resp, err := g.policyAPIclient.GetPolicy(ctx, policyNumber, g.authToken)
	if err != nil {
		errorMessage := fmt.Sprintf("Error %v\n", err)
		return &models.GetPolicyResponseBody{Message: errorMessage}, err
	}

	var parsedResp models.GetPolicyResponseBody
	err = json.Unmarshal(resp.Body, &parsedResp)
	if err != nil {
		return nil, err
	}

	statusCode := resp.HTTPResponse.StatusCode
	switch statusCode {
	case 401:
		return &parsedResp, errors.New(Unauthorized)
	case 404:
		return &parsedResp, errors.New(NotFound)
	}

	return &parsedResp, nil
}

func (g *gateway) returnMockSuccess() *models.GetPolicyResponseBody {
	return g.mockPolicyData
}

func GetMockPolicyData(configFile string) (*models.GetPolicyResponseBody, error) {
	var policyResponse *models.GetPolicyResponseBody
	content, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &policyResponse)
	if err != nil {
		return nil, err
	}
	return policyResponse, nil
}
