package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

var apiBase = os.Getenv("API_BASE")

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

type DuckRouter struct {
	Router *mux.Router
}

func NewDuckRouter() *DuckRouter {
	gorilla := mux.NewRouter()

	r := &DuckRouter{
		Router: gorilla,
	}

	deviceController := &DeviceController{}

	bootstrapController := &BootstrapController{
		Router: r,
	}

	gorilla.HandleFunc("/devices", deviceController.NewDevice).Methods("PUT").Name("newDevice")
	gorilla.HandleFunc("/bootstrap", bootstrapController.Bootstrap).Methods("GET")

	return r
}

func (self *DuckRouter) GetHref(name string) string {
	path, err := self.Router.Get(name).URLPath()
	if err != nil {
		panic(err)
	}
	return apiBase + path.String()
}
