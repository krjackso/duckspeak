package controllers

import (
	"google.golang.org/appengine"
	"net/http"

	"models"
	"util/auth"
)

type IAuthController interface {
	RefreshAuth(http.ResponseWriter, *http.Request)
}

type AuthController struct {
	Authenticator *auth.Authenticator
}

type RefreshAuthRequest struct {
	DeviceId string `json:"deviceId"`
}

type RefreshAuthResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
}

func (self *AuthController) RefreshAuth(w http.ResponseWriter, r *http.Request) {
	var requestData RefreshAuthRequest
	if err := parseJson(r, &requestData); err != nil {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	ctx := appengine.NewContext(r)
	device, err := models.GetDevice(ctx, requestData.DeviceId)

	if err != nil {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	if device == nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	token, expiresIn := self.Authenticator.NewAccessToken(device.Id)

	response := &RefreshAuthResponse{
		AccessToken: token,
		ExpiresIn:   expiresIn,
	}
	jsonResponse(w, response, http.StatusOK)
}
