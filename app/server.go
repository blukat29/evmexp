package app

import (
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()

	// https://github.com/gin-gonic/gin/issues/75
	r.Use(static.Serve("/", static.LocalFile("./front/dist", true)))

	api := r.Group("/api")
	{
		api.POST("/code/upload", ApiCodeUpload)
		api.GET("/deco/:id", ApiDeco)
	}

	r.Run(":8000")
}

func MatchErrorCode(err error) int {
	switch err.(type) {
	case *InputError:
		return http.StatusBadRequest
	case *NotFoundError:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func ApiCodeUpload(c *gin.Context) {
	req := &CodeUploadRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if res, err := CodeUpload(req); err != nil {
		log.Print(err)
		c.JSON(MatchErrorCode(err), &Response{Error: err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func ApiDeco(c *gin.Context) {
	req := &DecoRequest{}
	if err := c.BindUri(req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if res, err := Deco(req); err != nil {
		log.Print(err)
		c.JSON(MatchErrorCode(err), &Response{Error: err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
