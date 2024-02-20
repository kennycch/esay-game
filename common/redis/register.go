package redis

import (
	"easy-game/common/lifecycle"
)

type Redis struct{}

func (r *Redis) Start() {
	RedisInit()
}

func (r *Redis) Priority() uint32 {
	return lifecycle.HighPriority + 10000
}

func (r *Redis) Stop() {
	RD.Close()
}

func NewRedis() *Redis {
	return &Redis{}
}
