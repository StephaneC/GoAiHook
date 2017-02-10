package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

/**
 * JSON PING response
 */
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Ok")
}

func ApiAiHookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	json.NewEncoder(w).Encode(id)
}
