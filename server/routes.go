package server

import "net/http"

type RouteConfig struct {
	Endpoint string
	Method   string
	Handler  func(w http.ResponseWriter, r *http.Request)
}

type Routes []RouteConfig
