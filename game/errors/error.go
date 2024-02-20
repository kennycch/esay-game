package errors

import (
	"easy-game/pb"

	"google.golang.org/protobuf/proto"
)

func GetErrorMsg(errorCode pb.ErrorCode) *pb.Msg {
	message := ""
	if msg, ok := Errors[errorCode]; ok {
		message = msg
	}
	body := &pb.ErrorBody{
		Code:    errorCode,
		Message: message,
	}
	b, _ := proto.Marshal(body)
	return &pb.Msg{
		Cmd:  pb.CmdId_CMD_Error,
		Body: b,
	}
}
