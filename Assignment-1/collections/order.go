package collections

import (
	"context"
	"errors"
	"sellerapp/base/db/mongodb"
	"time"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
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
