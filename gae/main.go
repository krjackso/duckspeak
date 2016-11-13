package duckspeak

import (
	"controllers"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", quack)
	http.HandleFunc("/cron/", cron)
	http.HandleFunc("/bootstrap", controllers.Bootstrap)
}

func quack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Quack quack!")
}

func cron(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "cron cron!")
}
