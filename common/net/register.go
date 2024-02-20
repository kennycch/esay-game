package net

import (
	"easy-game/common/lifecycle"
	"easy-game/common/net/middleware"
	"easy-game/config"
	"easy-game/game/client"
	"fmt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/kennycch/gotools/worker"
)

type Net struct{}

func (n *Net) Start() {
	route = gin.Default()
	// 允许跨域
	route.Use(middleware.Cors())
	// 注册pprof
	pprof.Register(route, "pprof")
	// 注册Http路由
	httpRoute()
	// websocket处理
	websocketRoute(config.Http.WSRoute, client.WebsocketUpgrader)
	// 开启服务
	worker.AddTask(func() {
		route.Run(fmt.Sprintf(":%d", config.Http.Port))
	})
}

func (n *Net) Priority() uint32 {
	return lifecycle.LowPriority
}

func (n *Net) Stop() {

}

func NewNet() *Net {
	return &Net{}
}
