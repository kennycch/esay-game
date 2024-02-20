package errors

import "easy-game/pb"

var (
	Errors = map[pb.ErrorCode]string{
		pb.ErrorCode_EC_VaildFail:   "参数异常",
		pb.ErrorCode_EC_OtherClient: "账户在其他地方登陆",
	}
)
