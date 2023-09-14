// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"encoding/xml"
)

// AlphanumericString ...
type AlphanumericString string

// Email ...
type Email string

// CARRIERADDRESS ...
type CARRIERADDRESS struct {
	XMLName                xml.Name `xml:"CARRIER_ADDRESS"`
	CARRIERADDRLINE1       string   `xml:"CARRIER_ADDR_LINE1"`
	CARRIERADDRLINE2       string   `xml:"CARRIER_ADDR_LINE2"`
	CARRIERADDRLINE3       string   `xml:"CARRIER_ADDR_LINE3"`
	CARRIERCITY            string   `xml:"CARRIER_CITY"`
	CARRIERSTATE           string   `xml:"CARRIER_STATE"`
	CARRIERZIP             int      `xml:"CARRIER_ZIP"`
	CARRIERADDRCOUNTRYCODE string   `xml:"CARRIER_ADDR_COUNTRY_CODE"`
	CARRIERADDRTYPE        string   `xml:"CARRIER_ADDR_TYPE"`
	CARRIERADDRSTATUS      string   `xml:"CARRIER_ADDR_STATUS"`
}

// CARRIERADDRESSES ...
type CARRIERADDRESSES struct {
	XMLName        xml.Name          `xml:"CARRIER_ADDRESSES"`
	CARRIERADDRESS []*CARRIERADDRESS `xml:"CARRIER_ADDRESS"`
}

// CARRIERPHONE ...
type CARRIERPHONE struct {
	XMLName        xml.Name `xml:"CARRIER_PHONE"`
	CARRIERPHNNUM  string   `xml:"CARRIER_PHN_NUM"`
	CARRIERPHNTYPE string   `xml:"CARRIER_PHN_TYPE"`
}

// CARRIERPHONES ...
type CARRIERPHONES struct {
	XMLName      xml.Name        `xml:"CARRIER_PHONES"`
	CARRIERPHONE []*CARRIERPHONE `xml:"CARRIER_PHONE"`
}

// CARRIEREMAIL ...
type CARRIEREMAIL struct {
	XMLName          xml.Name `xml:"CARRIER_EMAIL"`
	CARRIEREMAILADDR string   `xml:"CARRIER_EMAIL_ADDR"`
	CARRIEREMAILTYPE string   `xml:"CARRIER_EMAIL_TYPE"`
}

// CARRIEREMAILS ...
type CARRIEREMAILS struct {
	XMLName      xml.Name        `xml:"CARRIER_EMAILS"`
	CARRIEREMAIL []*CARRIEREMAIL `xml:"CARRIER_EMAIL"`
}

// CARRIERWEBLINK ...
type CARRIERWEBLINK struct {
	XMLName            xml.Name `xml:"CARRIER_WEB_LINK"`
	CARRIERWEBLINKURL  string   `xml:"CARRIER_WEB_LINK_URL"`
	CARRIERWEBLINKTYPE string   `xml:"CARRIER_WEB_LINK_TYPE"`
}

// CARRIERWEBLINKS ...
type CARRIERWEBLINKS struct {
	XMLName        xml.Name          `xml:"CARRIER_WEB_LINKS"`
	CARRIERWEBLINK []*CARRIERWEBLINK `xml:"CARRIER_WEB_LINK"`
}

// CARRIER ...
type CARRIER struct {
	CARRIERCONT         string            `xml:"CARRIER_CONT"`
	CARRIERID           string            `xml:"CARRIER_ID"`
	CARRIERDISPLAYNAME  string            `xml:"CARRIER_DISPLAY_NAME"`
	CARRIERBUSINESSNAME string            `xml:"CARRIER_BUSINESS_NAME"`
	CARRIERADDRESSES    *CARRIERADDRESSES `xml:"CARRIER_ADDRESSES"`
	CARRIERPHONES       *CARRIERPHONES    `xml:"CARRIER_PHONES"`
	CARRIEREMAILS       *CARRIEREMAILS    `xml:"CARRIER_EMAILS"`
	CARRIERWEBLINKS     *CARRIERWEBLINKS  `xml:"CARRIER_WEB_LINKS"`
	CARRIEROFFICEHOURS  string            `xml:"CARRIER_OFFICE_HOURS"`
	CARRIEROFFICEDAYS   string            `xml:"CARRIER_OFFICE_DAYS"`
	CARRIERBUDGETCNTR   string            `xml:"CARRIER_BUDGET_CNTR"`
}

