package controllers

import (
	"net/http"
)

type BootstrapController struct {
	Router *DuckRouter
}

type BootstrapResponse struct {
	NewDevice   string `json:"newDevice"`
	RefreshAuth string `json:"refreshAuth"`
}

func (self *BootstrapController) Bootstrap(w http.ResponseWriter, r *http.Request) {
	response := &BootstrapResponse{
		NewDevice:   self.Router.GetHref(r, "newDevice"),
		RefreshAuth: self.Router.GetHref(r, "refreshAuth"),
	}
	jsonResponse(w, response, http.StatusOK)
}
