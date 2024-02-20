package config

import "flag"

var (
	// 配置路径
	filePath = flag.String("c", "./env.ini", "config file")

	// 总配置
	App = &struct {
		AppName    string // 进程名
		Debug      bool   // 是否开启debug模式
		RegionName string // 区服名称
	}{}

	// Http配置
	Http = &struct {
		Port    int    // 服务监听端口
		WSRoute string // websocket路由
	}{}

	Kafka = &struct {
		Host           string
		ServerTopic    string // 服务日志topic
		StatisticTopic string // 统计信息日志topic
	}{}

	// 日志配置
	Log = &struct {
		LogPath      string // 输出类型
		RemoteEnable bool   // 是否开启远程日志
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
