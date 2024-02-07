package net

import "github.com/gin-gonic/gin"

func websocketRoute(rule string, handle func(*gin.Context)) {
	route.GET(rule, handle)
}