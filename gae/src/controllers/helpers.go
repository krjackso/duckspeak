package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"util/auth"
)

func jsonResponse(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func parseJson(r *http.Request, dest interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, dest)
	if err != nil {
		return err
	}

	return nil
}

func withAccessAuth(authenticator *auth.Authenticator, w http.ResponseWriter, r *http.Request, next func(id string)) {
	withAuth(authenticator, auth.AccessAudience, w, r, next)
}

func withAdminAuth(authenticator *auth.Authenticator, w http.ResponseWriter, r *http.Request, next func(id string)) {
	withAuth(authenticator, auth.AdminAudience, w, r, next)
}

func withAuth(authenticator *auth.Authenticator, audience string, w http.ResponseWriter, r *http.Request, next func(id string)) {
	bearer, ok := auth.BearerFromHeader(r.Header)

	if !ok {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	id, valid := authenticator.VerifyToken(bearer, audience)

	if !valid {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	next(id)
}
