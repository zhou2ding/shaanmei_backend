package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"
	"os"
	"shaanmei_backend/pkg/conf"
	"shaanmei_backend/pkg/logger"
	"shaanmei_backend/pkg/service"
	"shaanmei_backend/pkg/utils"
	"shaanmei_backend/pkg/version"
)

func main() {
	for k, v := range gcmd.GetOptAll() {
		switch k {
		case "h", "help":
			figure.NewFigure("Shipment-Index", "", false).Print()
			return
		case "v", "version":
			version.PrintVersion()
			return
		case "c", "conf", "config":
			conf.InitConf(v)
		}
	}
	if conf.Conf == nil {
		conf.InitConf("./configs/shipment.yaml")
	}

	logger.InitLogger("shipment")

	path, _ := os.Getwd()
	p := &Program{}
	sConfig := &service.Config{
		Name:             "permission",
		DisplayName:      "permission",
		WorkingDirectory: path,
	}
	s, err := service.New(p, sConfig)
	if err != nil {
		logger.Logger.Errorf("create new service error. %s", err)
		return
	}
	//Arg参数  无-
	cmd := gconv.String(gcmd.GetArg(1))
	switch cmd {
	case "help":
		figure.NewFigure("Linkview", "", false).Print()
		return
	case "version":
		version.PrintVersion()
		return
	case "install", "uninstall", "start", "stop", "restart":
		if err = service.Control(s, cmd); err != nil {
			logger.Logger.Errorf("%s\n", err)
			fmt.Printf("%s %s failed.\n", sConfig.DisplayName, cmd)
			utils.PauseExit()
		} else {
			fmt.Printf("%s %s success.\n", sConfig.DisplayName, cmd)
		}
		return
	default:
		if e := s.Run(); e != nil {
			logger.Logger.Errorf("service run error. %s", e)
			utils.PauseExit()
		}
	}
}
