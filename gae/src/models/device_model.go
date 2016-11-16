package models

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"
)

type Device struct {
	Id      int64 `datastore:"-"`
	Created time.Time
	Type    string
}

func NewDevice(ctx context.Context, deviceType string) (*Device, error) {
	key := datastore.NewIncompleteKey(ctx, "Device", nil)

	device := &Device{
		Type:    deviceType,
		Created: time.Now().UTC(),
	}

	key, err := datastore.Put(ctx, key, device)
	if err != nil {
		return nil, err
	}

	device.Id = key.IntID()
	return device, nil
}
