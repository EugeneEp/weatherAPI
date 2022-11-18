package service

import (
	"github.com/zhashkevych/scheduler"
	"go.uber.org/zap"
	"projectname/internal/project/infrastructure/background"
)

func (a *app) startWorkers() {
	var workers *scheduler.Scheduler

	if err := a.ctn.Fill(background.ServiceName, &workers); err != nil {
		a.logger.Error(`Failed To Start Background Workers`, zap.Error(err))
	}
}
