package service

import (
	"context"
	"fmt"

	"correspondence-composer/models"
	"correspondence-composer/utils/log"
)

type Composer struct {
	Logger       log.Logger
	DataFetcher  dataFetcher
	RuleExecutor ruleExecutor
	Uploader     uploader
	XMLBuilder   xmlBuilder
}

type uploader interface {
	Upload(filepath string, key string) error
}

type dataFetcher interface {
	FetchAnniversaryData(ctx context.Context, policyNumber string) (*models.AnniversaryStatement, error)
}

type ruleExecutor interface {
	ValidateAnniversaryData(ctx context.Context, data *models.AnniversaryStatement) ([]*models.RuleValidation, error)
}

type xmlBuilder interface {
	BuildDataServicesLetter(statement *models.AnniversaryStatement) (string, error)
}

func (c *Composer) RunProcess(ctx context.Context, event *models.KafkaEvent) error {
	switch event.Name {
	case "anniversary":
		policy := event.PolicyData
		anniversaryData, err := c.DataFetcher.FetchAnniversaryData(ctx, policy.PolicyNumber)
		if err != nil {
			c.Logger.ErrorWithFields(err, log.Fields{
				"policyNumber": policy.PolicyNumber,
			})
			return err
		}

		validationFailures, err := c.RuleExecutor.ValidateAnniversaryData(ctx, anniversaryData)
		if err != nil {
			c.Logger.ErrorWithFields(err, log.Fields{
				"policyNumber": policy.PolicyNumber,
			})
			return err
		}
		if len(validationFailures) > 0 {
			c.Logger.InfoWithFields("Validations failed - no XML generated", log.Fields{
				"failedRule": validationFailures[0].RuleName,
			})
			// Save data in S3 and the reason why data is not valid and XML can't be generated
			// Produce kafka status message with details on which rules failed.
			return nil
		}

		// write & save file w/ anniversaryData snapshot regardless of XML creation
		filename := "test.json"
		key := "anniversary-data/<policyNumber>-<timestamp>"
		uploadErr := c.Uploader.Upload(filename, key)
		if err != nil {
			// handle
			return uploadErr
		}

		// If all the rules pass, we proceed to building XML for that correspondence type

		preparedXML, err := c.XMLBuilder.BuildDataServicesLetter(anniversaryData)

		if err != nil {
			fmt.Printf("Error %v\n", err)
			return err
		}

		fmt.Printf("Prepared XML %v\n", preparedXML)

		// Once we build the xml, we ensure it is valid

		// Once we ensure it is valid, we send it to CDS

		return nil
	}

	return nil
}
