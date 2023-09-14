package models

type GetPolicyResponseBody struct {
	Message string `json:"message"`
	Data    Policy `json:"data,omitempty"`
}

type Policy struct {
	Event                         string               `json:"event,omitempty"`
	MaturityDate                  string               `json:"maturityDate,omitempty"`
	PlanCode                      string               `json:"planCode,omitempty"`
	PolicyNumber                  string               `json:"policyNumber,omitempty"`
	PolicyStatus                  string               `json:"policyStatus,omitempty"`
	PlanName                      string               `json:"planName,omitempty"`
	ShortName                     string               `json:"shortName,omitempty"`
	MarketingName                 string               `json:"marketingName,omitempty"`
	QualificationType             string               `json:"qualificationType,omitempty"`
	LineOfBusiness                string               `json:"lineOfBusiness,omitempty"`
	IssueDate                     string               `json:"issueDate,omitempty"`
	IssueState                    string               `json:"issueState,omitempty"`
	IssueType                     string               `json:"issueType,omitempty"`
	PolicyYear                    int                  `json:"policyYear,omitempty"`
	ProductType                   string               `json:"productType,omitempty"`
	PolicyStartDate               string               `json:"policyStartDate,omitempty"`
	PolicyTerminationDate         string               `json:"policyTerminationDate,omitempty"`
	ResidenceState                string               `json:"residenceState,omitempty"`
	PolicyTerm                    int                  `json:"policyTerm,omitempty"`
	PremiumFrequency              string               `json:"premiumFrequency,omitempty"`
	ApplicationDate               string               `json:"applicationDate,omitempty"`
	PreviousPolicyAnniversaryDate string               `json:"previousPolicyAnniversaryDate,omitempty"`
	NextAnniversaryDate           string               `json:"nextAnniversaryDate,omitempty"`
	CarrierID                     string               `json:"carrierID,omitempty"`
	LoanValues                    *LoanValues          `json:"loanValues,omitempty"`
	WithDrawalValues              *WithDrawalValues    `json:"withDrawalValues,omitempty"`
	PolicyValues                  *PolicyValues        `json:"policyValues,omitempty"`
	DistributionDetails           *DistributionDetails `json:"distributionDetails,omitempty"`
	Coverage                      *Coverage            `json:"coverage,omitempty"`
	PolicyFeatures                []*PolicyFeature     `json:"policyFeatures,omitempty"`
	Parties                       []*Party             `json:"parties,omitempty"`
	PartyRoles                    []*PartyRole         `json:"partyRoles,omitempty"`
	SystematicPrograms            []*SystematicProgram `json:"systematicPrograms,omitempty"`
	Timestamp                     string               `json:"timestamp,omitempty"`
}

type SystematicProgram struct {
	Amount    float64 `json:"amount,omitempty"`
	Frequency string  `json:"frequency,omitempty"`
	ArrType   string  `json:"arrType,omitempty"`
	Reason    string  `json:"reason,omitempty"`
}

type PolicyValues struct {
	BeginningAcctValue float64 `json:"beginningAcctValue,omitempty"`
	EndingAcctValue    float64 `json:"endingAcctValue,omitempty"`
	TotalYTDPremiumAmt float64 `json:"totalYTDPremiumAmt,omitempty"`
	SurrenderValue     float64 `json:"surrenderValue,omitempty"`
}

type DistributionDetails struct {
	LoanSegments []*LoanSegments `json:"loanSegments,omitempty"`
}

type LoanSegments struct {
	LoanInterestRate string `json:"loanInterestRate,omitempty"`
	LoanCreditRate   string `json:"loanCreditRate,omitempty"`
}

type WithDrawalValues struct {
	TotalWithdrawalAmt float64 `json:"totalWithdrawalAmt,omitempty"`
}

type LoanValues struct {
	TotalLoanAccruedInterest float64 `json:"totalLoanAccuredInterest,omitempty"`
	TotalLoanPrincipal       float64 `json:"totalLoanPrincipal,omitempty"`
	TotalLoanBalance         float64 `json:"totalLoanBalance,omitempty"`
	LoanPayoffAmt            float64 `json:"totalLoanPayoffAmt,omitempty"`
}

type PolicyFeature struct {
	PaymentAmt float64 `json:"paymentAmt,omitempty"`
}

type Coverage struct {
	CoverageLayers   []*CoverageLayer `json:"coverageLayers,omitempty"`
	TotalCoverageAmt float64          `json:"totalCoverageAmt,omitempty"`
}

type CoverageLayer struct {
	CoverageParticipants []*CoverageParticipant `json:"coverageParticipants,omitempty"`
	OriginalCoverageAmt  float64                `json:"originalCoverageAmt,omitempty"`
	GrossDeathBenefitAmt float64                `json:"grossDeathBenefitAmt,omitempty"`
	CoverageType         string                 `json:"coverageType,omitempty"`
}

type CoverageParticipant struct {
	RiskClass string `json:"riskClass,omitempty"`
	IssueAge  int    `json:"issueAge,omitempty"`
	PartyID   string `json:"partyID,omitempty"`
}

type Party struct {
	PartyID               string            `json:"partyID,omitempty"`
	FullName              string            `json:"fullName,omitempty"`
	PartyType             string            `json:"partyType,omitempty"`
	Prefix                string            `json:"prefix,omitempty"`
	FirstName             string            `json:"firstName,omitempty"`
	MiddleName            string            `json:"middleName,omitempty"`
	LastName              string            `json:"lastName,omitempty"`
	Suffix                string            `json:"suffix,omitempty"`
	DateOfBirth           string            `json:"dateOfBirth,omitempty"`
	AttainedAge           int               `json:"attainedAge,omitempty"`
	Gender                string            `json:"gender,omitempty"`
	PrefCommunicationType string            `json:"prefCommunicationType,omitempty"`
	BeneficiaryPercentage int               `json:"beneficiaryPercentage,omitempty"`
	Emails                []*Email          `json:"emails,omitempty"`
	Phones                []*Phone          `json:"phones,omitempty"`
	Addresses             []*Address        `json:"addresses,omitempty"`
	Identifications       []*Identification `json:"identifications,omitempty"`
}

type Email struct {
	EmailAddress string `json:"emailAddress,omitempty"`
	EmailType    string `json:"emailType,omitempty"`
}

type Phone struct {
	CountryCode string `json:"countryCode,omitempty"`
	AreaCode    string `json:"areaCode,omitempty"`
	DialNumber  string `json:"dialNumber,omitempty"`
	PhoneType   string `json:"phoneType,omitempty"`
}

type Address struct {
	AddrLine1      string `json:"addrLine1,omitempty"`
	AddrLine2      string `json:"addrLine2,omitempty"`
	AddrLine3      string `json:"addrLine3,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	ZipCode        string `json:"zipCode,omitempty"`
	AddrCountry    string `json:"addrCountry,omitempty"`
	AddressType    string `json:"addressType,omitempty"`
	PrefAddressInd string `json:"prefAddressInd,omitempty"`
	EndDate        string `json:"endDate,omitempty"`
}

type PartyAttribute struct {
	StartDate string  `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
}

type Identification struct {
	IdentificationValue string `json:"identificationValue,omitempty"`
	IdentificationType  string `json:"identificationType,omitempty"`
}

type PartyRole struct {
	PartyRole         string `json:"partyRole,omitempty"`
	RelationToInsured string `json:"relationToInsured,omitempty"`
	PartyID           string `json:"partyID,omitempty"`
}
