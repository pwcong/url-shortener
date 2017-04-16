package controller

import (
	"fmt"
	"net/http"

	"pwcong.me/url-shortener/mux"
)

func GetIndex(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello")

}
