package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/garrypolley/mux_trial/logging"
	"github.com/garrypolley/mux_trial/response"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

var log *logrus.Logger

func init() {
	log = logging.Logger
}

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	response.Render.JSON(w, 200, map[string]string{"Home": "json"})
}

func ProductIndex(w http.ResponseWriter, r *http.Request) {
	response.Render.JSON(w, 200, map[string]string{"Products": "json"})
}

func ProductById(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	id, ok := pathParams["id"]

	realId, err := strconv.Atoi(id)

	if ok && realId <= 10000 && err == nil {
		response.Render.JSON(w, 200, map[string]string{"Products": "json", "ProductId": id})
	} else {
		log.Errorf("Failed to get product with id %d", realId)
		http.Error(w, fmt.Sprintf("Failed to get product with id %d", realId), 404)
	}
}
