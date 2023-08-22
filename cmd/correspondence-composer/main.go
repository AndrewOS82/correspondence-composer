package main

import (
	"encoding/xml"
	"fmt"

	. "correspondence-composer/schemas" //nolint:all
)

func main() {
	fmt.Println("Hello, world.")
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
