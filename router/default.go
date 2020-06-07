package router

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"url-shortener/config"
	"url-shortener/controller"
	"url-shortener/service"
)

func Init(e *echo.Echo, conf *config.Config, db *gorm.DB, client *redis.Client) {

	baseService := &service.BaseService{Conf: conf, DB: db, Client: client}
	baseController := &controller.BaseController{Conf: conf, Service: baseService}
	urlController := &controller.UrlController{Base: baseController}

	e.GET("/:id", urlController.Redirect)

	apiGroup := e.Group("/api")
	apiGroup.POST("/url/l2s", urlController.CL2S)
	apiGroup.GET("/url/s2l/:id", urlController.CS2L)
}
