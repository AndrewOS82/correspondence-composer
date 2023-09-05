package usecases

import (
	"encoding/xml"

	"correspondence-composer/models"
	. "correspondence-composer/models/generated" //nolint:all
)

type XMLBuilder struct{}

func (xb *XMLBuilder) BuildDataServicesLetter(statement *models.AnniversaryStatement) (string, error) {
	roles := ROLE{
		ROLCONT:         "",
		ROLCODE:         "",
		ROLPCT:          "",
		ROLFULLNAME:     "",
		ROLPREFIX:       "",
		ROLFSTNAME:      "",
		ROLMI:           "",
		ROLLSTNAME:      "",
		ROLSUFFIX:       "",
		ROLNAMETYPE:     "",
		ADDRLINE1:       "",
		ADDRLINE2:       "",
		ADDRLINE3:       "",
		ADDRCITY:        "",
		ADDRSTATE:       "",
		ADDRZIP:         0,
		ADDRCNTRYCODE:   "",
		ADDRTYPE:        "",
		ROLEMAIL:        "",
		ROLPHONENUM:     0,
		ROLPHONETYPE:    "",
		ROLEXTERNALID:   "",
		ROLMASTERNUM:    "",
		ROLTAXID:        "",
		ROLDELIVERYFLAG: "",
		ROLDELIVERYDESC: "",
		ROLNAMEID:       "",
		ROLROLEID:       0,
	}

	carrier := CARRIER{
		CARRIERCONT:          "",
		CARRIERNAME:          "",
		CARRIERADDRLINE1:     "",
		CARRIERADDRLINE2:     "",
		CARRIERADDRLINE3:     "",
		CARRIERCITY:          "",
		CARRIERSTATE:         "",
		CARRIERZIP:           0,
		CARRIERWEBADDR:       "",
		CARRIERPHONENUM:      0,
		CARRIERFAXNUM:        "",
		CARRIEROFFICEHOURS:   "",
		CARRIEROFFICEDAYS:    "",
		CARRIERBUDGETCNTR:    "",
		CARRIERPOBADDRLINE1:  "",
		CARRIERPOBADDRLINE2:  "",
		CARRIERPOBADDRLINE3:  "",
		CARRIERPOBCITY:       "",
		CARRIERPOBSTATE:      "",
		CARRIERPOBZIP:        0,
		CARRIERMAILADDRLINE1: "",
		CARRIERMAILADDRLINE2: "",
		CARRIERMAILADDRLINE3: "",
		CARRIERMAILCITY:      "",
		CARRIERMAILSTATE:     "",
		CARRIERMAILZIP:       0,
		CARRIERSERVICECENTER: "",
	}

	policy := POLICY{
		POLCONT:      statement.Policy.PolicyNumber,
		POLPLANCODE:  statement.Policy.PlanCode,
		POLPRODUCT:   statement.Policy.PlanName,
		POLQUALDESC:  statement.Policy.QualificationType,
		POLCOMPNAME:  statement.Policy.MarketingName,
		POLCYCLEDATE: "",
		// TODO: Pull name and desc from ref table
		POLSERVICENAME: "Statement",
		POLSERVICEDESC: "UL Anniversary statement",
		POLSOURCE:      "Zinnia",
		POLSTATUS:      statement.Policy.PolicyStatus,
		POLPOLNUM:      statement.Policy.PolicyNumber,
		// TODO: Pull below from reftable
		POLPRDCTCOMP: "",
		POLLOB:       statement.Policy.LineOfBusiness,
		// TODO: Pull below from reftable
		POLPRCSGCOMP:  "",
		POLTRACKINGID: "",
		// TODO: Pull below from reftable
		POLDOCTYPE:       "ANN",
		POLPRDCTMKTGNAME: statement.Policy.MarketingName,
		POLISSUEDATE:     statement.Policy.IssueDate,
		POLISSUESTATE:    statement.Policy.IssueState,
		POLCURRYR:        "",
		// TODO: Pull below from reftable
		POLSYSCODE:         "",
		POLDTHBNFTGAURAGE:  "",
		POLDBTYPECODE:      "",
		POLRESIDENCESTATE:  "",
		POLTELTXFRIND:      "",
		POLHMRIND:          0,
		POLGROUPIND:        "",
		POLPLANTYPE:        statement.Policy.ProductType,
		POLCONTYEAR:        0,
		POLIRRSTRATEGY:     "",
		POLRUNDATE:         "",
		POLAGE55:           "",
		POLPOLYEAR:         "",
		ROLES:              &roles,
		CARRIERINFORMATION: &carrier,
	}

	policiesArray := []*POLICY{&policy}

	policies := POLICIES{
		POLICY: policiesArray,
	}

	letter := &DataServicesLetter{
		POLICIES: &policies,
	}

	preparedXML, err := xml.MarshalIndent(letter, "  ", "    ")

	if err != nil {
		return "", err
	}

	return string(preparedXML), nil
}
