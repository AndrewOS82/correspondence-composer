package mapper

import (
	"time"

	"correspondence-composer/models"
	xmlgenmodels "correspondence-composer/models/generated"
)

// MapAnniversaryData - map policy API response to final XML struct
// any calculations/coercion for the XML values should go here
func MapAnniversaryData(policyData *models.Policy) *xmlgenmodels.POLICY {

	finalXMLData := mapPolicyData(policyData)
	return finalXMLData
}

func MapPartiesToPolicy(data *models.Policy) *xmlgenmodels.PARTIES {
	var xmlParties []*xmlgenmodels.PARTY
	partyData := getAllPartyData(data)

	for _, p := range partyData {
		xmlparty := &xmlgenmodels.PARTY{
			PARTYCONT: data.PolicyNumber,
			PARTYID:   p.Party.PartyID,
			//REF TABLE FROM RULES DOC
			PARTYROLEID:    0,
			PARTYTYPE:      p.Party.PartyType,
			PARTYROLENAME:  p.PartyRole.PartyRole,
			PARTYRELTOINS:  p.PartyRole.RelationToInsured,
			PARTYPCT:       p.Party.BeneficiaryPercentage,
			PARTYFULLNAME:  p.Party.FullName,
			PARTYPREFIX:    p.Party.Prefix,
			PARTYFSTNAME:   p.Party.FirstName,
			PARTYMI:        p.Party.MiddleName,
			PARTYLSTNAME:   p.Party.LastName,
			PARTYSUFFIX:    p.Party.Suffix,
			PARTYDOB:       p.Party.DateOfBirth,
			PARTYATTAINAGE: p.Party.AttainedAge,
			PARTYGENDER:    p.Party.Gender,
		}

		xmlParties = append(xmlParties, xmlparty)
	}

	return &xmlgenmodels.PARTIES{
		PARTY: xmlParties,
	}
}

func mapPolicyData(data *models.Policy) *xmlgenmodels.POLICY {
	policy := &xmlgenmodels.POLICY{
		RECTYPE:     "",
		POLCONT:     data.PolicyNumber,
		POLPLANCODE: data.PlanCode,
		POLPRODUCT:  data.PlanName,
		POLQUALTYPE: data.QualificationType,
		//NO DATAPOINT DEFINED
		POLPLANTYPE:      "",
		POLPRODSHORTNAME: data.ShortName,
		POLPRDCTMKTGNAME: data.MarketingName,
		//NO DATAPOINT DEFINED
		POLCYCLEDATE: "",
		//REF TABLE FROM RULES DOC
		POLSERVICENAME: "",
		POLSERVICEDESC: "",
		POLSTATUS:      data.PolicyStatus,
		POLLOB:         data.LineOfBusiness,
		POLDOCTYPE:     data.Event,
		//Requires separate call to /carriers endpoint
		POLPRDCTCOMP:  "Everly Life",
		POLPRCSGCOMP:  "",
		POLISSUEDATE:  data.IssueDate,
		POLISSUESTATE: data.IssueState,
		POLCURRYR:     data.PolicyYear,
		//NO DATAPOINT DEFINED
		POLSYSCODE:        "",
		POLPRODUCTTYPE:    data.ProductType,
		POLSTARTDATE:      data.PolicyStartDate,
		POLMATURITYDATE:   data.MaturityDate,
		POLTERMDATE:       data.PolicyTerminationDate,
		POLRESIDENCESTATE: data.ResidenceState,
		POLTERM:           data.PolicyTerm,
		POLPREMMODE:       data.PremiumFrequency,
		POLEFFECTIVEDATE:  data.ApplicationDate,
		POLLASTANNIVDATE:  data.PreviousPolicyAnniversaryDate,
		//No DATAPOINT DEFINED
		POLCURRANNIVDATE:    "",
		POLNEXTANNIVENDDATE: data.NextAnniversaryDate,
		POLFACEVALUE:        int(data.Coverage.TotalCoverageAmt),
		POLICYVALUES:        mapPolicyValuesToPolicy(data),
		PARTIES:             MapPartiesToPolicy(data),
		// nolint
		//CARRIERINFORMATION: mapCarrierDataToPolicy()
	}

	systematicProgram := getSystematicProgram(data.SystematicPrograms)
	if systematicProgram != nil {
		policy.POLPREMAMT = systematicProgram.Amount
	}

	coverageParticipant := getCoverageParticipant(data.Coverage.CoverageLayers)
	if coverageParticipant != nil {
		policy.POLRISKCLASS = coverageParticipant.RiskClass
		policy.POLISSUEAGE = coverageParticipant.IssueAge
	}

	datelayout := "2006-01-02"
	// POLSTARTANNIVDATE = PreviousPolicyAnniversaryDate - 1 year
	prevPolicyAnniversaryDate, err := time.Parse(datelayout, data.PreviousPolicyAnniversaryDate)
	if err == nil {
		// subtract a year
		oneYearAgo := prevPolicyAnniversaryDate.AddDate(-1, 0, 0)
		// format string
		policy.POLSTARTANNIVDATE = oneYearAgo.Format(datelayout)
	}

	// POLCURRANNIVENDDATE = NextAnniversaryDate - 1 day
	nextAnniversaryDate, err := time.Parse(datelayout, data.NextAnniversaryDate)
	if err == nil {
		// subtract a day
		oneDayAgo := nextAnniversaryDate.AddDate(0, 0, 1)
		// format string
		policy.POLCURRANNIVENDDATE = oneDayAgo.Format(datelayout)
	}

	return policy
}