// ADDRESS ...
type ADDRESS struct {
	ADDRLINE1     string `xml:"ADDR_LINE1"`
	ADDRLINE2     string `xml:"ADDR_LINE2"`
	ADDRLINE3     string `xml:"ADDR_LINE3"`
	ADDRCITY      string `xml:"ADDR_CITY"`
	ADDRSTATE     string `xml:"ADDR_STATE"`
	ADDRZIP       string `xml:"ADDR_ZIP"`
	ADDRCNTRYCODE string `xml:"ADDR_CNTRY_CODE"`
	ADDRTYPE      string `xml:"ADDR_TYPE"`
	ADDRSTATUS    string `xml:"ADDR_STATUS"`
	ADDRPREFIND   string `xml:"ADDR_PREF_IND"`
}

// ADDRESSES ...
type ADDRESSES struct {
	ADDRESS *ADDRESS `xml:"ADDRESS"`
}

// EMAILADDRESS ...
type EMAILADDRESS struct {
	XMLName   xml.Name `xml:"EMAIL_ADDRESS"`
	EMAILADDR string   `xml:"EMAIL_ADDR"`
	EMAILTYPE string   `xml:"EMAIL_TYPE"`
}

// EMAILADDRESSES ...
type EMAILADDRESSES struct {
	XMLName      xml.Name      `xml:"EMAIL_ADDRESSES"`
	EMAILADDRESS *EMAILADDRESS `xml:"EMAIL_ADDRESS"`
}

// PARTYPHONE ...
type PARTYPHONE struct {
	XMLName   xml.Name `xml:"PARTY_PHONE"`
	PHNNUM    string   `xml:"PHN_NUM"`
	PHNTYPE   string   `xml:"PHN_TYPE"`
	PREFERRED string   `xml:"PREFERRED"`
}

// PARTYPHONES ...
type PARTYPHONES struct {
	XMLName    xml.Name    `xml:"PARTY_PHONES"`
	PARTYPHONE *PARTYPHONE `xml:"PARTY_PHONE"`
}

// PARTY ...
type PARTY struct {
	PARTYCONT         string            `xml:"PARTY_CONT"`
	PARTYID           string            `xml:"PARTY_ID"`
	PARTYROLEID       int               `xml:"PARTY_ROLE_ID"`
	PARTYTYPE         string            `xml:"PARTY_TYPE"`
	PARTYROLENAME     string            `xml:"PARTY_ROLE_NAME"`
	PARTYRELTOINS     string            `xml:"PARTY_REL_TO_INS"`
	PARTYPCT          int               `xml:"PARTY_PCT"`
	PARTYFULLNAME     string            `xml:"PARTY_FULL_NAME"`
	PARTYPREFIX       string            `xml:"PARTY_PREFIX"`
	PARTYFSTNAME      string            `xml:"PARTY_FST_NAME"`
	PARTYMI           string            `xml:"PARTY_MI"`
	PARTYLSTNAME      string            `xml:"PARTY_LST_NAME"`
	PARTYSUFFIX       string            `xml:"PARTY_SUFFIX"`
	ADDRESSES         []*ADDRESSES      `xml:"ADDRESSES"`
	EMAILADDRESSES    []*EMAILADDRESSES `xml:"EMAIL_ADDRESSES"`
	PARTYPHONES       []*PARTYPHONES    `xml:"PARTY_PHONES"`
	PARTYDOB          string            `xml:"PARTY_DOB"`
	PARTYATTAINAGE    int               `xml:"PARTY_ATTAIN_AGE"`
	PARTYGENDER       string            `xml:"PARTY_GENDER"`
	PARTYDELIVERYFLAG string            `xml:"PARTY_DELIVERY_FLAG"`
	PARTYDELIVERYDESC string            `xml:"PARTY_DELIVERY_DESC"`
}

