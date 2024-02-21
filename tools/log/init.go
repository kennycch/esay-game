package log

import "easy-game/tools/lifecycle"

type LogRegister struct{}

func (e *LogRegister) Start() {
	NewKafkaHooker()
	NewLogger()
}

func (e *LogRegister) Priority() uint32 {
	return lifecycle.HighPriority + 100
}

func (e *LogRegister) Stop() {

}

func NewLogRegister() *LogRegister {
	return &LogRegister{}
}
