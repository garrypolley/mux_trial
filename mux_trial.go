package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	hd "github.com/garrypolley/mux_trial/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	port    = 8080
	timeout = 5
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", hd.HomeHandler)
	r.HandleFunc("/products", hd.ProductsHandler)
	r.HandleFunc("/products/{id:[0-9]+}", hd.ProductsIdHandler)
	r.HandleFunc("/articles", hd.ArticlesHandler)

	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, r)
	compressedAndLogged := handlers.CompressHandler(loggedRouter)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      compressedAndLogged,
		ReadTimeout:  time.Minute * timeout,
		WriteTimeout: time.Minute * timeout,
	}

	log.Printf("Starting server on :%d with a timeout of %d minutes", port, timeout)
	server.ListenAndServe()
}