// POLICYVALUE ...
type POLICYVALUE struct {
	PVALCONT                 string  `xml:"PVAL_CONT"`
	PVALBEGINDATE            string  `xml:"PVAL_BEGIN_DATE"`
	PVALENDDATE              string  `xml:"PVAL_END_DATE"`
	PVALBEGINCOVERVAL        float64 `xml:"PVAL_BEGIN_COVER_VAL"`
	PVALENDCOVERVAL          float64 `xml:"PVAL_END_COVER_VAL"`
	PVALCHGCOVERVAL          float64 `xml:"PVAL_CHG_COVER_VAL"`
	PVALBEGINDEATHBNFTVAL    float64 `xml:"PVAL_BEGIN_DEATH_BNFT_VAL"`
	PVALENDDEATHBNFTVAL      float64 `xml:"PVAL_END_DEATH_BNFT_VAL"`
	PVALCHGDEATHBNFTVAL      float64 `xml:"PVAL_CHG_DEATH_BNFT_VAL"`
	PVALBEGINSURRVAL         float64 `xml:"PVAL_BEGIN_SURR_VAL"`
	PVALENDSURRVAL           float64 `xml:"PVAL_END_SURR_VAL"`
	PVALCHGSURRVAL           float64 `xml:"PVAL_CHG_SURR_VAL"`
	PVALBEGINACCTVAL         float64 `xml:"PVAL_BEGIN_ACCT_VAL"`
	PVALENDACCTVAL           float64 `xml:"PVAL_END_ACCT_VAL"`
	PVALCHGACCTVAL           float64 `xml:"PVAL_CHG_ACCT_VAL"`
	PVALTOTALPAYRCVD         float64 `xml:"PVAL_TOTAL_PAY_RCVD"`
	PVALTOTALCOICOST         float64 `xml:"PVAL_TOTAL_COI_COST"`
	PVALCOSTOFINSVAL         float64 `xml:"PVAL_COST_OF_INS_VAL"`
	PVALTOTALUNITCHARGE      float64 `xml:"PVAL_TOTAL_UNIT_CHARGE"`
	PVALTOTALEXPENSECHARGE   float64 `xml:"PVAL_TOTAL_EXPENSE_CHARGE"`
	PVALTOTALPARTWITHDRAWVAL float64 `xml:"PVAL_TOTAL_PART_WITHDRAW_VAL"`
	PVALTOTALLOANAMT         float64 `xml:"PVAL_TOTAL_LOAN_AMT"`
	PVALTOTALLOANBALNAMT     float64 `xml:"PVAL_TOTAL_LOAN_BALN_AMT"`
	PVALTOTALLOANREPAYAMT    float64 `xml:"PVAL_TOTAL_LOAN_REPAY_AMT"`
	PVALTOTALINTCHARGE       float64 `xml:"PVAL_TOTAL_INT_CHARGE"`
	PVALTOTALINTCREDIT       float64 `xml:"PVAL_TOTAL_INT_CREDIT"`
	PVALINTCREDITCURRRATE    int     `xml:"PVAL_INT_CREDIT_CURR_RATE"`
	PVALLOANINTCURRRATE      int     `xml:"PVAL_LOAN_INT_CURR_RATE"`
	PVALLOANCREDITCURRRATE   int     `xml:"PVAL_LOAN_CREDIT_CURR_RATE"`
	PVALINTCREDITNEXTYRRATE  int     `xml:"PVAL_INT_CREDIT_NEXT_YR_RATE"`
	PVALLOANINTNEXTYRRATE    int     `xml:"PVAL_LOAN_INT_NEXT_YR_RATE"`
	PVALLOANCREDITNEXTYRRATE int     `xml:"PVAL_LOAN_CREDIT_NEXT_YR_RATE"`
	PVALCURRPAYAMT           float64 `xml:"PVAL_CURR_PAY_AMT"`
	PVALCURRPAYMODE          string  `xml:"PVAL_CURR_PAY_MODE"`
	PVALANNLPSAMT            float64 `xml:"PVAL_ANN_LPS_AMT"`
}

