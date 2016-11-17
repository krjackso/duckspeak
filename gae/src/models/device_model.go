package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"
)

type Device struct {
	Id      string `datastore:"-"`
	Created time.Time
	Type    string
}

func deviceKey(ctx context.Context, deviceId string) *datastore.Key {
	return datastore.NewKey(ctx, "Device", deviceId, 0, nil)
}

func GetDevice(ctx context.Context, deviceId string) (*Device, error) {
	key := deviceKey(ctx, deviceId)

	device := &Device{}
	err := datastore.Get(ctx, key, device)

	switch err {
	case datastore.ErrNoSuchEntity:
		return nil, nil
	case nil:
		break
	default:
		return nil, err
	}

	device.Id = deviceId
	return device, nil
}

func NewDevice(ctx context.Context, deviceType string) (*Device, error) {
	newId := uuid.NewV4().String()
	key := deviceKey(ctx, newId)

	device := &Device{
		Type:    deviceType,
		Created: time.Now().UTC(),
	}

	key, err := datastore.Put(ctx, key, device)
	if err != nil {
		return nil, err
	}

	device.Id = key.StringID()
	return device, nil
}
