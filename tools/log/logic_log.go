package log

import (
	"easy-game/config"
	"fmt"

	"github.com/rs/zerolog"
)

// 逻辑日志
func LogicLog(level zerolog.Level, info LogInfo) {
	// 处理日志内容
	msg := fmt.Sprintf("%s-%s-%s-%s-%s",
		config.App.RegionName,
		info.PlayerName,
		info.PlayerId,
		info.Module,
		info.Message,
	)
	// 根据日志等级调用指定方法
	switch level {
	case zerolog.DebugLevel:
		Debug(msg)
	case zerolog.InfoLevel:
		Info(msg)
	case zerolog.WarnLevel:
		Warn(msg)
	case zerolog.ErrorLevel:
		Error(msg)
	case zerolog.PanicLevel:
		Panic(msg)
	}
}
