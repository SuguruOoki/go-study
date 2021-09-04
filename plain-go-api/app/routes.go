package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// routeに対する構造を定義
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MemoIndex",
		"GET",
		"/memos",
		MemoIndex,
	},
	Route{
		"MemoShow",
		"GET",
		"/memos/{memoId}",
		MemoShow,
	},
}
