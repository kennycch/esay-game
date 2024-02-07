package player

import "easy-game/common/redis"

const (
	ID        = "id"
	USER_ID   = "userId"
	NICK_NAME = "nickName"
	EXP       = "exp"
	Level     = "level"
	HEROS     = "heros"
	ATTR      = "attr"
	BAG       = "bag"
)

type PlayerData struct {
	hashData *redis.HashData
	Id       string
	UserId   string
	NickName string
	Exp      int
	Level    int
	Heros    map[int]*Hero
	Attr     *Attr
	Bag      map[int]*Item
}

type Hero struct {
	HeroId   int    `json:"heroId"`
	HeroName string `json:"heroName"`
	Level    int    `json:"level"`
}

type Attr struct {
	Atk int `json:"atk"`
	Def int `json:"def"`
	Hp  int `json:"hp"`
}

type Item struct {
	ItemId   int   `json:"itemId"`
	Num      int   `json:"num"`
	ItemType uint8 `json:"itemType"`
}
