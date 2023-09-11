package usecases

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
		//NO DATAPOINT DEFINED
		POLPRCSGCOMP: "",
		POLDOCTYPE:   data.Event,
		//NO DATAPOINT DEFINED, should it be hardcoded?
		POLPRDCTCOMP:  "Everly Life",
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
		// TO DO: Determine which SystematicProgram to pull from
		POLPREMAMT:  data.SystematicPrograms[0].Amount,
		POLPREMMODE: data.PremiumFrequency,
		// TO DO: Determine logic for which coverage layer / participant to pull from
		POLRISKCLASS:     data.Coverage.CoverageLayers[0].CoverageParticipants[0].RiskClass,
		POLEFFECTIVEDATE: data.ApplicationDate,
		// TO DO: Determine logic for which coverage layer / participant to pull from
		POLISSUEAGE:      data.Coverage.CoverageLayers[0].CoverageParticipants[0].IssueAge,
		POLLASTANNIVDATE: data.PreviousPolicyAnniversaryDate,
		//No DATAPOINT DEFINED
		POLCURRANNIVDATE:    "",
		POLNEXTANNIVENDDATE: data.NextAnniversaryDate,
		POLFACEVALUE:        int(data.Coverage.TotalCoverageAmt),
		POLICYVALUES:        mapPolicyValuesToPolicy(data),
		//PARTYINFORMATION:    mapPartyInfoToPolicy(data),
		// nolint
		//CARRIERINFORMATION: mapCarrierDataToPolicy()
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
		PVALCONT:      "",
		PVALBEGINDATE: "",
		PVALENDDATE:   "",
		// TO DO: Determine logic for which coverage layer to pull from
		PVALBEGINCOVERVAL: data.Coverage.CoverageLayers[0].OriginalCoverageAmt,
		PVALENDCOVERVAL:   data.Coverage.TotalCoverageAmt,
		PVALCHGCOVERVAL:   data.Coverage.TotalCoverageAmt - data.Coverage.CoverageLayers[0].OriginalCoverageAmt,
		// TO DO: Need clarity on values to subtract (policy data from last anniversary year)
		// end value of last year - beginning value of current year
		PVALBEGINDEATHBNFTVAL: 0,
		// TO DO: Determine logic for which coverage layer to pull from
		PVALENDDEATHBNFTVAL: data.Coverage.CoverageLayers[0].GrossDeathBenefitAmt,
		// TO DO: PVALCHGDEATHBNFTVAL = PVALBEGINDEATHBNFTVAL - PVALENDDEATHBNFTVAL
		PVALCHGDEATHBNFTVAL: 0,
		// TO DO: Need clarity on values to subtract (policy data from last anniversary year)
		// end value of last year - beginning value of current year
		PVALBEGINSURRVAL: 0,
		PVALENDSURRVAL:   float64(int(data.PolicyValues.SurrenderValue)),
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
		// TO DO: Determine logic for which SystematicProgram/PolicyFeatures to pull from
		PVALCURRPAYAMT:  data.SystematicPrograms[0].Amount,
		PVALCURRPAYMODE: data.SystematicPrograms[0].Frequency,
		PVALANNLPSAMT:   data.PolicyFeatures[0].PaymentAmt,
	}

	return policyvalues
}

//func mapPartyInfoToPolicy(data *models.Policy) *xmlgenmodels.PARTY {
//	// TO DO: The below will just use the first party listed;
//	// this is a placeholder until we have finalized logic
//
//	party := &xmlgenmodels.PARTY{
//		PARTYCONT: data.PolicyNumber,
//		PARTYID:   data.Parties[0].PartyID,
//		//REF TABLE FROM RULES DOC
//		PARTYROLEID:    0,
//		PARTYTYPE:      data.Parties[0].PartyType,
//		PARTYROLENAME:  data.PartyRoles[0].PartyRole,
//		PARTYRELTOINS:  data.PartyRoles[0].RelationToInsured,
//		PARTYPCT:       data.Parties[0].BeneficiaryPercentage,
//		PARTYFULLNAME:  data.Parties[0].FullName,
//		PARTYPREFIX:    data.Parties[0].Prefix,
//		PARTYFSTNAME:   data.Parties[0].FirstName,
//		PARTYMI:        data.Parties[0].MiddleName,
//		PARTYLSTNAME:   data.Parties[0].LastName,
//		PARTYSUFFIX:    data.Parties[0].Suffix,
//		ADDRLINE1:      data.Parties[0].Addresses[0].AddrLine1,
//		ADDRLINE2:      data.Parties[0].Addresses[0].AddrLine2,
//		ADDRLINE3:      data.Parties[0].Addresses[0].AddrLine3,
//		ADDRCITY:       data.Parties[0].Addresses[0].City,
//		ADDRSTATE:      data.Parties[0].Addresses[0].State,
//		ADDRZIP:        data.Parties[0].Addresses[0].ZipCode,
//		ADDRCNTRYCODE:  data.Parties[0].Addresses[0].AddrCountry,
//		ADDRTYPE:       data.Parties[0].Addresses[0].AddressType,
//		ADDRPREFIND:    data.Parties[0].Addresses[0].PrefAddressInd,
//		PARTYDOB:       data.Parties[0].DateOfBirth,
//		PARTYATTAINAGE: data.Parties[0].AttainedAge,
//		PARTYGENDER:    data.Parties[0].Gender,
//		EMAILADDR:      data.Parties[0].Emails[0].EmailAddress,
//		EMAILTYPE:      data.Parties[0].Emails[0].EmailType,
//		PHNNUM:         constructPhoneNum(data.Parties[0].Phones[0]),
//		PHNTYPE:        data.Parties[0].Phones[0].PhoneType,
//		//NO DATAPOINTS DEFINED
//		PREFERRED:         "",
//		PARTYDELIVERYFLAG: "",
//		PARTYDELIVERYDESC: "",
//	}
//
//	if data.Parties[0].Addresses[0].EndDate != "null" {
//		party.ADDRSTATUS = "Inactive"
//	} else {
//		party.ADDRSTATUS = "Active"
//	}
//
//	return party
//}
//
//func constructPhoneNum(phoneData *models.Phone) string {
//	countryCode := phoneData.CountryCode
//	areaCode := phoneData.AreaCode
//	dialNumber := phoneData.DialNumber
//
//	return fmt.Sprintf("%s %s %s", countryCode, areaCode, dialNumber)
//}
