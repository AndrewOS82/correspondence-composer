package usecases

import (
	"encoding/xml"

	"correspondence-composer/models"
	xmlgenmodels "correspondence-composer/models/generated"
)

type XMLBuilder struct{}

func (xb *XMLBuilder) BuildDataServicesLetter(policyData *models.AnniversaryStatement) (string, error) {
	xmlPolicyData := MapAnniversaryData(policyData.Policy)

	policiesArray := []*xmlgenmodels.POLICY{xmlPolicyData}

	policies := xmlgenmodels.POLICIES{
		POLICY: policiesArray,
	}

	letter := &xmlgenmodels.DataServicesLetter{
		POLICIES: &policies,
	}

	preparedXML, err := xml.MarshalIndent(letter, "  ", "    ")

	if err != nil {
		return "", err
	}

	return string(preparedXML), nil
}
