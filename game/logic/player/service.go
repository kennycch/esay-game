package player

import (
	"easy-game/game/client"
	"easy-game/pb"

	"google.golang.org/protobuf/proto"
)

// 推送客户端物品更改
func (h *HPlayer) PushBagChange(items []*pb.Item) {
	bagChange := &pb.BagChange{
		PlayerId: h.UserId,
		Changes:  items,
	}
	b, _ := proto.Marshal(bagChange)
	msg := &pb.Msg{
		Cmd:  pb.CmdId_CMD_BagChange,
		Body: b,
	}
	client.PushByPlayer(h.UserId, msg)
}

// 玩家基本信息
func PlayerInfo(playerId string) {
	hPlayer := GetPlayer(playerId, AllFields)
	playerInfo := hPlayer.toPb()
	b, _ := proto.Marshal(playerInfo)
	msg := &pb.Msg{
		Cmd:  pb.CmdId_CMD_PlayerInfo,
		Body: b,
	}
	client.PushByPlayer(playerId, msg)
}
