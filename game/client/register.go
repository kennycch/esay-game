package client

import (
	"easy-game/game/errors"
	"easy-game/pb"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kennycch/gotools/log"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func WebsocketUpgrader(ctx *gin.Context) {
	// 升级服务
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Error("Websocket Upgrade error", zap.Error(err))
		conn.Close()
		return
	}
	// 验证入参
	vaild := &Vaild{}
	if err := ctx.ShouldBind(vaild); err != nil {
		log.Error("Params not vaild", zap.Error(err))
		PushConnetError(conn, pb.ErrorCode_EC_VaildFail)
		return
	}
	// 解析Token #######
	playerConn := &PlayerConn{
		Conn:      conn,
		Reconnet:  make(chan bool),
		Disconnet: make(chan bool),
		Msg:       make(chan *pb.Msg),
		PlayerId:  "test001",
		Lock:      &sync.Mutex{},
	}
	RegisterTree(playerConn)
	playerConn.Listen()
}

// 客户端注册树
func RegisterTree(playerConn *PlayerConn) {
	WsConnTree.Lock.Lock()
	defer WsConnTree.Lock.Unlock()
	if oldConn, ok := WsConnTree.Conns[playerConn.PlayerId]; ok {
		oldConn.Reconnet <- true
	}
	WsConnTree.Conns[playerConn.PlayerId] = playerConn
	// 触发连接事件
	ConnectChan <- playerConn.PlayerId
}

// 客户端注销注册
func UnRegisterTree(playerConn *PlayerConn) {
	WsConnTree.Lock.Lock()
	defer WsConnTree.Lock.Unlock()
	delete(WsConnTree.Conns, playerConn.PlayerId)
	// 触发断开事件
	DisconnectChan <- playerConn.PlayerId
}

// 未注册客户端前发送异常
func PushConnetError(conn *websocket.Conn, errorCode pb.ErrorCode) {
	msg := errors.GetErrorMsg(errorCode)
	m, _ := proto.Marshal(msg)
	conn.WriteMessage(websocket.BinaryMessage, m)
	conn.Close()
}
