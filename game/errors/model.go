package errors

import "easy-game/pb"

var (
	Errros = map[pb.ErrorCode]string{
		pb.ErrorCode_EC_VaildFail: "参数异常",
	}
)