// PARTIES ...
type PARTIES struct {
	PARTY []*PARTY `xml:"PARTY"`
}

// POLICY ...
type POLICY struct {
	RECTYPE             string       `xml:"REC_TYPE"`
	POLCONT             string       `xml:"POL_CONT"`
	POLPLANCODE         string       `xml:"POL_PLAN_CODE"`
	POLPRODUCT          string       `xml:"POL_PRODUCT"`
	POLQUALTYPE         string       `xml:"POL_QUAL_TYPE"`
	POLPLANTYPE         string       `xml:"POL_PLAN_TYPE"`
	POLPRODSHORTNAME    string       `xml:"POL_PROD_SHORT_NAME"`
	POLPRDCTMKTGNAME    string       `xml:"POL_PRDCT_MKTG_NAME"`
	POLCYCLEDATE        string       `xml:"POL_CYCLE_DATE"`
	POLSERVICENAME      string       `xml:"POL_SERVICE_NAME"`
	POLSERVICEDESC      string       `xml:"POL_SERVICE_DESC"`
	POLSTATUS           string       `xml:"POL_STATUS"`
	POLLOB              string       `xml:"POL_LOB"`
	POLPRCSGCOMP        string       `xml:"POL_PRCSG_COMP"`
	POLDOCTYPE          string       `xml:"POL_DOC_TYPE"`
	POLPRDCTCOMP        string       `xml:"POL_PRDCT_COMP"`
	POLISSUEDATE        string       `xml:"POL_ISSUE_DATE"`
	POLISSUESTATE       string       `xml:"POL_ISSUE_STATE"`
	POLCURRYR           int          `xml:"POL_CURR_YR"`
	POLSYSCODE          string       `xml:"POL_SYS_CODE"`
	POLPRODUCTTYPE      string       `xml:"POL_PRODUCT_TYPE"`
	POLSTARTDATE        string       `xml:"POL_START_DATE"`
	POLMATURITYDATE     string       `xml:"POL_MATURITY_DATE"`
	POLTERMDATE         string       `xml:"POL_TERM_DATE"`
	POLRESIDENCESTATE   string       `xml:"POL_RESIDENCE_STATE"`
	POLTERM             int          `xml:"POL_TERM"`
	POLPREMAMT          float64      `xml:"POL_PREM_AMT"`
	POLPREMMODE         string       `xml:"POL_PREM_MODE"`
	POLRISKCLASS        string       `xml:"POL_RISK_CLASS"`
	POLEFFECTIVEDATE    string       `xml:"POL_EFFECTIVE_DATE"`
	POLISSUEAGE         int          `xml:"POL_ISSUE_AGE"`
	POLSTARTANNIVDATE   string       `xml:"POL_START_ANNIV_DATE"`
	POLCURRANNIVENDDATE string       `xml:"POL_CURR_ANNIV_END_DATE"`
	POLLASTANNIVDATE    string       `xml:"POL_LAST_ANNIV_DATE"`
	POLCURRANNIVDATE    string       `xml:"POL_CURR_ANNIV_DATE"`
	POLNEXTANNIVENDDATE string       `xml:"POL_NEXT_ANNIV_END_DATE"`
	POLFACEVALUE        int          `xml:"POL_FACE_VALUE"`
	POLSOURCE           string       `xml:"POL_SOURCE"`
	PARTIES             *PARTIES     `xml:"PARTIES"`
	POLICYVALUES        *POLICYVALUE `xml:"POLICYVALUES"`
	CARRIER             *CARRIER     `xml:"CARRIER"`
}

// POLICIES ...
type POLICIES struct {
	POLICY []*POLICY `xml:"POLICY"`
}

// DataServicesLetter ...
type DataServicesLetter struct {
	POLICIES *POLICIES `xml:"POLICIES"`
}
