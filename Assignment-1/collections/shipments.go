package collections

import (
	"context"
	"errors"
	"sellerapp/base/db/mongodb"
	"time"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
)

type ShipmentDetails struct {
	ID          objectid.ObjectID `bson:"_id,omitempty"`
	Name        string            `bson:"name"`
	CompanyName string            `bson:"company_name"`
	Address1    string            `bson:"address1"`
	Town        string            `bson:"town"`
	PostCode    string            `bson:"post_code"`
	IsoCountry  string            `bson:"iso_country"`
	Version     string            `bson:"version"`
	CaptureTime time.Time         `bson:"capture_time"`
}

type CarrierDetails struct {
	ID          objectid.ObjectID `bson:"_id,omitempty"`
	Code        string            `bson:"name"`
	Service     string            `bson:"company_name"`
	Version     string            `bson:"version"`
	CaptureTime time.Time         `bson:"capture_time"`
}

type AllShipments struct {
	ID            objectid.ObjectID `bson:"_id,omitempty"`
	SourceOrderId string            `bson:"source_order_id"`
	ShipToId      objectid.ObjectID `bson:"shipto_id"`
	CarrierId     objectid.ObjectID `bson:"carrier_id"`
	Version       string            `bson:"version"`
	CaptureTime   time.Time         `bson:"capture_time"`
}

func InsertShipmentDetails(data ShipmentDetails) (objectid.ObjectID, error) {
	var oid objectid.ObjectID
	var ok bool

	if res, err := mongodb.CreateCollection("shipments_details").InsertOne(context.Background(), data); err != nil {
		return oid, err
	} else {
		oid, ok = res.InsertedID.(objectid.ObjectID)
		if !ok {
			return oid, errors.New("Something went wrong.")
		}
	}
	return oid, nil

}

func InsertCarrier(code, service, version string) (objectid.ObjectID, error) {
	var oid objectid.ObjectID
	var ok bool
	data := CarrierDetails{
		Code:        code,
		Service:     service,
		Version:     version,
		CaptureTime: time.Now(),
	}

	if res, err := mongodb.CreateCollection("all_carriers").InsertOne(context.Background(), data); err != nil {
		return oid, err
	} else {
		oid, ok = res.InsertedID.(objectid.ObjectID)
		if !ok {
			return oid, errors.New("Something went wrong.")
		}
	}
	return oid, nil
}

func InsertAllShipments(shiptoId, carrierId objectid.ObjectID, orderId, version string) (objectid.ObjectID, error) {
	var oid objectid.ObjectID
	var ok bool
	data := AllShipments{
		SourceOrderId: orderId,
		ShipToId:      shiptoId,
		CarrierId:     carrierId,
		Version:       version,
		CaptureTime:   time.Now(),
	}

	if res, err := mongodb.CreateCollection("all_shipments").InsertOne(context.Background(), data); err != nil {
		return oid, err
	} else {
		oid, ok = res.InsertedID.(objectid.ObjectID)
		if !ok {
			return oid, errors.New("Something went wrong.")
		}
	}
	return oid, nil
}
