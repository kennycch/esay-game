package redis

import "github.com/go-redis/redis"

const (
	Player = "player"
)

var (
	RD *redis.Client
)

type HashData struct {
	CacheKey       string
	Fields    []string
	SaveValus map[string]interface{}
}
