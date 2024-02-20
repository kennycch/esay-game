package config

import (
	"github.com/go-ini/ini"
)

// 配置赋值
func MapConfig(env *ini.File) {
	env.Section("app").MapTo(App)
	env.Section("log").MapTo(Log)
	env.Section("kafka").MapTo(Kafka)
	env.Section("task").MapTo(Task)
	env.Section("http").MapTo(Http)
	env.Section("redis").MapTo(Redis)
}
