package handler

import (
	"github.com/gin-gonic/gin"
	"shaanmei_backend/pkg/errcode"
	"shaanmei_backend/pkg/global"
	"shaanmei_backend/pkg/logger"
	"shaanmei_backend/types"
)

type VipUserHandler struct{}

func (h *VipUserHandler) GetPriceTimeline(c *gin.Context) {
	reqId := c.GetString(global.RequestIdCtx)

	var req types.VipUserGetPriceTimelineReq
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Logger.Errorf("requestId %s parse query error %v", reqId, err)
		c.Set(global.ErrCtx, errcode.ParamError)
		return
	}
}

func (h *VipUserHandler) GetPriceIndex(c *gin.Context) {
	reqId := c.GetString(global.RequestIdCtx)

	var req types.VipUserGetPriceIndexReq
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Logger.Errorf("requestId %s parse query error %v", reqId, err)
		c.Set(global.ErrCtx, errcode.ParamError)
		return
	}
}

func (h *VipUserHandler) GetLineUnitPrice(c *gin.Context) {
	reqId := c.GetString(global.RequestIdCtx)

	var req types.VipUserGetLineUnitPriceReq
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Logger.Errorf("requestId %s parse query error %v", reqId, err)
		c.Set(global.ErrCtx, errcode.ParamError)
		return
	}
}

func (h *VipUserHandler) GetPriceRadar(c *gin.Context) {
}

func (h *VipUserHandler) GetPriceHeatmap(c *gin.Context) {

}
