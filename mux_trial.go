package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/garrypolley/mux_trial/handlers"

	"github.com/gorilla/mux"
)

const (
	port    = 8080
	timeout = 5
)

func handleFunc(r *mux.Router, url string, handler func(http.ResponseWriter, *http.Request)) {
	r.HandleFunc(url, handlers.LogRequest(handler))
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	handleFunc(r, "/", handlers.HomeHandler)
	handleFunc(r, "/products", handlers.ProductsHandler)
	handleFunc(r, "/products/{id:[0-9]+}", handlers.ProductsIdHandler)
	handleFunc(r, "/articles", handlers.ArticlesHandler)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		ReadTimeout:  time.Minute * timeout,
		WriteTimeout: time.Minute * timeout,
	}

	log.Printf("Starting server on :%d with a timeout of %d minutes", port, timeout)
	server.ListenAndServe()
}
