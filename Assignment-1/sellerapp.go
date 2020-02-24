package main

import (
	"log"
	"net/http"
	"sellerapp/Assignment-1/collections"
	"sellerapp/Assignment-1/models"
	"sellerapp/base/utils"
	"time"

	"github.com/gorilla/mux"
)

// any order data post
func OrderDataPost() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		version := mux.Vars(r)["version"]
		var err error
		body := &models.OrderDataRequest{}
		utils.ParseBody(r, body)

		//Enbedded approach
		_, err = collections.InsertOrUpdateOrder(body.OrderData.SourceOrderId, version, body.OrderData.Items, body.OrderData.Shipments)

		// check error if nil
		if err == nil {
			utils.Json("{\"success\":true, \"data\":{\"message\":\"Store Order data with success.\"}}")(w, r)
		} else {
			log.Println("Error:", err)
			utils.Json("{\"success\":false, \"data\":{\"message\":\"Something went wrong.\"}}")(w, r)

		}
	}
}

// refrenced approach
func refrencedApproach(body models.OrderDataRequest, version string) error {
	var err error
	// write into mongodb
	// insert id into all_orders
	_, err = collections.InsertNewOrder(body.OrderData.SourceOrderId, version)
	// insert order items
	for _, item := range body.OrderData.Items {
		_, err = collections.InsertOrderItem(body.OrderData.SourceOrderId, item.Sku, item.SourceItemId, version)
		// insert item components
		for _, row := range item.Components {
			_, err = collections.InsertItemComponent(item.SourceItemId, row.Code, row.Path, version, row.Fetch)
		}
	}
	//insert shipments
	for _, item := range body.OrderData.Shipments {
		// insert shipTo
		shipToId, _ := collections.InsertShipmentDetails(collections.ShipmentDetails{
			Name:        item.ShipTo.Name,
			CompanyName: item.ShipTo.CompanyName,
			Address1:    item.ShipTo.Address1,
			Town:        item.ShipTo.Town,
			PostCode:    item.ShipTo.PostCode,
			IsoCountry:  item.ShipTo.IsoCountry,
			Version:     version,
			CaptureTime: time.Now(),
		})
		// insert carrier
		carrierId, _ := collections.InsertCarrier(item.Carrier.Code, item.Carrier.Service, version)

		// insert all_shipments or sort of mappings of shipments
		_, err = collections.InsertAllShipments(shipToId, carrierId, body.OrderData.SourceOrderId, version)
	}
	return err
}

// any order data get
func OrderDataGet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Json("{\"success\":true, \"data\":{\"message\":\"success.\"}}")(w, r)
	}
}
