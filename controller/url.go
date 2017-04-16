package controller

import (
	"net/http"
	"time"

	"encoding/json"

	"fmt"

	Init "pwcong.me/url-shortener/init"
	"pwcong.me/url-shortener/model"
	"pwcong.me/url-shortener/mux"
	"pwcong.me/url-shortener/utils/httpstatus"
	"pwcong.me/url-shortener/utils/shortener"
)

type UrlJSONResponse struct {
	Err      string
	LongUrl  string
	ShortUrl string
}

var UrlJSONResponsePrefix string

type UrlController struct{}

var Url UrlController

func handleConvertLongURL2ShortURLResponse(w http.ResponseWriter, format string, longUrl string, shortUrl string) {

	if format == "json" {

		res, err := json.Marshal(UrlJSONResponse{
			LongUrl:  longUrl,
			ShortUrl: UrlJSONResponsePrefix + shortUrl,
		})

		if err != nil {
			httpstatus.StatusInternalServerError(w)
			return
		}

		w.Write(res)

	} else {

		fmt.Fprint(w, UrlJSONResponsePrefix+shortUrl)

	}
}

func (c UrlController) GetConvertLongURL2ShortURL(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

	longURL := r.URL.Query().Get("url")

	if longURL == "" {
		httpstatus.StatusBadRequest(w)
		return
	}

	format := r.URL.Query().Get("format")

	/*
	 * search from redis
	 */
	val := mux.RedisClient.Get(longURL).Val()

	if val != "" {

		handleConvertLongURL2ShortURLResponse(w, format, longURL, val)

		return
	}

	var url model.Url
	notFound := mux.DB.First(&url, "source = ?", longURL).RecordNotFound()

	if notFound {

		mux.DB.Create(&model.Url{Source: longURL, CreatedAt: time.Now()})

		notFound = mux.DB.First(&url, "source = ?", longURL).RecordNotFound()

		if notFound {
			httpstatus.StatusInternalServerError(w)
			return
		}

		shortURL := shortener.ConvertInt64ToString(url.ID)

		err := mux.RedisClient.Set(longURL, shortURL, 0).Err()
		if err != nil {
			httpstatus.StatusInternalServerError(w)
			return
		}

		handleConvertLongURL2ShortURLResponse(w, format, longURL, shortURL)

		return

	}

	shortURL := shortener.ConvertInt64ToString(url.ID)

	err := mux.RedisClient.Set(longURL, shortURL, 0).Err()
	if err != nil {
		httpstatus.StatusInternalServerError(w)
		return
	}

	handleConvertLongURL2ShortURLResponse(w, format, longURL, shortURL)

}

func (c UrlController) GetConvertShortURL2LongURL(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

	shortURL := r.URL.Path[1:]

	val := mux.RedisClient.Get(shortURL).Val()

	if val != "" {
		http.Redirect(w, r, val, 302)

		return
	}

	id, err := shortener.ConvertStringToInt64(shortURL)

	if err != nil {
		httpstatus.StatusInternalServerError(w)
		return
	}

	var url model.Url

	notFound := mux.DB.First(&url, id).RecordNotFound()

	if notFound {

		httpstatus.StatusNotFound(w)
		return

	}

	err = mux.RedisClient.Set(shortURL, url.Source, 0).Err()
	if err != nil {
		httpstatus.StatusInternalServerError(w)
		return
	}

	http.Redirect(w, r, url.Source, 302)

}

func init() {

	if Init.Config.Port == "80" {
		UrlJSONResponsePrefix = Init.Config.Server + "/"
	} else {
		UrlJSONResponsePrefix = Init.Config.Server + ":" + Init.Config.Port + "/"
	}
}
