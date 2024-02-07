package redis

import (
	"easy-game/config"
	"fmt"

	"github.com/go-redis/redis"
)

func RedisInit() {
	con := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,                            // 连接地址（含端口）
		Password: config.Redis.PassWord,                        // 密码
		DB:       config.Redis.DB,                              // 选择的库
		PoolSize: config.Task.Player + config.Task.Public + 10, // 连接池数量为任务管道总数量+10
	})
	_, err := con.Ping().Result()
	if err != nil {
		panic(err)
	}
	RD = con
}

func GetPlayerKey(userId string) string {
	return fmt.Sprintf("%s_%s", Player, userId)
}

func (h *HashData) Save() {
	RD.HMSet(h.CacheKey, h.SaveValus).Result()
}
