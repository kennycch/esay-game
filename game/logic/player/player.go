package player

import (
	"easy-game/common/redis"
	"easy-game/pb"
	"encoding/json"
	"strconv"

	"github.com/kennycch/gotools/general"
)

// 获取用户信息
func GetPlayer(userId string, fields []string) *HPlayer {
	// 获取Redis数据
	key := redis.GetPlayerKey(userId)
	datas, _ := redis.RD.HGetAll(key).Result()
	// 赋值到用户
	playerData := newHPlayer(key, fields)
	playerData.formatHPlayer(datas)
	return playerData
}

func newHPlayer(key string, fields []string) *HPlayer {
	return &HPlayer{
		hashData: &redis.HashData{
			CacheKey:  key,
			Fields:    fields,
			SaveValus: make(map[string]interface{}),
		},
		Heros: make(map[int32]*pb.Hero, 0),
		Attr:  &pb.Attr{},
		Bag:   make(map[int32]*pb.Item, 0),
	}
}

// 保存用户信息
func (h *HPlayer) SaveValus() {
	h.formatSave()
	h.hashData.Save()
}

// 赋值对象
func (h *HPlayer) formatHPlayer(datas map[string]string) {
	for _, field := range h.hashData.Fields {
		if value, ok := datas[field]; ok {
			h.formatValue(field, value)
		}
	}
}

// 赋值参数到对象
func (h *HPlayer) formatValue(field, value string) {
	switch field {
	case ID:
		h.Id = strToInt(value)
	case USER_ID:
		h.UserId = value
	case NICK_NAME:
		h.NickName = value
	case EXP:
		h.Exp = strToInt(value)
	case LEVEL:
		h.Level = strToInt(value)
	case HEROS:
		json.Unmarshal([]byte(value), &h.Heros)
	case ATTR:
		json.Unmarshal([]byte(value), h.Attr)
	case BAG:
		json.Unmarshal([]byte(value), &h.Bag)
	}
}

// 赋值对象参数到保存对象
func (h *HPlayer) formatSave() {
	for _, field := range h.hashData.Fields {
		switch field {
		case ID:
			h.hashData.SaveValus[ID] = h.Id
		case USER_ID:
			h.hashData.SaveValus[USER_ID] = h.UserId
		case NICK_NAME:
			h.hashData.SaveValus[NICK_NAME] = h.NickName
		case EXP:
			h.hashData.SaveValus[EXP] = h.Exp
		case LEVEL:
			h.hashData.SaveValus[LEVEL] = h.Level
		case HEROS:
			s, _ := json.Marshal(h.Heros)
			h.hashData.SaveValus[HEROS] = s
		case ATTR:
			s, _ := json.Marshal(h.Attr)
			h.hashData.SaveValus[ATTR] = s
		case BAG:
			s, _ := json.Marshal(h.Bag)
			h.hashData.SaveValus[BAG] = s
		}
	}
}

func (h *HPlayer) toPb() *pb.PlayerInfo {
	return &pb.PlayerInfo{
		Id:       h.Id,
		UserId:   h.UserId,
		NickName: h.NickName,
		Exp:      h.Exp,
		Level:    h.Level,
		Heros:    general.MapValues(h.Heros),
		Attr:     h.Attr,
		Bag:      general.MapValues(h.Bag),
	}
}

func strToInt(value string) int32 {
	num, _ := strconv.Atoi(value)
	return int32(num)
}
