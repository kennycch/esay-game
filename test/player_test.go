package test

import (
	"easy-game/game/logic/player"
	"easy-game/tools/lifecycle"
	"fmt"
	"testing"
)

/*
测试获取玩家信息方法
测试从Redis中获取不存在的哈希字段时是否会处理异常
*/
func TestGetPlayer(t *testing.T) {
	// 服务注册器
	register()
	// 服务开启事件
	lifecycle.Start()
	// 服务结束事件
	defer lifecycle.Stop()
	p := player.GetPlayer("test001", []string{
		player.LEVEL,
		"ids",
		player.BAG,
	})
	fmt.Println(p)
}
