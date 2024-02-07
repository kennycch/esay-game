package bag

import (
	"easy-game/game/client"
	"easy-game/game/logic/player"
	"easy-game/pb"

	"google.golang.org/protobuf/proto"
)

// 背包变更
func BagChange(task *client.Task) {
	// 解析请求
	playerBagChange := &pb.PlayerBagChange{}
	if err := proto.Unmarshal(task.Msg.Body, playerBagChange); err != nil {
		return
	}
	// 加载玩家信息
	playerData := player.GetPlayer(task.Player.PlayerId, BagChangeFields)
	// 开始处理背包
	for _, item := range playerBagChange.Changes {
		if _, ok := playerData.Bag[int(item.ItemId)]; ok {
			playerData.Bag[int(item.ItemId)].Num += int(item.Num)
		} else {
			playerData.Bag[int(item.ItemId)] = &player.Item{
				ItemId: int(item.ItemId),
				Num:    int(item.Num),
			}
		}
	}
	// 保存变更
	playerData.SaveValus()
	// 通知客户端变更
	playerData.PushBagChange(playerBagChange.Changes)
}
