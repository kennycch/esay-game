package game

import (
	"easy-game/game/client"
	"easy-game/game/logic"
)

// 注册分配器、执行器树
func MapToTree() {
	for _, task := range taskMap {
		// 注册分配器
		client.RegisterTaskChannle(task.Cmd, task.ChannelType, task.Target)
		// 注册执行器
		logic.RegisterHandle(task.Cmd, task.Handle)
		// 注册客户端请求黑名单
		if task.ConnBlackList {
			client.RegisterBlackList(task.Cmd)
		}
	}
}
