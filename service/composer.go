package service

import (
	"correspondence-composer/models"
)

type Composer struct {
	DataFetcher  dataFetcher
	RuleExecutor ruleExecutor
}

type dataFetcher interface {
	FetchAnniversaryData() (*models.AnniversaryStatement, error)
}

type ruleExecutor interface {
	ValidateAnniversaryData(data *models.AnniversaryStatement) ([]*models.RuleValidation, error)
}

func (c *Composer) RunProcess(correspondenceType string) error {
	switch correspondenceType {
	case "anniversary":
		anniversaryData, err := c.DataFetcher.FetchAnniversaryData()
		if err != nil {
			// handle
			return err
		}

		validationFailures, err := c.RuleExecutor.ValidateAnniversaryData(anniversaryData)
		if err != nil {
			// handle
			return err
		}
		if len(validationFailures) > 0 {
			// Save data in S3 and the reason why data is not valid and XML can't be generated
			// Produce kafka status message with details on which rules failed.
			return nil
		}
		// If all the rules pass, we proceed to building XML for that correspondence type
		// c.XmlGenerator.GenerateXml(anniversaryData)

		return nil
	}

	return nil
}
