package duckspeak

import (
	"google.golang.org/appengine"
	"net/http"
	"os"

	"controllers"
	"util/auth"
)

func init() {
	authIssuer := os.Getenv("AUTH_ISSUER")
	authSecret := os.Getenv("AUTH_SECRET")

	switch "" {
	case authIssuer:
		fallthrough
	case authSecret:
		panic("Missing environment variables")
	}

	authenticator := auth.NewAuthenticator(authIssuer, authSecret)

	httpScheme := determineHttpScheme()
	duck := controllers.NewDuckRouter(httpScheme, authenticator)

	http.Handle("/", duck.Router)
}

func determineHttpScheme() string {
	if appengine.IsDevAppServer() {
		return "http"
	} else {
		return "https"
	}
}
