package controllers

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
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
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	token, expiresIn := self.Authenticator.NewToken(device.Id, auth.AccessAudience)

	response := &RefreshAuthResponse{
		AccessToken: token,
		ExpiresIn:   expiresIn,
	}
	jsonResponse(w, response, http.StatusOK)
}

type AdminAuthResponse struct {
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
}

func (self *AuthController) AdminAuth(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	adminUser := user.Current(ctx)
	if adminUser == nil || !adminUser.Admin {
		http.Error(w, "", http.StatusForbidden)
		return
	}

	token, expiresIn := self.Authenticator.NewToken(adminUser.Email, auth.AdminAudience)

	response := &AdminAuthResponse{
		Email:       adminUser.Email,
		AccessToken: token,
		ExpiresIn:   expiresIn,
	}
	jsonResponse(w, response, http.StatusOK)
}
