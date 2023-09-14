package mapper_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"correspondence-composer/mapper"
	"correspondence-composer/models"
	xmlgenmodels "correspondence-composer/models/generated"
)

func TestMapper_BasePolicy(t *testing.T) {
	apiData, mappedData := setup()

	assert.Equal(t, mappedData.POLCONT, apiData.PolicyNumber)
	assert.Equal(t, mappedData.POLPLANCODE, apiData.PlanCode)
	assert.Equal(t, mappedData.POLQUALTYPE, apiData.QualificationType)
	assert.Equal(t, mappedData.POLPRDCTMKTGNAME, apiData.MarketingName)
	assert.Equal(t, mappedData.POLLOB, apiData.LineOfBusiness)
	assert.Equal(t, mappedData.POLISSUEDATE, apiData.IssueDate)

	faceValue := int(apiData.Coverage.TotalCoverageAmt)
	assert.Equal(t, mappedData.POLFACEVALUE, faceValue)
}

func TestMapper_CoverageParticipant(t *testing.T) {
	apiData, mappedData := setup()
	coveragelayer := apiData.Coverage.CoverageLayers[0]
	coverageparticipant := coveragelayer.CoverageParticipants[0]

	assert.Equal(t, mappedData.POLRISKCLASS, coverageparticipant.RiskClass)
	assert.Equal(t, mappedData.POLISSUEAGE, coverageparticipant.IssueAge)
	assert.Equal(t, mappedData.POLICYVALUES.PVALBEGINCOVERVAL, coveragelayer.OriginalCoverageAmt)
	assert.Equal(t, mappedData.POLICYVALUES.PVALENDDEATHBNFTVAL, coveragelayer.GrossDeathBenefitAmt)

	coverageChangeValue := apiData.Coverage.TotalCoverageAmt - coveragelayer.OriginalCoverageAmt
	assert.Equal(t, mappedData.POLICYVALUES.PVALCHGCOVERVAL, coverageChangeValue)
}

func TestMapper_SystematicProgram(t *testing.T) {
	apiData, mappedData := setup()
	systematicProgram := apiData.SystematicPrograms[0]

	assert.Equal(t, "PREMIUMREASON", systematicProgram.Reason)
	assert.Equal(t, mappedData.POLPREMAMT, systematicProgram.Amount)
	assert.Equal(t, mappedData.POLICYVALUES.PVALCURRPAYMODE, systematicProgram.Frequency)
}

func TestMapper_PolicyValues(t *testing.T) {
	apiData, mappedData := setup()

	assert.Equal(t, apiData.PolicyValues.BeginningAcctValue, mappedData.POLICYVALUES.PVALBEGINACCTVAL)
	assert.Equal(t, apiData.PolicyValues.EndingAcctValue, mappedData.POLICYVALUES.PVALENDACCTVAL)

	changeValue := apiData.PolicyValues.EndingAcctValue - apiData.PolicyValues.BeginningAcctValue
	assert.Equal(t, changeValue, mappedData.POLICYVALUES.PVALCHGACCTVAL)
}

func TestMapper_MapPartiesToPolicy(t *testing.T) {
	apiData, mappedData := setup()

	assert.IsType(t, &xmlgenmodels.PARTIES{}, mappedData.PARTIES)

	var parties []*xmlgenmodels.PARTY
	for _, party := range mappedData.PARTIES.PARTY {
		assert.Equal(t, party.PARTYCONT, apiData.PolicyNumber)
		if party.PARTYID == "Party_PI_1" {
			assert.Equal(t, party.PARTYFSTNAME, "Molly")
			assert.Equal(t, party.PARTYDOB, "2003-10-05")
			assert.Equal(t, party.PARTYPCT, 0)
		}
		if party.PARTYID == "Party_PB_Primary_Bene_1" {
			assert.Equal(t, party.PARTYLSTNAME, "Padilla")
			assert.Equal(t, party.PARTYROLENAME, "PRIMARYBENEFICIARY")
			assert.Equal(t, party.PARTYRELTOINS, "Stepparent")
			assert.Equal(t, party.PARTYPCT, 100)
		}
		parties = append(parties, party)
	}
	assert.Equal(t, len(parties), len(apiData.PartyRoles))
}

func setup() (*models.Policy, *xmlgenmodels.POLICY) {
	var policyResponse *models.GetPolicyResponseBody
	content, _ := os.ReadFile("../gateways/policyapi/sample_policy_data.json")
	json.Unmarshal(content, &policyResponse)

	mockapiData := &models.AnniversaryStatement{
		Policy: &policyResponse.Data,
	}

	mappedXMLData := mapper.MapAnniversaryData(mockapiData.Policy)

	return mockapiData.Policy, mappedXMLData
}
