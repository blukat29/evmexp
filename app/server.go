package app

import (
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
	}

	r.Run(":8000")
}
