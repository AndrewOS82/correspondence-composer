package main

import (
	"encoding/xml"
	"fmt"

	rulesgateway "correspondence-composer/gateways/rulesengine"
	. "correspondence-composer/schemas" //nolint:all
	"correspondence-composer/service"
	"correspondence-composer/usecases"
)

func main() {
	rulesEngineGateway := rulesgateway.New()
	composer := service.Composer{
		DataFetcher: &usecases.DataFetcher{},
		RuleExecutor: &usecases.RuleExecutor{
			RulesEngine: rulesEngineGateway,
		},
		// XmlGenerator: &usecases.XmlGenerator{},
	}

	// This is a placeholder. The correspondence type will be determined by incoming kafka message.
	correspondenceType := "anniversary"

	fmt.Println("Running composer process...")
	err := composer.RunProcess(correspondenceType)
	if err != nil {
		fmt.Printf("Error running composer process: %v\n", err)
	}
	fmt.Println("Composer finished successfully")

	// The code below is part of the POC and it will be moved into an XML generator
	// with relevant types like Policy replacing the example PurchaseOrder.
	shipAddress := USAddress{
		CountryAttr: "USA",
		Name:        "Clark Kent",
		Street:      "123 Main Street",
		City:        "Metropolis",
		State:       "NJ",
		Zip:         10016,
	}

	billAddress := USAddress{
		CountryAttr: "USA",
		Name:        "Bruce Wayne",
		Street:      "123 Wayne Road",
		City:        "Gotham",
		State:       "NY",
		Zip:         19106,
	}

	shippingAddresses := []*USAddress{&shipAddress}

	purchaseOrder := PurchaseOrderType{
		OrderDateAttr: "01/01/2023",
		ShipTo:        shippingAddresses,
		BillTo:        &billAddress,
	}

	finalOrder := PurchaseOrder(&purchaseOrder)

	out, _ := xml.MarshalIndent(finalOrder, " ", "  ")

	println(string(out))

}
