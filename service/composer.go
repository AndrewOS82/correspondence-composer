package service

import (
	"context"
	"fmt"

	"correspondence-composer/models"
)

type Composer struct {
	DataFetcher  dataFetcher
	RuleExecutor ruleExecutor
	Uploader     uploader
}

type uploader interface {
	Upload(filepath string, key string) error
}

type dataFetcher interface {
	FetchAnniversaryData(ctx context.Context, policyNumber string) (*models.AnniversaryStatement, error)
}

type ruleExecutor interface {
	ValidateAnniversaryData(data *models.AnniversaryStatement) ([]*models.RuleValidation, error)
}

func (c *Composer) RunProcess(ctx context.Context, event *models.KafkaEvent) error {
	switch event.Name {
	case "anniversary":
		policy := event.PolicyData
		anniversaryData, err := c.DataFetcher.FetchAnniversaryData(ctx, policy.PolicyNumber)
		if err != nil {
			fmt.Printf("Error %v\n", err)
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

		// write & save file w/ anniversaryData snapshot regardless of XML creation
		filename := "test.json"
		key := "anniversary-data/<policyNumber>-<timestamp>"
		uploadErr := c.Uploader.Upload(filename, key)
		if err != nil {
			// handle
			return uploadErr
		}

		// If all the rules pass, we proceed to building XML for that correspondence type
		// c.XmlGenerator.GenerateXml(anniversaryData)

		return nil
	}

	return nil
}
