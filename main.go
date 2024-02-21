package main

import (
	"easy-game/config"
	"easy-game/game"
	"easy-game/tools/cron"
	"easy-game/tools/lifecycle"
	"easy-game/tools/log"
	"easy-game/tools/net"
	"easy-game/tools/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 服务注册器
	register()
	// 服务开启事件
	lifecycle.Start()
	// 服务结束事件
	defer lifecycle.Stop()
	// 信号监听
	loop()
}

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
	// 定时任务
	lifecycle.AddLifecycle(cron.NewCron())
}

// 信号监听
func loop() {
	log.Info("server started")
	signals := make(chan os.Signal, 1)
	// kill -9 无法被捕获
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
label:
	for {
		s := <-signals
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			break label
		case syscall.SIGHUP:
			continue
		}
	}
	log.Info("server close")
}
