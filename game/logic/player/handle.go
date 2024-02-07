package player

import (
	"easy-game/game/client"
	"easy-game/pb"

	"github.com/kennycch/gotools/general"
)

// 心跳
func HeartBeat(task *client.Task) {
	msg := &pb.Msg{
		Cmd:  pb.CmdId_HeartBeat,
		Time: general.NowUnix(),
	}
	client.PushByPlayer(task.Player.PlayerId, msg)
}
