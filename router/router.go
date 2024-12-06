package router

import (
	"github.com/gin-gonic/gin"
	"shaanmei_backend/handler"
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

	r.GET("/v1/mine/list", handler.GetMineList)
	r.GET("/v1/carrier/list", handler.GetCarrierList)

	// 普通用户接口
	normalUserHandler := handler.NormalUserHandler{}
	normal := r.Group("/v1/normal")
	normal.GET("/price/index", normalUserHandler.GetPriceIndex)
	normal.GET("/carrier/rank", normalUserHandler.GetCarrierRank)
	normal.GET("/price/timeline", normalUserHandler.GetPriceTimeline)
	normal.GET("/line/price/index", normalUserHandler.GetLinePriceIndex)
	normal.GET("/mine/price/index", normalUserHandler.GetMinePriceIndex)
	normal.GET("/line/unitprice", normalUserHandler.GetLineUnitPrice)

	// vip用户结接口
	vipUserHandler := handler.VipUserHandler{}
	vip := r.Group("/v1/vip")
	vip.GET("/price/timeline", vipUserHandler.GetPriceTimeline)
	vip.GET("/price/index", vipUserHandler.GetPriceIndex)
	vip.GET("/line/unitprice", vipUserHandler.GetLineUnitPrice)
	vip.GET("/price/radar", vipUserHandler.GetPriceRadar)
	vip.GET("/price/heatmap", vipUserHandler.GetPriceHeatmap)

	_ = r.Run(":" + conf.Conf.GetString("http.port"))
}
