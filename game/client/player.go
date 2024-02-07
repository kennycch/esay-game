package client

import (
	"easy-game/game/errors"
	"easy-game/pb"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kennycch/gotools/log"
	"github.com/kennycch/gotools/worker"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (p *PlayerConn) Listen() {
	p.ReadMsg()
	worker.AddTask(func() {
	label:
		for {
			select {
			// 客户端断线
			case <-p.Disconnet:
				p.Conn.Close()
				UnRegisterTree(p)
				break label
			// 客户端多点登录
			case <-p.Reconnet:
				msg := errors.GetErrorMsg(pb.ErrorCode_EC_OtherClient)
				PushByPlayer(p.PlayerId, msg)
				p.Conn.Close()
				break label
			// 客户端发送信息
			case msg := <-p.Msg:
				// 获取任务管道
				_, blackList := connBlack[msg.Cmd]
				if channelTask, ok := channelTaskMap[msg.Cmd]; ok && !blackList {
					// 玩家策略自动添加玩家ID
					if channelTask.ChannelType == PlayerType {
						channelTask.Target = p.PlayerId
					}
					channel := GetTaskChannle(channelTask)
					task := &Task{
						Player: p,
						Msg:    msg,
					}
					channel <- task
				}
			}
		}
	})
}

// 读取信息
func (p *PlayerConn) ReadMsg() {
	worker.AddTask(func() {
		for {
			// 设置心跳时间
			p.Conn.SetReadDeadline(time.Now().Add(time.Minute))
			// 读取信息
			mType, m, err := p.Conn.ReadMessage()
			if err != nil { // 读取异常触发关闭事件
				log.Error("read message err", zap.Error(err))
				p.Disconnet <- true
				break
			} else if mType != websocket.BinaryMessage { // 数据类型不符，忽略本次信息
				continue
			} else { // 解析信息
				msg := &pb.Msg{}
				if err = proto.Unmarshal(m, msg); err == nil {
					p.Msg <- msg
				}
			}
		}
	})
}
