package logger

import (
	"easy-game/config"
	"easy-game/lifecycle"

	"github.com/kennycch/gotools/log"
)

type Logger struct{}

func (l *Logger) Start() {
	log.InitLog(
		config.Log.LogPath,
		config.App.AppName,
		config.Log.LogDay,
		log.Level(config.Log.LogLevel))
}

func (l *Logger) Priority() uint32 {
	return lifecycle.HighPriority + 9999
}

func (l *Logger) Stop() {

}

func NewLogger() *Logger {
	return &Logger{}
}
