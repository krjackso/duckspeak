package duckspeak

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", quack)
}

func quack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Quack quack!")
}
