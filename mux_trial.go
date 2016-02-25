package main

import (
	"fmt"

	"github.com/garrypolley/mux_trial/handlers"
	"github.com/garrypolley/mux_trial/logging"
	"github.com/garrypolley/mux_trial/middleware"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

const (
	port    = 8080
	timeout = 5
)

var log *logrus.Logger

func init() {
	log = logging.Logger
}
func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", handlers.HomeIndex)
	r.HandleFunc("/products", handlers.ProductIndex)
	r.HandleFunc("/products/{id:[0-9]+}", handlers.ProductById)

	server := negroni.New()

	server.UseHandler(r)
	server.Use(middleware.NewLoggingMiddleware())

	server.Run(fmt.Sprintf(":%d", port))
}
