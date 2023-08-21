# Correspondence Orchestrator

## Run Instructions

In order to run the report and generate sample output, run the following code: 
> go run main.go

## Regenerating types from XSDs

The following snippet will regenerate your Go types from your xsds. The existing ones were regenerated as follows using [xgen](https://github.com/xuri/xgen)

> xgen -i ./xsds/PurchaseOrderSchema.xsd -o ./schemas/purchase_order -l Go

This is an update to the README.
