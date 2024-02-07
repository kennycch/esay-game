package game

import (
	"easy-game/game/client"
	"easy-game/game/logic/bag"
	"easy-game/game/logic/player"
	"easy-game/pb"
)

type Task struct {
	Cmd           pb.CmdId                // 协议ID
	ChannelType   client.ChannleType      // 任务分配策略
	Target        string                  // 任务分配标的（用户策略无须配置）
	Handle        func(task *client.Task) // 处理方法
	ConnBlackList bool                    // 是否客户端请求黑名单（黑名单中协议仅能通过内部调用，客户端调用无效）
}

var (
	// 任务分配器、映射方法注册器
	taskMap = []Task{
		// 心跳
		{
			Cmd:         pb.CmdId_HeartBeat,
			ChannelType: client.PlayerType,
			Handle:      player.HeartBeat,
		},
		// 背包更变
		{
			Cmd:           pb.CmdId_BagChange,
			ChannelType:   client.PlayerType,
			ConnBlackList: true,
			Handle:        bag.BagChange,
		},
	}
)
