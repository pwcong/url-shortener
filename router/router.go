package router

import (
	"net/http"

	"regexp"

	"log"

	"pwcong.me/url-shortener/controller"
	"pwcong.me/url-shortener/mux"
	"pwcong.me/url-shortener/utils/httpstatus"
)

const (
	ROUTES_GET_INDEX_INDEX = "^/$"    // index
	ROUTES_GET_URL_CL2S    = "^/api$" // convert long url to short url
	ROUTES_GET_URL_CS2L    = "^/[0-9A-Z]+$"
)

type Route func(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request)

type Router struct{}

func MatchPath(pattern string, path string) bool {

	matched, err := regexp.MatchString(pattern, path)

	if err != nil {
		log.Fatal(err.Error())
	}

	return matched

}

func (rt Router) Routes(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	if MatchPath(ROUTES_GET_INDEX_INDEX, path) {

		Get(mux, w, r, controller.Index.GetIndex)

	} else if MatchPath(ROUTES_GET_URL_CL2S, path) {

		Get(mux, w, r, controller.Url.GetConvertLongURL2ShortURL)

	} else if MatchPath(ROUTES_GET_URL_CS2L, path) {

		Get(mux, w, r, controller.Url.GetConvertShortURL2LongURL)

	} else {
		httpstatus.StatusNotFound(w)
	}

}
