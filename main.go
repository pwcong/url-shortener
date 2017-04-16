package main

import (
	"net/http"

	"github.com/go-redis/redis"

	"github.com/rs/cors"
	Init "pwcong.me/url-shortener/init"
	"pwcong.me/url-shortener/mux"
	"pwcong.me/url-shortener/router"
)

func main() {

	myMux := mux.NewDBMux()

	myMux.OpenDBConnection(Init.Config.DB.MySQL.User, Init.Config.DB.MySQL.Password, Init.Config.DB.MySQL.Address, Init.Config.DB.MySQL.DBName)
	defer myMux.CloseDBConnection()

	myMux.OpenRedisClient(&redis.Options{
		Addr: Init.Config.DB.Redis.Address,
	})
	defer myMux.CloseRedisClient()

	myMux.RegisterRouter(router.Router{})

	mux := http.NewServeMux()
	mux.Handle("/", myMux)

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(Init.Config.Addr, handler)

}
