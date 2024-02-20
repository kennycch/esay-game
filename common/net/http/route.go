package http

import (
	"github.com/gin-gonic/gin"
)

func (h *HttpService) Route(route *gin.Engine) {
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"code": 200})
	})
}
