package controllers

import (
	"net/http"
	"os"
)

type BootstrapResponse struct {
	GetToken string `json:"getToken"`
	GetTopic string `json:"getTopic"`
}

var apiBase = os.Getenv("API_BASE")

func Bootstrap(w http.ResponseWriter, r *http.Request) {
	response := &BootstrapResponse{
		GetToken: apiBase + "/token",
		GetTopic: apiBase + "/dailyTopic",
	}
	jsonResponse(w, response, http.StatusTeapot)
}
