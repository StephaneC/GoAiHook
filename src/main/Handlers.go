package main

import (
	"net/http"
	"encoding/json"
	"log"
	"time"
	"strconv"
)

/**
 * JSON PING response
 */
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Ok")
}

func ApiAiHookHandler(w http.ResponseWriter, r *http.Request) {
	//first check that we have an input
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var error = "I don't understand what you said!"
	//get request
	var request ApiAiRequest
	var response ApiAiResponse
	request, err := decodeRequest(r)
	if err != nil {
		log.Println("Couldn't parse body for ApiAiHookHandler")
		response = ApiAiResponse{DisplayText: error}
		return;
	} else {
		log.Println("request.result.action "+response.DisplayText)
		switch (request.Result.Action){
		case "action.time":
			response = ApiAiResponse{DisplayText:strconv.Itoa(time.Now().Hour()), Speech:strconv.Itoa(time.Now().Hour()), Source: "SCA"}
			break;
		/** add your actions here */
		default:
			response = ApiAiResponse{DisplayText: error}
			break;

		}
	}
	log.Println("response.displaytest "+response.DisplayText)
	json.NewEncoder(w).Encode(response)
}

func decodeRequest(r *http.Request)(ApiAiRequest, error) {
	var t ApiAiRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		log.Println("Error parsing apiairequest", err)
		panic(err)
		return t, err
	}
	defer r.Body.Close()
	return t, err
}
