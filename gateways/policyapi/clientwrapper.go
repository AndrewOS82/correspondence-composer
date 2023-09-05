package policyapi

import (
	"context"
)

type clientWrapper struct {
	client *ClientWithResponses
}

func (cw *clientWrapper) GetPolicy(ctx context.Context, policyNumber string, token string) (*AppControllerGetPolicyResponse, error) {
	params := AppControllerGetPolicyParams{Authorization: "Bearer " + token}
	resp, err := cw.client.AppControllerGetPolicyWithResponse(ctx, policyNumber, &params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
