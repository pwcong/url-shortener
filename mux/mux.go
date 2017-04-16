package mux

import (
	"net/http"

	"log"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"pwcong.me/panorama-tour-sys/utils/httpstatus"
)

type Router interface {
	Routes(*ServeMux, http.ResponseWriter, *http.Request)
}

type ServeMux struct {
	redisClient *redis.Client
	db          *gorm.DB
	router      Router
}

func (mux *ServeMux) OpenRedisClient(opt *redis.Options) {

	client := redis.NewClient(opt)

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err.Error())
	}

	mux.redisClient = client

}

func (mux *ServeMux) CloseRedisClient() {
	mux.redisClient.Close()
}

func (mux *ServeMux) OpenDBConnection(user string, password string, address string, dbname string) {

	connectionURL := user + ":" + password + "@" + "tcp(" + address + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", connectionURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	mux.db = db

}

func (mux *ServeMux) CloseDBConnection() {
	mux.db.Close()
}

func (mux *ServeMux) RegisterRouter(router Router) {
	mux.router = router
}

func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if mux.router != nil {
		mux.router.Routes(mux, w, req)
	} else {
		httpstatus.StatusBadRequest(w)
	}

}

func NewDBMux() *ServeMux { return new(ServeMux) }
