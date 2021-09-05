package main

import "net/http"

// routeに対する構造を定義
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// TODO: routerとrouteのファイル分離を行う
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
