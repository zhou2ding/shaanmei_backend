package handler

import (
	"github.com/gin-gonic/gin"
	"shaanmei_backend/logic"
	"shaanmei_backend/pkg/errcode"
	"shaanmei_backend/pkg/global"
	"shaanmei_backend/pkg/logger"
	"shaanmei_backend/types"
)

type NormalUserHandler struct {
}

func (h *NormalUserHandler) GetPriceIndex(c *gin.Context) {
	reqId := c.GetString(global.RequestIdCtx)

	var req types.NormalUserGetPriceIndexReq
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Logger.Errorf("requestId %s parse query error %v", reqId, err)
		c.Set(global.ErrCtx, errcode.ParamError)
		return
	}

	l := logic.NormalUserLogic{}
	ret, err := l.GetPriceIndex(c.Request.Context(), &req)
	if err != nil {
		c.Set(global.ErrCtx, err)
		return
	}

	c.Set(global.DataCtx, ret)
}

func (h *NormalUserHandler) GetCarrierRank(c *gin.Context) {
}

func (h *NormalUserHandler) GetPriceTimeline(c *gin.Context) {
}

func (h *NormalUserHandler) GetLinePriceIndex(c *gin.Context) {
}

func (h *NormalUserHandler) GetMinePriceIndex(c *gin.Context) {
}

func (h *NormalUserHandler) GetLineUnitPrice(c *gin.Context) {
}
