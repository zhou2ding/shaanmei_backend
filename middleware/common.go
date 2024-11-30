package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shaanmei_backend/pkg/errcode"
	"shaanmei_backend/pkg/global"
	"shaanmei_backend/pkg/utils"
)

func CommonResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(global.RequestIdCtx, utils.GetUUIDFull())
		c.Next()

		if c.Writer.Written() {
			return
		}

		var (
			reqId = c.GetString(global.RequestIdCtx)
			resp  = Response{RequestId: reqId}
		)
		if err, ok := c.Get(global.ErrCtx); ok {
			ec := err.(errcode.ErrCode)
			resp.Code, resp.Message = int(ec), ec.String()

			errMsg, _ := c.Get(global.ErrMsgCtx)
			if errMsg != nil {
				resp.Message = errMsg
				c.JSON(http.StatusBadRequest, resp)
			} else {
				c.JSON(http.StatusInternalServerError, resp)
			}
		} else {
			data, _ := c.Get(global.DataCtx)
			resp.Code, resp.Data = http.StatusOK, data
			c.JSON(http.StatusOK, resp)
		}
	}
}

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
}
