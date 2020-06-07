package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"

	"url-shortener/config"
	"url-shortener/controller"
	"url-shortener/db"
	"url-shortener/middleware"
	"url-shortener/model"
	"url-shortener/router"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func initMiddlewares(e *echo.Echo, conf *config.Config) {
	middleware.Init(e, conf)
}

func initRoutes(e *echo.Echo, conf *config.Config, db *gorm.DB, client *redis.Client) {

	router.Init(e, conf, db, client)

}

func initDB(db *gorm.DB) {
	db.AutoMigrate(&model.Url{})
}

func main() {

	// 初始化配置
	conf, err := config.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	mySQLConfig, ok := conf.Databases["mysql"]
	if !ok {
		log.Fatal("Can not load configuration of MySQL")
	}

	orm := db.ORM{DB: nil, Name: "mysql"}

	orm.Open(
		mySQLConfig.Username,
		mySQLConfig.Password,
		mySQLConfig.Host+":"+strconv.Itoa(mySQLConfig.Port),
		mySQLConfig.DBName)

	defer orm.Close()

	redisConfig, ok := conf.Databases["redis"]
	if !ok {
		log.Fatal("Can not load configuration of Redis")
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisConfig.Host + ":" + strconv.Itoa(redisConfig.Port),
	})
	_, err = client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()

	initDB(orm.DB)

	e := echo.New()

	// 全局错误处理
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.JSON(http.StatusOK, controller.BaseResponseJSON{
			Success: false,
			Code:    controller.STATUS_ERROR,
			Message: err.Error(),
		})
	}

	// 初始化中间件
	initMiddlewares(e, &conf)
	// 初始化路由
	initRoutes(e, &conf, orm.DB, client)

	// 运行服务
	if conf.Server.Port == 80 {
		e.Logger.Fatal(e.Start(conf.Server.Host))
	} else {
		e.Logger.Fatal(e.Start(conf.Server.Host + ":" + strconv.Itoa(conf.Server.Port)))
	}

}

func init() {

}
