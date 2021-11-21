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
		api.POST("/deco", func(c *gin.Context) {
			req := &ContractRequest{}
			if err := c.Bind(req); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			}
			if res, err := Decompile(req); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			} else {
				c.JSON(http.StatusOK, res)
			}
		})

		api.POST("/code/upload", ApiCodeUpload)
	}

	r.Run(":8000")
}

func MatchErrorCode(err error) int {
	switch err.(type) {
	case *InputError:
		return http.StatusBadRequest
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
