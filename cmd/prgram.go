package main

import (
	"shaanmei_backend/pkg/cache"
	"shaanmei_backend/pkg/database"
	"shaanmei_backend/pkg/logger"
	"shaanmei_backend/pkg/service"
	"shaanmei_backend/router"
)

type Program struct {
}

func (p *Program) Start(s service.Service) (e error) {
	logger.Logger.Warnf("start programme...")
	if db := database.InitIotDB(); db != nil {
		logger.Logger.Infof("connect to tidb success.")
	} else {
		_ = p.Stop(s)
	}
	if err := cache.InitRedis(); err != nil {
		_ = p.Stop(s)
	}
	go router.NewHttpServer()
	return nil
}

func (p *Program) Stop(s service.Service) (e error) {
	logger.Logger.Warnf("stop programme...")
	return nil
}
