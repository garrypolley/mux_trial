package main

import (
	"log"
	"net/http"
	"time"

	"github.com/garrypolley/mux_trial"

	"github.com/gorilla/mux"
)

func LogRequest(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler(w, r)
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", LogRequest(HomeHandler))
	r.HandleFunc("/products", LogRequest(ProductsHandler))
	r.HandleFunc("/products/{id:[0-9]+}", LogRequest(ProductsIdHandler))
	r.HandleFunc("/articles", LogRequest(ArticlesHandler))

	server := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  time.Minute * 5,
		WriteTimeout: time.Minute * 5,
	}

	log.Printf("Starting server on :%d", 8080)
	server.ListenAndServe()
}
