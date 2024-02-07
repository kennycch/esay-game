package player

import (
	"easy-game/game/client"
	"easy-game/pb"
)

// 心跳
func HeartBeat(task *client.Task) {
	client.PushByPlayer(task.Player.PlayerId, pb.CmdId_CMD_HeartBeat, nil)
}
