package net

import (
	"easy-game/config"
	"easy-game/tools/lifecycle"
	"easy-game/tools/net/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kennycch/gotools/worker"
)

type Net struct{}

func (n *Net) Start() {
	route = gin.Default()
	// 允许跨域
	route.Use(middleware.Cors())
	// 注册路由
	for _, service := range routeServices {
		service.Route(route)
	}
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
