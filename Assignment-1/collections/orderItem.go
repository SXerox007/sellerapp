package collections

import (
	"context"
	"errors"
	"sellerapp/base/db/mongodb"
	"time"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
)

type OrderItem struct {
	ID            objectid.ObjectID `bson:"_id,omitempty"`
	SourceOrderId string            `bson:"source_order_id"`
	Sku           string            `bson:"sku"`
	SourceItemId  string            `bson:"source_item_id"`
	Version       string            `bson:"version"`
	CaptureTime   time.Time         `bson:"capture_time"`
}

func InsertOrderItem(orderId, sku, itemId, version string) (objectid.ObjectID, error) {
	var oid objectid.ObjectID
	var ok bool
	data := OrderItem{
		SourceOrderId: orderId,
		Sku:           sku,
		SourceItemId:  itemId,
		Version:       version,
		CaptureTime:   time.Now(),
	}

	if res, err := mongodb.CreateCollection("all_orders_items").InsertOne(context.Background(), data); err != nil {
		return oid, err
	} else {
		oid, ok = res.InsertedID.(objectid.ObjectID)
		if !ok {
			return oid, errors.New("Something went wrong.")
		}

	}
	return oid, nil
}
