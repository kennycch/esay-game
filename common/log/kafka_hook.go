package log

import (
	"easy-game/config"
	_config "easy-game/config"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog"
)

const (
	chansize = 1000
)

var kafkaHooker kafkaHook
var kafkaHookerOnce sync.Once

type kafkaInfo struct {
	serverTopic    string
	statisticTopic string
	producer       sarama.SyncProducer
}

type kafkaHook struct {
	serverChan chan string
	kafkaInfo  *kafkaInfo
}

func NewKafkaHooker() {
	kafkaHookerOnce.Do(func() {
		kafkaHooker = kafkaHook{
			serverChan: make(chan string, 1000),
		}
		if config.Log.RemoteEnable {
			// 初始化kafka生产者
			kafkaHooker.initProducer()
			// 启动服务日志上传服务
			kafkaHooker.runServerLogUpload()
		}
	})
}

func (kh *kafkaHook) initProducer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{_config.Kafka.Host}, config)
	if err != nil {
		Panic("kafka初始化失败")
	}
	kh.kafkaInfo = &kafkaInfo{
		serverTopic:    _config.Kafka.ServerTopic,
		statisticTopic: _config.Kafka.StatisticTopic,
		producer:       producer,
	}
}

func (kh *kafkaHook) runServerLogUpload() {
	go func() {
		for {
			select {
			case logMsg := <-kh.serverChan:
				kh.sendMsg2kafka(kh.kafkaInfo.serverTopic, logMsg)
			default:
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()
}

func (kh *kafkaHook) sendMsg2kafka(topic, msg string) error {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	_, _, err := kh.kafkaInfo.producer.SendMessage(message)
	return err
}

func (kh kafkaHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level == zerolog.NoLevel { // 运营业务的日志 即时上传
		go kh.sendMsg2kafka(kh.kafkaInfo.statisticTopic, msg)
	} else { // 服务日志
		select {
		case kh.serverChan <- msg:
		default:
		}
	}
}
