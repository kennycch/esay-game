package game

import (
	"easy-game/game/client"
	"easy-game/game/logic"
	"easy-game/lifecycle"
)

type Game struct{}

func (g *Game) Start() {
	// 注册分配器、执行器树
	MapToTree()
	// 初始化任务分配器
	client.TaskInit(logic.Listen)
	// 监听连接事件
	connectEvent()
	// 监听断开事件
	disconnectEvent()
}

func (g *Game) Priority() uint32 {
	return lifecycle.HighPriority
}

func (g *Game) Stop() {

}

func NewGame() *Game {
	return &Game{}
}
