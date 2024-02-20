package config

import (
	"easy-game/common/lifecycle"
	"flag"

	"github.com/go-ini/ini"
)

type Config struct{}

func (c *Config) Start() {
	flag.Parse()
	// 读取env配置文件
	if env, err := ini.Load(*filePath); err != nil {
		panic(err)
	} else {
		MapConfig(env)
	}
}

func (c *Config) Priority() uint32 {
	return lifecycle.HighPriority + 10000
}

func (c *Config) Stop() {

}

func NewConfig() *Config {
	return &Config{}
}
