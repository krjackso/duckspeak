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
	DailyTopic  string `json:"dailyTopic"`
}

func (self *BootstrapController) Bootstrap(w http.ResponseWriter, r *http.Request) {
	response := &BootstrapResponse{
		NewDevice:   self.Router.GetHref(r, "newDevice"),
		RefreshAuth: self.Router.GetHref(r, "refreshAuth"),
		DailyTopic:  self.Router.GetHref(r, "dailyTopic"),
	}
	jsonResponse(w, response, http.StatusOK)
}
