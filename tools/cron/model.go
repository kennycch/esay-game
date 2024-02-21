package cron

import (
	"github.com/robfig/cron/v3"
)

var (
	// 定时器对象
	cronJob = cron.New(cron.WithSeconds())
	// 任务列表
	jobs = []job{
		{
			spec: "0 * * * * *",
			cmd:  expample,
		},
	}
)

type job struct {
	spec string
	cmd  func()
}
