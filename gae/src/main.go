package duckspeak

import (
	"controllers"

	"net/http"
)

func init() {
	duck := controllers.NewDuckRouter()

	http.Handle("/", duck.Router)
}
