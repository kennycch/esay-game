package test

import (
	"easy-game/config"
	"easy-game/game"
	"easy-game/tools/lifecycle"
	"easy-game/tools/log"
	"easy-game/tools/net"
	"easy-game/tools/redis"
)

// 服务注册器
func register() {
	// 加载配置
	lifecycle.AddLifecycle(config.NewConfig())
	// 初始化日志
	lifecycle.AddLifecycle(log.NewLogRegister())
	// 初始化游戏对象
	lifecycle.AddLifecycle(game.NewGame())
	// 开启服务
	lifecycle.AddLifecycle(net.NewNet())
	// Redis服务
	lifecycle.AddLifecycle(redis.NewRedis())
}
