package main

import (
	"log"
	"net/http"
	"sellerapp/Assignment-1/collections"
	"sellerapp/base/utils"
	"time"

	"github.com/gorilla/mux"
)

type OrderDataRequest struct {
	OrderData OrderData `json:"orderData" param:"orderData"`
}

type OrderData struct {
	SourceOrderId string      `json:"sourceOrderId" param:"sourceOrderId"`
	Items         []Items     `json:"items" param:"items"`
	Shipments     []Shipments `json:"shipments" param:"shipments"`
}

type Items struct {
	Sku          string       `json:"sku" param:"sku"`
	SourceItemId string       `json:"sourceItemId" param:"sourceItemId"`
	Components   []Components `json:"components" param:"components"`
}

type Components struct {
	Code  string `json:"code" param:"code"`
	Fetch bool   `json:"fetch" param:"fetch"`
	Path  string `json:"path" param:"path"`
}

type Shipments struct {
	ShipTo  ShipTo  `json:"shipTo" param:"shipTo"`
	Carrier Carrier `json:"carrier" param:"carrier"`
}

type ShipTo struct {
	Name        string `json:"name" param:"name"`
	CompanyName string `json:"companyName" param:"companyName"`
	Address1    string `json:"address1" param:"address1"`
	Town        string `json:"town" param:"town"`
	PostCode    string `json:"postcode" param:"postcode"`
	IsoCountry  string `json:"isoCountry" param:"isoCountry"`
}

type Carrier struct {
	Code    string `json:"code" param:"code"`
	Service string `json:"service" param:"service"`
}

// any order data post
func OrderDataPost() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		version := mux.Vars(r)["version"]
		var err error
		body := &OrderDataRequest{}
		utils.ParseBody(r, body)
		log.Println("Body:", body)

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

		// check error if nil
		if err == nil {
			utils.Json("{\"success\":true, \"data\":{\"message\":\"Store Order data with success.\"}}")(w, r)
		} else {
			utils.Json("{\"success\":false, \"data\":{\"message\":\"Something went wrong.\"}}")(w, r)

		}
	}
}
