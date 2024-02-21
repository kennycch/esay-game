package log

import (
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog"
)

const (
	fileMaxage   = 7 * 24 * time.Hour
	rotationTime = 24 * time.Hour
	chanSize     = 1000
)

const (
	stdoutType outputType = "stdout"
	stderrType outputType = "stderr"
	fileType   outputType = "file"
)

var (
	kafkaHooker     kafkaHook
	kafkaHookerOnce sync.Once
	logger          *Logger
)

type kafkaInfo struct {
	serverTopic    string
	statisticTopic string
	producer       sarama.SyncProducer
}

type kafkaHook struct {
	serverChan chan string
	kafkaInfo  *kafkaInfo
}

type outputType string

type Logger struct {
	log  zerolog.Logger
	name string
}

type LogInfo struct {
	PlayerName string // 用户名
	PlayerId   string // 用户Id
	Module     string // 模块
	Message    string // 事件
}
