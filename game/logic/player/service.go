package player

import (
	"easy-game/game/client"
	"easy-game/pb"

	"google.golang.org/protobuf/proto"
)

// 推送客户端物品更改
func (p *PlayerData) PushBagChange(items []*pb.Item) {
	playerBagChange := &pb.PlayerBagChange{
		PlayerId: p.UserId,
		Changes:  items,
	}
	b, _ := proto.Marshal(playerBagChange)
	msg := &pb.Msg{
		Cmd:  pb.CmdId_BagChange,
		Body: b,
	}
	client.PushByPlayer(p.UserId, msg)
}
