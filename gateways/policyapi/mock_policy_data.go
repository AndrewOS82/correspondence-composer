package policyapi

import (
	"correspondence-composer/models"
)

func (g *gateway) returnMockSuccess() *models.GetPolicyResponseBody {
	body := &models.GetPolicyResponseBody{
		Message: "SUCCESS",
		Data: models.Policy{
			PlanCode:          "SBFIXUL1",
			PolicyNumber:      "AR38613036",
			PolicyStatus:      "ACTIVE",
			PlanName:          "SB UL Premium Match",
			MarketingName:     "Everly Life",
			QualificationType: "NONQUALIFIED",
			LineOfBusiness:    "LIFE",
			IssueDate:         "2023-06-14",
			IssueState:        "VT",
			PolicyYear:        1,
			ProductType:       "UNIVERSALLIFE",
		},
	}

	return body
}
