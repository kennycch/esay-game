package cron

import (
	"easy-game/tools/lifecycle"
	"easy-game/tools/log"
)

type Cron struct{}

func (n *Cron) Start() {
	for _, job := range jobs {
		if _, err := cronJob.AddFunc(job.spec, job.cmd); err != nil {
			log.Panic(err.Error())
		}
	}
	cronJob.Start()
}

func (n *Cron) Priority() uint32 {
	return lifecycle.LowPriority + 100
}

func (n *Cron) Stop() {

}

func NewCron() *Cron {
	return &Cron{}
}
