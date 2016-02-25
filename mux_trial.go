package main

import (
	"log"
	"net/http"
	"time"

	"github.com/garrypolley/mux_trial/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", handlers.LogRequest(handlers.HomeHandler))
	r.HandleFunc("/products", handlers.LogRequest(handlers.ProductsHandler))
	r.HandleFunc("/products/{id:[0-9]+}", handlers.LogRequest(handlers.ProductsIdHandler))
	r.HandleFunc("/articles", handlers.LogRequest(handlers.ArticlesHandler))

	server := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  time.Minute * 5,
		WriteTimeout: time.Minute * 5,
	}

	log.Printf("Starting server on :%d", 8080)
	server.ListenAndServe()
}
