package errors

import (
	"easy-game/pb"

	"google.golang.org/protobuf/proto"
)

func GetErrorMsg(errorCode pb.ErrorCode) *pb.Msg {
	body := &pb.ErrorBody{
		Code:    errorCode,
		Message: pb.ErrorCode_name[int32(errorCode)],
	}
	b, _ := proto.Marshal(body)
	return &pb.Msg{
		Cmd:  pb.CmdId_CMD_Error,
		Body: b,
	}
}
