package main

import (
	"net/http"

	"github.com/go-redis/redis"

	Init "pwcong.me/url-shortener/init"
	"pwcong.me/url-shortener/mux"
)

func main() {

	mux := mux.NewDBMux()

	mux.OpenDBConnection(Init.Config.DB.MySQL.User, Init.Config.DB.MySQL.Password, Init.Config.DB.MySQL.Address, Init.Config.DB.MySQL.DBName)
	defer mux.CloseDBConnection()

	mux.OpenRedisClient(&redis.Options{
		Addr: Init.Config.DB.Redis.Address,
	})
	defer mux.CloseRedisClient()

	http.ListenAndServe(":80", mux)

}
