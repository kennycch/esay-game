package net

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func httpRoute() {
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 111})
	})
}
