package controllers

import (
	"google.golang.org/appengine"
	"net/http"
	"strings"

	"models"
)

type NewDeviceRequest struct {
	DeviceType string `json:"deviceType"`
}

type NewDeviceResponse struct {
	DeviceId string `json:"deviceId"`
}

func NewDevice(w http.ResponseWriter, r *http.Request) {
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
	jsonResponse(w, response, http.StatusCreated)
}

func validateDeviceType(deviceType string) bool {
	switch deviceType {
	case "ios", "android":
		return true
	default:
		return false
	}
}
