package player

import (
	"easy-game/common/redis"
	"encoding/json"
)

// 获取用户信息
func GetPlayer(userId string, fields []string) *PlayerData {
	// 获取Redis数据
	key := redis.GetPlayerKey(userId)
	datas, _ := redis.RD.HGetAll(key).Result()
	// 赋值到用户
	playerData := newPlayerData(key, fields)
	playerData.formatPlayerData(datas)
	return playerData
}

func newPlayerData(key string, fields []string) *PlayerData {
	return &PlayerData{
		hashData: &redis.HashData{
			CacheKey:  key,
			Fields:    fields,
			SaveValus: make(map[string]interface{}),
		},
		Heros: make(map[int]*Hero, 0),
		Attr:  &Attr{},
		Bag:   make(map[int]*Item, 0),
	}
}

// 保存用户信息
func (p *PlayerData) SaveValus() {
	p.formatSave()
	p.hashData.Save()
}

// 赋值对象
func (p *PlayerData) formatPlayerData(datas map[string]string) {
	for _, field := range p.hashData.Fields {
		if value, ok := datas[field]; ok {
			p.formatValue(field, value)
		}
	}
}

// 赋值参数到对象
func (p *PlayerData) formatValue(field, value string) {
	switch field {
	case ID:
		p.Id = value
	case USER_ID:
		p.UserId = value
	case NICK_NAME:
		p.NickName = value
	case HEROS:
		json.Unmarshal([]byte(value), &p.Heros)
	case ATTR:
		json.Unmarshal([]byte(value), p.Attr)
	case BAG:
		json.Unmarshal([]byte(value), &p.Bag)
	}
}

// 赋值对象参数到保存对象
func (p *PlayerData) formatSave() {
	for _, field := range p.hashData.Fields {
		switch field {
		case ID:
			p.hashData.SaveValus[ID] = p.Id
		case USER_ID:
			p.hashData.SaveValus[USER_ID] = p.UserId
		case NICK_NAME:
			p.hashData.SaveValus[NICK_NAME] = p.NickName
		case HEROS:
			s, _ := json.Marshal(p.Heros)
			p.hashData.SaveValus[HEROS] = s
		case ATTR:
			s, _ := json.Marshal(p.Attr)
			p.hashData.SaveValus[ATTR] = s
		case BAG:
			s, _ := json.Marshal(p.Bag)
			p.hashData.SaveValus[BAG] = s
		}
	}
}
