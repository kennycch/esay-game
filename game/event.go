package game

import (
	"easy-game/game/client"

	"github.com/kennycch/gotools/worker"
)

// 登录事件
func connectEvent() {
	worker.AddTask(func() {
		for {
			playerId := <-client.ConnectChan
			for _, event := range connectEvents {
				event(playerId)
			}
		}
	})
}

// 登出事件
func disconnectEvent() {
	worker.AddTask(func() {
		for {
			playerId := <-client.DisconnectChan
			for _, event := range disconnectEvents {
				event(playerId)
			}
		}
	})
}
