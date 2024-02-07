package logic

import (
	"easy-game/game/client"
	"easy-game/pb"
)

var (
	// 处理方法映射
	handleMap = map[pb.CmdId]func(task *client.Task){}
)
