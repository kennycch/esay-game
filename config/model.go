package config

import "flag"

var (
	// 配置路径
	filePath = flag.String("c", "./env.ini", "config file")
	// 总配置
	App = &struct {
		AppName string // 进程名
	}{}
	// 日志配置
	Log = &struct {
		LogPath  string // 日志存放路径
		LogDay   int    // 日志持久化天数
		LogLevel int    // 日志等级
	}{}
	// Http配置
	Http = &struct {
		Port    int    // 服务监听端口
		WSRoute string // websocket路由
	}{}
	// 任务分配器配置
	Task = &struct {
		Player int // 玩家任务管道数
		Public int // 公共任务管道数
	}{}
	// Redis配置
	Redis = &struct {
		Addr     string // Redis地址
		PassWord string // 密码
		DB       int    // 选择库
	}{}
)
