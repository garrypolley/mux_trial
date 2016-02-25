package mux_trial

import (
	"io"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My server: "+r.URL.String())
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My server: "+r.URL.String())
}

func ProductsIdHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	id, ok := pathParams["id"]

	if ok {
		io.WriteString(w, "My server: "+r.URL.String()+" id is "+id)
	} else {
		log.Println("failed to get an id on the route, this should never happen")
	}
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My server: "+r.URL.String())
}
