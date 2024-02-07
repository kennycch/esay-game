package bag

import (
	"easy-game/game/client"
	"easy-game/pb"

	"github.com/kennycch/gotools/worker"
	"google.golang.org/protobuf/proto"
)

/*
异步背包物品变更
当执行逻辑的任务分发器非玩家策略时，使用此方法更新背包
*/
func AsyncChangeItems(playerId string, items []*pb.Item) {
	// 获取任务管道
	channelTask := client.ChannelTask{
		ChannelType: client.PlayerType,
		Target:      playerId,
	}
	channel := client.GetTaskChannle(channelTask)
	// 准备请求信息
	playerBagChange := &pb.PlayerBagChange{
		PlayerId: playerId,
		Changes:  items,
	}
	b, _ := proto.Marshal(playerBagChange)
	msg := &pb.Msg{
		Cmd:  pb.CmdId_BagChange,
		Body: b,
	}
	task := &client.Task{
		Player: &client.PlayerConn{
			PlayerId: playerId,
		},
		Msg: msg,
	}
	// 异步压入管道
	worker.AddTask(func() {
		channel <- task
	})
}
