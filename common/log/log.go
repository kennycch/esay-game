package log

import (
	"easy-game/config"
	"io"
	"os"

	"github.com/rs/zerolog"
)

type outputType string

const (
	stdoutType outputType = "stdout"
	stderrType outputType = "stderr"
	fileType   outputType = "file"
)

var logger *Logger

type Logger struct {
	log  zerolog.Logger
	name string
}

func NewLogger() {
	var zerologger zerolog.Logger
	if config.App.Debug { // 如果是全局debug=true模式 则默认输出到标准输出以及kafka
		zerologger = zerolog.New(setLogOutput(stdoutType))
	} else { // 否则默认输出到文件以及输出到kafka
		zerologger = zerolog.New(setLogOutput(stdoutType))
	}

	logger = &Logger{
		name: config.App.RegionName,
		log: zerologger.With().
			Timestamp().
			Caller().
			Logger().
			Hook(kafkaHooker),
	}
}

func setLogOutput(level outputType) io.Writer {
	switch level {
	case stderrType:
		return os.Stderr
	case fileType:
		return newFile(config.Log.LogPath)
	}
	return os.Stdout
}
