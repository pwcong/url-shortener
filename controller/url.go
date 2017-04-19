package controller

import (
	"net/http"
	"time"

	"encoding/json"

	"fmt"

	"regexp"

	Init "github.com/pwcong/url-shortener/init"
	"github.com/pwcong/url-shortener/model"
	"github.com/pwcong/url-shortener/mux"
	"github.com/pwcong/url-shortener/utils/httpstatus"
	"github.com/pwcong/url-shortener/utils/logger"
	"github.com/pwcong/url-shortener/utils/shortener"
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
			logger.Log2Error("UrlController", "handleConvertLongURL2ShortURLResponse", err.Error())
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

	logger.Log2Server("UrlController", "GetConvertLongURL2ShortURL", longURL)

	if longURL == "" {
		httpstatus.StatusBadRequest(w)
		return
	}

	m1, err := regexp.MatchString("^http://", longURL)
	m2, err := regexp.MatchString("^https://", longURL)

	if err != nil {
		logger.Log2Error("UrlController", "GetConvertLongURL2ShortURL", err.Error())
		httpstatus.StatusInternalServerError(w)
		return
	}

	if !(m1 || m2) {
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
			logger.Log2Error("UrlController", "GetConvertLongURL2ShortURL", err.Error())
			httpstatus.StatusInternalServerError(w)
			return
		}

		handleConvertLongURL2ShortURLResponse(w, format, longURL, shortURL)

		return

	}

	shortURL := shortener.ConvertInt64ToString(url.ID)

	err = mux.RedisClient.Set(longURL, shortURL, 0).Err()
	if err != nil {
		httpstatus.StatusInternalServerError(w)
		return
	}

	handleConvertLongURL2ShortURLResponse(w, format, longURL, shortURL)

}

func (c UrlController) GetConvertShortURL2LongURL(mux *mux.ServeMux, w http.ResponseWriter, r *http.Request) {

	shortURL := r.URL.Path[1:]

	logger.Log2Server("UrlController", "GetConvertShortURL2LongURL", shortURL)

	val := mux.RedisClient.Get(shortURL).Val()

	if val != "" {
		http.Redirect(w, r, val, 302)

		return
	}

	id, err := shortener.ConvertStringToInt64(shortURL)

	if err != nil {
		logger.Log2Error("UrlController", "GetConvertShortURL2LongURL", err.Error())
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
		logger.Log2Error("UrlController", "GetConvertShortURL2LongURL", err.Error())
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
