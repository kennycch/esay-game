package logic

import (
	"easy-game/game/client"
	"easy-game/pb"

	"github.com/kennycch/gotools/worker"
)

// 任务执行器
func Listen(channel chan *client.Task) {
	worker.AddTask(func() {
		for {
			task := <-channel
			if handle, ok := handleMap[task.Msg.Cmd]; ok {
				handle(task)
			}
		}
	})
}

// 注册执行器
func RegisterHandle(cmd pb.CmdId, handle func(task *client.Task)) {
	handleMap[cmd] = handle
}
