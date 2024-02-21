// 日志格式规范尽量做到：区服名称-用户名称-用户id-什么功能模块-发生了什么
package log

func Debug(msg string) {
	logger.log.Debug().Str("regionName", logger.name).Msg(msg)
}

func Info(msg string) {
	logger.log.Info().Str("regionName", logger.name).Msg(msg)
}

func Warn(msg string) {
	logger.log.Warn().Str("regionName", logger.name).Msg(msg)
}

func Error(msg string) {
	logger.log.Error().Str("regionName", logger.name).Msg(msg)
}

func Panic(msg string) {
	logger.log.Panic().Str("regionName", logger.name).Msg(msg)
}

// Statistic 上报统计信息(如道具日志、用户操作日志)
func Statistic(msg string) {
	logger.log.Log().Str("regionName", logger.name).Msg(msg)
}
