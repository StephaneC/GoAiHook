package main


import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Health", "GET", "/healthcheck", HealthcheckHandler},
	Route{"ApiAiHook", "POST", "/apiaihook", ApiAiHookHandler},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		log.Printf("routing : " + route.Pattern)
		router.
		Methods(route.Method).Path(route.Pattern).
			Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}
