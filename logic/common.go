package logic

import (
	"context"
	"shaanmei_backend/dao/model"
	"shaanmei_backend/pkg/database"
	"shaanmei_backend/pkg/global"
	"shaanmei_backend/pkg/logger"
	"shaanmei_backend/types"
)

func GetMineList(ctx context.Context, req *types.GetMineListReq) ([]*types.Mine, error) {
	subQuery := database.GetDB().Model(model.MineSeriesDatum{}).
		Select("MIN(id) AS min_id").
		Group("mine_name")

	query := database.GetDB().Model(model.MineSeriesDatum{}).InnerJoins("JOIN (?) AS sub ON mine_series_data.id = sub.min_id", subQuery)
	if req.Page != 0 {
		query = query.Offset(req.Page)
	}
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}

	var mines []*model.MineSeriesDatum
	if err := query.Find(&mines).Error; err != nil {
		logger.Logger.Errorf("requestId: %v find mines in database error: %v", ctx.Value(global.RequestIdCtx), err)
		return nil, err
	}

	if len(mines) == 0 {
		return nil, nil
	}

	res := make([]*types.Mine, 0)
	for _, mine := range mines {
		res = append(res, &types.Mine{
			ID:        mine.ID,
			Name:      mine.MineName,
			Longitude: mine.MineLongitude,
			Latitude:  mine.MineLatitude,
		})
	}
	return res, nil
}

func GetCarrierList(ctx context.Context, req *types.GetCarrierListReq) ([]*types.Carrier, error) {
	query := database.GetDB().Model(model.CarrierSeriesDatum{}).Distinct("carrier_name")
	if req.Page != 0 {
		query = query.Offset(req.Page)
	}
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}

	var carriers []*model.CarrierSeriesDatum
	if err := query.Find(&carriers).Error; err != nil {
		logger.Logger.Errorf("requestId: %v find carriers in database error: %v", ctx.Value(global.RequestIdCtx), err)
		return nil, err
	}

	res := make([]*types.Carrier, 0)
	for _, carrier := range carriers {
		res = append(res, &types.Carrier{
			ID:   carrier.ID,
			Name: carrier.CarrierName,
		})
	}

	return res, nil
}
