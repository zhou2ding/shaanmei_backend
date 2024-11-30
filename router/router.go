package router

import (
	"github.com/gin-gonic/gin"
	mw "shaanmei_backend/middleware"
	"shaanmei_backend/pkg/conf"
	"shaanmei_backend/pkg/logger"
)

func NewHttpServer() {
	logger.Logger.Info("start http server...")
	if conf.Conf.GetBool("log.stdout") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(mw.CommonResponse())
	r.NoRoute(mw.NotFound())

	// 不应用JWTAuth中间件的接口
	r.POST("/v1/mine/list")

	_ = r.Run(":" + conf.Conf.GetString("http.port"))
}
