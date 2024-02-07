package client

import (
	"easy-game/pb"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/serialx/hashring"
)

const (
	PlayerType = iota
	PublicType
)

type Vaild struct {
	Token string `form:"token" binding:"required" json:"token"`
}

type Task struct {
	Player *PlayerConn
	Msg    *pb.Msg
}

// 玩家客户端
type PlayerConn struct {
	Conn      *websocket.Conn // websocket客户端
	Reconnet  chan bool       // 是否重复连接
	Disconnet chan bool       // 断线
	Msg       chan *pb.Msg    // 客户端主动发送信息
	PlayerId  string          // 玩家ID
	Lock      *sync.Mutex     // 写锁
}

// 客户端注册树
type ConnTree struct {
	Conns map[string]*PlayerConn
	Lock  *sync.RWMutex
}

type ChannleType uint8

type ChannelTask struct {
	// 管道类型
	ChannelType ChannleType
	// 管道选择标的
	Target string
}

var (
	// websocket更新
	upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 注册树
	WsConnTree = &ConnTree{
		Conns: map[string]*PlayerConn{},
		Lock:  &sync.RWMutex{},
	}
	// 玩家任务分配映射
	playerChannels = make(map[string]chan *Task, 0)
	// 公共分配映射
	publicChannels = make(map[string]chan *Task, 0)
	// 玩家哈希环
	playerHashring *hashring.HashRing
	// 公共哈希环
	publicHashring *hashring.HashRing
	// 任务分配器映射
	channelTaskMap = map[pb.CmdId]ChannelTask{}
	// 客户端请求黑名单（注册到这里的协议禁止客户端请求，仅供内部调用）
	connBlack = map[pb.CmdId]struct{}{}
	// 连接事件管道
	ConnectChan = make(chan string)
	// 关闭事件管道
	DisconnectChan = make(chan string)
)