func mapPolicyValuesToPolicy(data *models.Policy) *xmlgenmodels.POLICYVALUE {
	policyvalues := &xmlgenmodels.POLICYVALUE{
		//NO DATAPOINTS DEFINED FOR THE BELOW
		PVALCONT:        "",
		PVALBEGINDATE:   "",
		PVALENDDATE:     "",
		PVALENDCOVERVAL: data.Coverage.TotalCoverageAmt,
		// TO DO: Need clarity on values to subtract (policy data from last anniversary year)
		// end value of last year - beginning value of current year
		PVALBEGINDEATHBNFTVAL: 0,
		// TO DO: PVALCHGDEATHBNFTVAL = PVALBEGINDEATHBNFTVAL - PVALENDDEATHBNFTVAL
		PVALCHGDEATHBNFTVAL: 0,
		// TO DO: Need clarity on values to subtract (policy data from last anniversary year)
		// end value of last year - beginning value of current year
		PVALBEGINSURRVAL: 0,
		PVALENDSURRVAL:   data.PolicyValues.SurrenderValue,
		// TO DO: PVALCHGDEATHBNFTVAL = PVALBEGINSURRVAL - PVALENDSURRVAL
		PVALCHGSURRVAL:   0,
		PVALBEGINACCTVAL: data.PolicyValues.BeginningAcctValue,
		PVALENDACCTVAL:   data.PolicyValues.EndingAcctValue,
		PVALCHGACCTVAL:   data.PolicyValues.EndingAcctValue - data.PolicyValues.BeginningAcctValue,
		PVALTOTALPAYRCVD: data.PolicyValues.TotalYTDPremiumAmt,
		// TO DO: PVALTOTALCOICOST = PVALCOSTOFINSVAL + PVALTOTALUNITCHARGE + PVALTOTALEXPENSECHARGE
		// REQUIRES SEPARATE CALL TO /transactions
		PVALTOTALCOICOST:       0,
		PVALCOSTOFINSVAL:       0,
		PVALTOTALUNITCHARGE:    0,
		PVALTOTALEXPENSECHARGE: 0,
		PVALTOTALINTCREDIT:     0,

		PVALTOTALPARTWITHDRAWVAL: data.WithDrawalValues.TotalWithdrawalAmt,
		PVALTOTALLOANAMT:         data.LoanValues.TotalLoanPrincipal,
		PVALTOTALLOANBALNAMT:     data.LoanValues.TotalLoanBalance,
		PVALTOTALLOANREPAYAMT:    data.LoanValues.LoanPayoffAmt,
		PVALTOTALINTCHARGE:       data.LoanValues.TotalLoanAccruedInterest,
		//NO DATAPOINT DEFINED
		PVALINTCREDITCURRRATE: 0,

		//REQUIRES SEPARATE CALL TO /distribution
		PVALLOANINTCURRRATE:      0,
		PVALLOANCREDITCURRRATE:   0,
		PVALLOANINTNEXTYRRATE:    0,
		PVALLOANCREDITNEXTYRRATE: 0,

		//NO DATAPOINT DEFINED
		PVALINTCREDITNEXTYRRATE: 0,
		PVALANNLPSAMT:           data.PolicyFeatures[0].PaymentAmt,
	}

	systematicProgram := getSystematicProgram(data.SystematicPrograms)
	if systematicProgram != nil {
		policyvalues.PVALCURRPAYAMT = systematicProgram.Amount
		policyvalues.PVALCURRPAYMODE = systematicProgram.Frequency
	}

	coverageLayer := getCoverageLayer(data.Coverage.CoverageLayers)
	if coverageLayer != nil {
		policyvalues.PVALBEGINCOVERVAL = coverageLayer.OriginalCoverageAmt
		policyvalues.PVALCHGCOVERVAL = data.Coverage.TotalCoverageAmt - coverageLayer.OriginalCoverageAmt
		policyvalues.PVALENDDEATHBNFTVAL = coverageLayer.GrossDeathBenefitAmt
	}

	return policyvalues
}

func getPartyDataByID(id string, parties []*models.Party) *models.Party {
	for _, party := range parties {
		if party.PartyID == id {
			return party
		}
	}
	return nil
}

type AllPartyData struct {
	Party     *models.Party
	PartyRole *models.PartyRole
}

func getAllPartyData(data *models.Policy) []*AllPartyData {
	var allPartyData []*AllPartyData

	for _, role := range data.PartyRoles {
		partyDataFromAPI := getPartyDataByID(role.PartyID, data.Parties)
		if partyDataFromAPI != nil {
			partydata := &AllPartyData{
				Party:     partyDataFromAPI,
				PartyRole: role,
			}
			allPartyData = append(allPartyData, partydata)
		}
	}

	return allPartyData
}

func getSystematicProgram(allSysPrograms []*models.SystematicProgram) *models.SystematicProgram {
	for _, sps := range allSysPrograms {
		if sps.ArrType == "PAYMENT" && sps.Reason == "PREMIUMREASON" {
			return sps
		}
	}

	return nil
}

func getCoverageLayer(allCoverageLayers []*models.CoverageLayer) *models.CoverageLayer {
	for _, covlayer := range allCoverageLayers {
		if covlayer.CoverageType == "BASE" {
			return covlayer
		}
	}

	return nil
}

func getCoverageParticipant(allCoverageLayers []*models.CoverageLayer) *models.CoverageParticipant {
	covlayer := getCoverageLayer(allCoverageLayers)
	if covlayer != nil {
		for _, participant := range covlayer.CoverageParticipants {
			if participant.PartyID == "Party_PI_1" {
				return participant
			}
		}
		return nil
	}

	return nil
}

//func constructPhoneNum(phoneData *models.Phone) string {
//	countryCode := phoneData.CountryCode
//	areaCode := phoneData.AreaCode
//	dialNumber := phoneData.DialNumber
//
//	return fmt.Sprintf("%s %s %s", countryCode, areaCode, dialNumber)
//}
