package client

import (
	"easy-game/pb"

	"github.com/gorilla/websocket"
	"github.com/kennycch/gotools/general"
	"google.golang.org/protobuf/proto"
)

// 发送信息给玩家
func PushByPlayer(playerId string, msg *pb.Msg) {
	msg.Time = general.NowMilli()
	if player := getPlayer(playerId); player != nil {
		player.Lock.Lock()
		defer player.Lock.Unlock()
		if b, err := proto.Marshal(msg); err == nil {
			player.Conn.WriteMessage(websocket.BinaryMessage, b)
		}
	}
}

// 推送全服
func PushAll(msg *pb.Msg) {
	msg.Time = general.NowMilli()
	if b, err := proto.Marshal(msg); err == nil {
		WsConnTree.Lock.Lock()
		defer WsConnTree.Lock.Unlock()
		for _, player := range WsConnTree.Conns {
			player.Lock.Lock()
			defer player.Lock.Unlock()
			player.Conn.WriteMessage(websocket.BinaryMessage, b)
		}
	}
}

// 获取玩家客户端
func getPlayer(playerId string) *PlayerConn {
	WsConnTree.Lock.Lock()
	defer WsConnTree.Lock.Unlock()
	if player, ok := WsConnTree.Conns[playerId]; ok {
		return player
	}
	return nil
}
