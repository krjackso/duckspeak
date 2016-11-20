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

	topicController := &TopicController{
		Router:        r,
		Authenticator: authenticator,
	}

	gorilla.HandleFunc("/devices", NewDevice).Methods("PUT").Name("newDevice")
	gorilla.HandleFunc("/bootstrap", bootstrapController.Bootstrap).Methods("GET")
	gorilla.HandleFunc("/auth/refresh", authController.RefreshAuth).Methods("POST").Name("refreshAuth")
	gorilla.HandleFunc("/auth/admin", authController.AdminAuth).Methods("GET")

	gorilla.HandleFunc("/dailyTopic", topicController.GetDailyTopic).Methods("GET").Name("dailyTopic")
	gorilla.HandleFunc("/topic", topicController.ListTopics).Methods("GET")
	gorilla.HandleFunc("/topic", topicController.NewTopic).Methods("PUT")
	gorilla.HandleFunc("/topic/{topicId}", topicController.GetTopic).Methods("GET").Name("getTopic")

	return r
}

func (self *DuckRouter) GetHref(r *http.Request, name string, params ...string) string {
	host := r.Host

	params = append(params, "host", host)

	path, err := self.Router.Get(name).URL(params...)
	if err != nil {
		panic(err)
	}

	path.Scheme = self.HttpScheme
	return path.String()
}

func (self *DuckRouter) GetVar(r *http.Request, name string) string {
	vars := mux.Vars(r)
	return vars[name]
}
