package policyapi

import (
	"correspondence-composer/models"
)

func (g *gateway) returnMockSuccess() *models.GetPolicyResponseBody {
	body := &models.GetPolicyResponseBody{
		Message: "SUCCESS",
		Data: models.Policy{
			PolicyNumber: "AR38613036",
			PlanCode:     "SBFIXUL1",
		},
	}

	return body
}
