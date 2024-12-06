package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"shaanmei_backend/logic"
	"shaanmei_backend/pkg/errcode"
	"shaanmei_backend/pkg/global"
	"shaanmei_backend/pkg/logger"
	"shaanmei_backend/types"
)

func GetMineList(c *gin.Context) {
	reqId := c.GetString(global.RequestIdCtx)

	var pagination types.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		logger.Logger.Errorf("requestId %s parse query error %v", reqId, err)
		c.Set(global.ErrCtx, errcode.ParamError)
		return
	}

	ctx := context.WithValue(c.Request.Context(), global.RequestIdCtx, reqId)

	mines, err := logic.GetMineList(ctx, &types.GetMineListReq{Pagination: pagination})
	if err != nil {
		c.Set(global.ErrCtx, errcode.QueryDatabaseError)
		return
	}

	c.Set(global.DataCtx, mines)
}

func GetCarrierList(c *gin.Context) {
	reqId := c.GetString(global.RequestIdCtx)

	var pagination types.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		logger.Logger.Errorf("requestId %s parse query error %v", reqId, err)
		c.Set(global.ErrCtx, errcode.ParamError)
		return
	}

	ctx := context.WithValue(c.Request.Context(), global.RequestIdCtx, reqId)

	mines, err := logic.GetCarrierList(ctx, &types.GetCarrierListReq{Pagination: pagination})
	if err != nil {
		c.Set(global.ErrCtx, errcode.QueryDatabaseError)
		return
	}

	c.Set(global.DataCtx, mines)
}
