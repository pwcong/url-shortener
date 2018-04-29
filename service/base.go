package service

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/pwcong/url-shortener/config"
)

type BaseService struct {
	Conf   *config.Config
	DB     *gorm.DB
	Client *redis.Client
}
