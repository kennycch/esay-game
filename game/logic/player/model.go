package player

import (
	"easy-game/pb"
	"easy-game/tools/redis"
)

const (
	ID        = "id"
	USER_ID   = "userId"
	NICK_NAME = "nickName"
	EXP       = "exp"
	LEVEL     = "level"
	HEROS     = "heros"
	ATTR      = "attr"
	BAG       = "bag"
)

type HPlayer struct {
	hashData *redis.HashData
	Id       int32
	UserId   string
	NickName string
	Exp      int32
	Level    int32
	Heros    map[int32]*pb.Hero
	Attr     *pb.Attr
	Bag      map[int32]*pb.Item
}

var (
	// 玩家全部字段
	AllFields = []string{ID, USER_ID, NICK_NAME, EXP, LEVEL, HEROS, ATTR, BAG}
)
