package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pwcong/url-shortener/service"
)

type UrlController struct {
	Base *BaseController
}

type UrlForm struct {
	SourceUrl string `json:"source_url" form:"source_url"`
}

type UrlJSONResponse struct {
	SourceUrl string `json:"source_url"`
	ShortUrl  string `json:"short_url"`
}

func (ctx *UrlController) Redirect(c echo.Context) error {

	service := service.UrlService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	sourceUrl, err := service.ConvertShort2Long(uint(id))
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, sourceUrl)

}

func (ctx *UrlController) CL2S(c echo.Context) error {

	service := service.UrlService{Base: ctx.Base.Service}

	form := new(UrlForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	if form.SourceUrl == "" {
		return errors.New("source url can not be empty")
	}

	id, err := service.ConvertLong2Short(form.SourceUrl)
	if err != nil {
		return err
	}

	proto, ok := PROTOS[c.Request().Proto]
	if !ok {
		return errors.New("invalid http proto")
	}

	return BaseResponse(c, true, STATUS_OK, "convert long url to short url successfully", UrlJSONResponse{
		SourceUrl: form.SourceUrl,
		ShortUrl:  proto + c.Request().Host + "/" + id,
	})

}
func (ctx *UrlController) CS2L(c echo.Context) error {

	service := service.UrlService{Base: ctx.Base.Service}

	_id := c.Param("id")

	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	sourceUrl, err := service.ConvertShort2Long(uint(id))
	if err != nil {
		return err
	}

	proto, ok := PROTOS[c.Request().Proto]
	if !ok {
		return errors.New("invalid http proto")
	}

	return BaseResponse(c, true, STATUS_OK, "convert long url to short url successfully", UrlJSONResponse{
		SourceUrl: sourceUrl,
		ShortUrl:  proto + c.Request().Host + "/" + _id,
	})

}
