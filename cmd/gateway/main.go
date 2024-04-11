package main

import (
	"go-temp/pkg/router"
	"net/http"
)

func main() {
	router := router.NewRouter("gin", ":1111")

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Gin"))
	})

	router.Listen()
}
