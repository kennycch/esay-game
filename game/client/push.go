package client

import (
	"easy-game/pb"

	"github.com/gorilla/websocket"
	"github.com/kennycch/gotools/general"
	"google.golang.org/protobuf/proto"
)

// 发送信息给玩家
func PushByPlayer(playerId string, cmd pb.CmdId, resp proto.Message) {
	msg := &pb.Msg{
		Cmd:  cmd,
		Time: general.NowMilli(),
	}
	if resp != nil {
		if b, err := proto.Marshal(resp); err == nil {
			msg.Body = b
		}
	}
	m, _ := proto.Marshal(msg)
	WsConnTree.Lock.RLock()
	defer WsConnTree.Lock.RUnlock()
	if player, ok := WsConnTree.Conns[playerId]; ok {
		player.Lock.Lock()
		defer player.Lock.Unlock()
		player.Conn.WriteMessage(websocket.BinaryMessage, m)
	}
}

// 推送全服
func PushAll(cmd pb.CmdId, resp proto.Message) {
	msg := &pb.Msg{
		Cmd:  cmd,
		Time: general.NowMilli(),
	}
	if resp != nil {
		if b, err := proto.Marshal(resp); err == nil {
			msg.Body = b
		}
	}
	m, _ := proto.Marshal(msg)
	WsConnTree.Lock.RLock()
	defer WsConnTree.Lock.RUnlock()
	for _, player := range WsConnTree.Conns {
		player.Lock.Lock()
		defer player.Lock.Unlock()
		player.Conn.WriteMessage(websocket.BinaryMessage, m)
	}
}
