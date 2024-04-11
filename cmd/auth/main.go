package main

import (
	"go-temp/pkg/router"
	"net/http"
)

func main() {
	router := router.NewRouter("chi", ":2222")

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from chi"))
	})

	router.Listen()
}
