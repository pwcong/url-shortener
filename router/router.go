package router

import (
	"net/http"

	"regexp"

	"pwcong.me/url-shortener/controller"
	"pwcong.me/url-shortener/mux"
	"pwcong.me/url-shortener/utils/httpstatus"
)

type Route func(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request)

type Router struct{}

func (rt Router) Routes(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	if matched, _ := regexp.MatchString("/$", path); matched {
		Get(mux, w, r, controller.Index.GetIndex)
	} else {
		httpstatus.StatusBadRequest(w)
	}

}
