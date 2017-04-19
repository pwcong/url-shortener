package controller

import (
	"fmt"
	"net/http"

	"github.com/pwcong/url-shortener/mux"
)

type IndexController struct{}

var Index IndexController

func (c IndexController) GetIndex(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello World")

}
