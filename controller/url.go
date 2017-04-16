package controller

import (
	"net/http"

	"pwcong.me/url-shortener/mux"
)

type UrlController struct{}

var Url UrlController

func (c UrlController) ConvertLongUrl2ShortUrl(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

}

func (c UrlController) ConvertShortUrl2LongUrl(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

}
