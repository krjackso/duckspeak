package controllers

import (
	"net/http"
)

type BootstrapController struct {
	Router *DuckRouter
}

type BootstrapResponse struct {
	NewDevice string `json:"newDevice"`
	GetTopic  string `json:"getTopic"`
}

func (self *BootstrapController) Bootstrap(w http.ResponseWriter, r *http.Request) {
	response := &BootstrapResponse{
		NewDevice: self.Router.GetHref("newDevice"),
	}
	jsonResponse(w, response, http.StatusTeapot)
}
