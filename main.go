package main

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	Init "pwcong.me/url-shortener/init"

	"github.com/rs/cors"
	"pwcong.me/url-shortener/model"
	"pwcong.me/url-shortener/mux"
	"pwcong.me/url-shortener/router"
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
