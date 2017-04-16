package router

import (
	"net/http"

	"pwcong.me/url-shortener/controller"
	"pwcong.me/url-shortener/mux"
	"pwcong.me/url-shortener/utils/httpstatus"
)

type Route func(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request)

type Router struct{}

func (rt Router) Routes(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

	indexController := controller.IndexController{}

	switch r.URL.Path {

	case "/":
		Get(mux, w, r, indexController.GetIndex)
	default:
		httpstatus.StatusBadRequest(w)
	}

}
