package collections

import (
	"context"
	"errors"
	"sellerapp/base/db/mongodb"
	"time"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
)

type ItemComponent struct {
	ID           objectid.ObjectID `bson:"_id,omitempty"`
	SourceItemId string            `bson:"source_item_id"`
	Code         string            `bson:"code"`
	Fetch        bool              `bson:"fetch"`
	Path         string            `bson:"path"`
	Version      string            `bson:"version"`
	CaptureTime  time.Time         `bson:"capture_time"`
}

func InsertItemComponent(itemId, code, path, version string, fetch bool) (objectid.ObjectID, error) {
	var oid objectid.ObjectID
	var ok bool
	data := ItemComponent{
		SourceItemId: itemId,
		Code:         code,
		Path:         path,
		Fetch:        fetch,
		Version:      version,
		CaptureTime:  time.Now(),
	}

	if res, err := mongodb.CreateCollection("item_components").InsertOne(context.Background(), data); err != nil {
		return oid, err
	} else {
		oid, ok = res.InsertedID.(objectid.ObjectID)
		if !ok {
			return oid, errors.New("Something went wrong.")
		}

	}
	return oid, nil

}
