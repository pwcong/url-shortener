package router

import (
	"net/http"

	"regexp"

	"pwcong.me/url-shortener/mux"
	"pwcong.me/url-shortener/utils/httpstatus"
)

// Request Method Name
const (
	GET  = "GET"
	POST = "POST"
)

// Get can register a get routes
func Get(pattern string, mux *mux.ServeMux, w http.ResponseWriter, r *http.Request, route Route) {
	handleByMethodAndPattern(pattern, GET, mux, w, r, route)
}

// Post can register a post routes
func Post(pattern string, mux *mux.ServeMux, w http.ResponseWriter, r *http.Request, route Route) {
	handleByMethodAndPattern(pattern, POST, mux, w, r, route)
}

func handleByMethodAndPattern(pattern string, method string, mux *mux.ServeMux, w http.ResponseWriter, r *http.Request, route Route) {

	matched, err := regexp.MatchString(pattern, r.URL.Path)

	if matched && err == nil {
		if r.Method == method {
			route(mux, w, r)
		} else {
			httpstatus.StatusMethodNotAllowed(w)
		}
	} else {

		if err != nil {
			httpstatus.StatusInternalServerError(w)
		} else {
			httpstatus.StatusBadRequest(w)
		}
	}
}
