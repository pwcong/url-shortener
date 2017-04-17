package mux

import (
	"net/http"

	"log"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"pwcong.me/panorama-tour-sys/utils/httpstatus"
	Init "pwcong.me/url-shortener/init"
	"pwcong.me/url-shortener/utils/logger"
)

type Router interface {
	Routes(*ServeMux, http.ResponseWriter, *http.Request)
}

type ServeMux struct {
	RedisClient *redis.Client
	DB          *gorm.DB
	router      Router
}

func (mux *ServeMux) OpenRedisClient(opt *redis.Options) {

	client := redis.NewClient(opt)

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err.Error())
	}

	mux.RedisClient = client

}

func (mux *ServeMux) CloseRedisClient() {
	mux.RedisClient.Close()
}

func (mux *ServeMux) OpenDBConnection(mysqlConfig Init.MySQLConfig, onConnected func(db *gorm.DB)) {

	connectionURL := mysqlConfig.User + ":" + mysqlConfig.Password + "@" + "tcp(" + mysqlConfig.Address + ")/" + mysqlConfig.DBName + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", connectionURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	onConnected(db)

	mux.DB = db

}

func (mux *ServeMux) CloseDBConnection() {
	mux.DB.Close()
}

func (mux *ServeMux) RegisterRouter(router Router) {
	mux.router = router
}

func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	logger.Log2Access(r.RemoteAddr, r.Method, r.URL.Path)

	if mux.router != nil {
		mux.router.Routes(mux, w, r)
	} else {
		httpstatus.StatusNotFound(w)
	}

}

func NewDBMux() *ServeMux { return new(ServeMux) }
