package main

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	Init "github.com/pwcong/url-shortener/init"

	"github.com/pwcong/url-shortener/model"
	"github.com/pwcong/url-shortener/mux"
	"github.com/pwcong/url-shortener/router"
	"github.com/rs/cors"
)

func main() {

	myMux := mux.NewDBMux()

	myMux.OpenDBConnection(Init.Config.DB.MySQL, func(db *gorm.DB) {
		db.AutoMigrate(&model.Url{})
	})
	defer myMux.CloseDBConnection()

	myMux.OpenRedisClient(&redis.Options{
		Addr: Init.Config.DB.Redis.Address,
	})
	defer myMux.CloseRedisClient()

	myMux.RegisterRouter(router.Router{})

	mux := http.NewServeMux()
	mux.Handle("/", myMux)

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(Init.Config.Host+":"+Init.Config.Port, handler)

}
