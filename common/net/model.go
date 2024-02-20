package net

import (
	"easy-game/common/net/http"
	"easy-game/common/net/pprof"
	"easy-game/common/net/websocket"

	"github.com/gin-gonic/gin"
)

type RouteService interface {
	Route(route *gin.Engine)
}

var (
	// Gin路由对象
	route = &gin.Engine{}
	// 要注册的服务路由
	routeServices = []RouteService{
		&pprof.PprofService{},
		&http.HttpService{},
		&websocket.WebsocketService{},
	}
)
