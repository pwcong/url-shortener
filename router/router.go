package router

import (
	"net/http"

	"pwcong.me/url-shortener/controller"
	"pwcong.me/url-shortener/mux"
)

type Route func(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request)

type Router struct{}

func (rt Router) Routes(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

	Get("/$", mux, w, r, controller.GetIndex)

}
