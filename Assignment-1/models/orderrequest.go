package models

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
