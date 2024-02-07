package client

import (
	"easy-game/config"
	"easy-game/pb"
	"fmt"

	"github.com/serialx/hashring"
)

// 初始化任务分配器
func TaskInit(listen func(channel chan *Task)) {
	// 创建玩家任务分配器
	playerKeys := make([]string, 0)
	for i := 0; i < int(config.Task.Player); i++ {
		channel := make(chan *Task)
		key := fmt.Sprintf("%d", i)
		// 注册到分配器
		playerChannels[key] = channel
		// 添加哈希环元素
		playerKeys = append(playerKeys, key)
		// 监听管道
		listen(channel)
	}
	// 注册玩家哈希环
	playerHashring = hashring.New(playerKeys)

	// 创建公共任务分配器
	publicKeys := make([]string, 0)
	for i := 0; i < int(config.Task.Public); i++ {
		channel := make(chan *Task)
		key := fmt.Sprintf("%d", i)
		// 注册到分配器
		publicChannels[key] = channel
		// 添加哈希环元素
		publicKeys = append(publicKeys, key)
		// 监听管道
		listen(channel)
	}
	// 注册公共哈希环
	publicHashring = hashring.New(publicKeys)
}

// 获取任务管道
func GetTaskChannle(channelTask ChannelTask) (channel chan *Task) {
	if channelTask.ChannelType == PlayerType {
		key, _ := playerHashring.GetNode(channelTask.Target)
		channel = playerChannels[key]
	} else {
		key, _ := publicHashring.GetNode(channelTask.Target)
		channel = publicChannels[key]
	}
	return
}

// 注册分配器
func RegisterTaskChannle(cmd pb.CmdId, channelType ChannleType, target string) {
	channelTask := ChannelTask{
		ChannelType: channelType,
		Target:      target,
	}
	channelTaskMap[cmd] = channelTask
}

// 注册客户端请求黑名单
func RegisterBlackList(cmd pb.CmdId) {
	connBlack[cmd] = struct{}{}
}
