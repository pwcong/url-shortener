package controller

import (
	"net/http"

	"pwcong.me/url-shortener/mux"
)

type UrlController struct{}

var Url UrlController

func (c UrlController) GenerateShortUrl(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

}
