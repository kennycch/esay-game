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
	bagChange := &pb.BagChange{}
	if err := proto.Unmarshal(task.Msg.Body, bagChange); err != nil {
		return
	}
	// 加载玩家信息
	playerData := player.GetPlayer(task.Player.PlayerId, BagChangeFields)
	// 开始处理背包
	for _, item := range bagChange.Changes {
		if _, ok := playerData.Bag[item.ItemId]; ok {
			playerData.Bag[item.ItemId].Num += item.Num
		} else {
			playerData.Bag[item.ItemId] = &pb.Item{
				ItemId: item.ItemId,
				Num:    item.Num,
				ItemType: item.ItemType,
			}
		}
	}
	// 保存变更
	playerData.SaveValus()
	// 通知客户端变更
	playerData.PushBagChange(bagChange.Changes)
}
