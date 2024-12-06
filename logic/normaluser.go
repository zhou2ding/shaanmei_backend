package logic

import (
	"context"
	"fmt"
	"shaanmei_backend/types"
	"time"
)

type NormalUserLogic struct{}

func (l *NormalUserLogic) GetPriceIndex(ctx context.Context, req *types.NormalUserGetPriceIndexReq) (*types.NormalUserGetPriceIndexResp, error) {
	fmt.Println(req.Date)
	fmt.Println(req.Type)
	res := &types.NormalUserGetPriceIndexResp{Date: time.Now().Format(time.DateTime), PriceIndex: 10}
	return res, nil
}
