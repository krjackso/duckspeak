package duckspeak

import (
	"net/http"
	"os"

	"controllers"
	"util/auth"
)

func init() {
	apiBase := os.Getenv("API_BASE")
	authIssuer := os.Getenv("AUTH_ISSUER")
	authSecret := os.Getenv("AUTH_SECRET")

	switch "" {
	case apiBase:
		fallthrough
	case authIssuer:
		fallthrough
	case authSecret:
		panic("Missing environment variables")
	}

	authenticator := auth.NewAuthenticator(authIssuer, authSecret)
	duck := controllers.NewDuckRouter(apiBase, authenticator)

	http.Handle("/", duck.Router)
}
