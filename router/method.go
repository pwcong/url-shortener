package router

import (
	"net/http"

	"pwcong.me/url-shortener/mux"
	"pwcong.me/url-shortener/utils/httpstatus"
)

// Request Method Name
const (
	GET  = "GET"
	POST = "POST"
)

// Get can register a get routes
func Get(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request, route Route) {
	handleByMethodAndPattern(GET, mux, w, r, route)
}

// Post can register a post routes
func Post(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request, route Route) {
	handleByMethodAndPattern(POST, mux, w, r, route)
}

func handleByMethodAndPattern(method string, mux *mux.ServeMux, w http.ResponseWriter, r *http.Request, route Route) {

	if r.Method == method {
		route(mux, w, r)
	} else {
		httpstatus.StatusMethodNotAllowed(w)
	}
}
