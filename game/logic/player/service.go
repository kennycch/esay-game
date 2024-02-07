package player

import (
	"easy-game/game/client"
	"easy-game/pb"
)

// 推送客户端物品更改
func (h *HPlayer) PushBagChange(items []*pb.Item) {
	bagChange := &pb.BagChange{
		PlayerId: h.UserId,
		Changes:  items,
	}
	client.PushByPlayer(h.UserId, pb.CmdId_CMD_BagChange, bagChange)
}

// 玩家基本信息
func PlayerInfo(playerId string) {
	hPlayer := GetPlayer(playerId, AllFields)
	playerInfo := hPlayer.toPb()
	client.PushByPlayer(playerId, pb.CmdId_CMD_PlayerInfo, playerInfo)
}
