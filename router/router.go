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

	r.GET("/v1/mine/list")
	r.GET("/v1/carrier/list")

	// 普通用户接口
	normal := r.Group("/v1/normal")
	normal.GET("/price/index")
	normal.GET("/carrier/rank")
	normal.GET("/price/timeline")
	normal.GET("/line/price/index")
	normal.GET("/mine/price/index")
	normal.GET("/line/unitprice")

	// vip用户结接口
	vip := r.Group("/v1/vip")
	vip.GET("/price/timeline")
	vip.GET("/price/index")
	vip.GET("/line/unitprice")
	vip.GET("/price/radar")
	vip.GET("/price/radar")

	_ = r.Run(":" + conf.Conf.GetString("http.port"))
}
