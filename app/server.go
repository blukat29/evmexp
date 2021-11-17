package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()
	r.POST("/api/deco", func(c *gin.Context) {
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
	r.Run()
}
