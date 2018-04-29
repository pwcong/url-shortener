package service

import (
	"errors"
	"strconv"

	"github.com/pwcong/url-shortener/model"
)

type UrlService struct {
	Base *BaseService
}

func (ctx *UrlService) ConvertLong2Short(sourceUrl string) (string, error) {
	db := ctx.Base.DB
	client := ctx.Base.Client

	id := client.Get(sourceUrl).Val()
	if id != "" {
		return id, nil
	}

	var url model.Url
	if notFound := db.Where("source_url = ?", sourceUrl).First(&url).RecordNotFound(); notFound {
		url = model.Url{
			SourceUrl: sourceUrl,
		}
		if err := db.Create(&url).Error; err != nil {
			return "", err
		}
	}

	return strconv.Itoa(int(url.ID)), nil

}

func (ctx *UrlService) ConvertShort2Long(id uint) (string, error) {
	db := ctx.Base.DB
	client := ctx.Base.Client

	sourceUrl := client.Get(strconv.Itoa(int(id))).Val()
	if sourceUrl != "" {
		return sourceUrl, nil
	}

	var url model.Url
	if notFound := db.Where("id = ?", id).First(&url).RecordNotFound(); notFound {
		return "", errors.New("short url has not been registered")
	}

	return url.SourceUrl, nil
}
