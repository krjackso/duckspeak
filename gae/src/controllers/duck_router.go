package controllers

import (
	"github.com/gorilla/mux"
	"net/http"

	"util/auth"
)

type DuckRouter struct {
	ApiBase string
	Router  *mux.Router
}

func NewDuckRouter(apiBase string, authenticator *auth.Authenticator) *DuckRouter {
	gorilla := mux.NewRouter().Host("{host}").Schemes("{scheme:https?}").Subrouter()

	r := &DuckRouter{
		ApiBase: apiBase,
		Router:  gorilla,
	}

	bootstrapController := &BootstrapController{
		Router: r,
	}

	authController := &AuthController{
		Authenticator: authenticator,
	}

	gorilla.HandleFunc("/devices", NewDevice).Methods("PUT").Name("newDevice")
	gorilla.HandleFunc("/bootstrap", bootstrapController.Bootstrap).Methods("GET")
	gorilla.HandleFunc("/auth/refresh", authController.RefreshAuth).Methods("POST").Name("refreshAuth")

	return r
}

func (self *DuckRouter) GetHref(r *http.Request, name string) string {
	host := r.Host
	scheme := r.URL.Scheme
	path, err := self.Router.Get(name).URL("host", host, "scheme", scheme)
	if err != nil {
		panic(err)
	}
	return path.String()
}
