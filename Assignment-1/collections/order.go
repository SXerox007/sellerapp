package collections

import (
	"context"
	"errors"
	"log"
	"sellerapp/Assignment-1/models"
	"sellerapp/base/db/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	objectid "go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            objectid.ObjectID `bson:"_id,omitempty"`
	SourceOrderId string            `bson:"source_order_id"`
	Version       string            `bson:"version"`
	CaptureTime   time.Time         `bson:"capture_time"`
}

func InsertNewOrder(orderId, version string) (objectid.ObjectID, error) {
	var oid objectid.ObjectID
	var ok bool
	data := Order{
		SourceOrderId: orderId,
		Version:       version,
		CaptureTime:   time.Now(),
	}

	if res, err := mongodb.CreateCollection("all_orders").InsertOne(context.Background(), data); err != nil {
		return oid, err
	} else {
		oid, ok = res.InsertedID.(objectid.ObjectID)
		if !ok {
			return oid, errors.New("Something went wrong.")
		}
	}
	return oid, nil

}

type OrderDetails struct {
	ID            objectid.ObjectID  `bson:"_id,omitempty"`
	SourceOrderId string             `bson:"source_order_id"`
	Items         []models.Items     `bson:"items"`
	Shipments     []models.Shipments `bson:"shipments"`
	Version       string             `bson:"version"`
	CaptureTime   time.Time          `bson:"capture_time"`
}

func InsertOrUpdateOrder(orderId, version string, items []models.Items, shipments []models.Shipments) (objectid.ObjectID, error) {
	var oid objectid.ObjectID
	var ok bool
	data := OrderDetails{
		SourceOrderId: orderId,
		Items:         items,
		Shipments:     shipments,
		Version:       version,
		CaptureTime:   time.Now(),
	}

	// add filter to find the same order id
	filter := bson.M{"source_order_id": orderId}
	err := mongodb.CreateCollection("all_orders_details").FindOne(context.Background(), filter).Decode(&OrderDetails{})
	if err == nil {
		// already exist update the order details
		updateFilter := bson.M{"$set": bson.M{"items": items, "shipments": shipments, "capture_time": time.Now()}}
		_, err = mongodb.CreateCollection("all_orders_details").UpdateOne(context.Background(), filter, updateFilter)

	} else {
		if res, err := mongodb.CreateCollection("all_orders_details").InsertOne(context.Background(), data); err != nil {
			return oid, err
		} else {
			log.Println("Error in Insert Data in all_orders_details Collection:", err)
			oid, ok = res.InsertedID.(objectid.ObjectID)
			if !ok {
				return oid, errors.New("Something went wrong.")
			}
		}
	}
	return oid, nil

}
