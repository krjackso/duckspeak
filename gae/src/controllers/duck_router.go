package controllers

import (
	"github.com/gorilla/mux"
	"net/http"

	"util/auth"
)

type DuckRouter struct {
	HttpScheme string
	Router     *mux.Router
}

func NewDuckRouter(httpScheme string, authenticator *auth.Authenticator) *DuckRouter {
	gorilla := mux.NewRouter().Host("{host:.*}").Schemes("{scheme}").Subrouter()

	r := &DuckRouter{
		HttpScheme: httpScheme,
		Router:     gorilla,
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
	path, err := self.Router.Get(name).URL("host", host)
	if err != nil {
		panic(err)
	}

	path.Scheme = self.HttpScheme
	return path.String()
}
