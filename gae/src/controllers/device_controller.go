package controllers

import (
	"google.golang.org/appengine"
	"net/http"
	"strings"

	"models"
)

type DeviceController struct{}

type NewDeviceRequest struct {
	DeviceType string `json:"deviceType"`
}

type NewDeviceResponse struct {
	DeviceId int64 `json:"deviceId"`
}

func (self *DeviceController) NewDevice(w http.ResponseWriter, r *http.Request) {
	var requestData NewDeviceRequest
	if err := parseJson(r, &requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deviceType := strings.ToLower(requestData.DeviceType)

	if !validateDeviceType(deviceType) {
		http.Error(w, "invalid deviceType", http.StatusBadRequest)
		return
	}

	ctx := appengine.NewContext(r)
	device, err := models.NewDevice(ctx, deviceType)

	if err != nil {
		panic(err)
	}

	response := &NewDeviceResponse{
		DeviceId: device.Id,
	}
	jsonResponse(w, response, http.StatusTeapot)
}

func validateDeviceType(deviceType string) bool {
	switch deviceType {
	case "ios", "android":
		return true
	default:
		return false
	}
}
