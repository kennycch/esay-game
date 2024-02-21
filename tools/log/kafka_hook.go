package log

import (
	"easy-game/config"
	"time"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog"
)

func NewKafkaHooker() {
	kafkaHookerOnce.Do(func() {
		kafkaHooker = kafkaHook{
			serverChan: make(chan string, chanSize),
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
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Retry.Max = 5
	conf.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{config.Kafka.Host}, conf)
	if err != nil {
		Panic("kafka初始化失败")
	}
	kh.kafkaInfo = &kafkaInfo{
		serverTopic:    config.Kafka.ServerTopic,
		statisticTopic: config.Kafka.StatisticTopic,
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
