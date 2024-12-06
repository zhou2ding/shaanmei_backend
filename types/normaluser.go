package types

type (
	NormalUserGetPriceIndexReq struct {
		Date string `form:"date" binding:"required"`
		Type string `form:"type" binding:"required,oneof=short medium long"`
	}
	NormalUserGetPriceIndexResp struct {
		Date       string  `json:"date"`
		PriceIndex float32 `json:"price_index"`
	}
)

type (
	NormalUserGetCarrierRankResp struct {
		List  []*CarrierRank `json:"list"`
		Total int            `json:"total"`
	}
	CarrierRank struct {
		Type     string             `json:"type"`
		RankList []*CarrierRankItem `json:"rank_list"`
	}
	CarrierRankItem struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
		Rank int    `json:"rank"`
	}
)

type (
	NormalUserGetPriceTimelineRes struct {
		TimeRange *TimeRange            `json:"time_range"`
		List      []*NormalUserTimeline `json:"list"`
	}
	NormalUserTimeline struct {
		Title  string   `json:"title"`
		Points []*Point `json:"points"`
	}
)

type (
	NormalUserGetLinePriceIndexRes struct {
		LastPeriodDate    string                      `json:"last_period_date"`
		CurrentPeriodDate string                      `json:"current_period_date"`
		List              []*NormalUserLinePriceIndex `json:"list"`
	}
	NormalUserLinePriceIndex struct {
		Name                    string  `json:"name"`
		LastPeriodPriceIndex    float32 `json:"last_period_price_index"`
		CurrentPeriodPriceIndex float32 `json:"current_period_price_index"`
		ChangePercentage        float32 `json:"change_percentage"`
	}
)

type (
	NormalUserGetMinePriceIndexRes struct {
		List []*NormalUserMinePriceIndex `json:"list"`
	}
	NormalUserMinePriceIndex struct {
		Name       string  `json:"name"`
		Lng        float32 `json:"lng"`
		Lat        float32 `json:"lat"`
		PriceIndex float32 `json:"price_index"`
	}
)

type (
	NormalUserGetLineUnitPriceRes struct {
		List []*NormalUserLineUnitPrice `json:"list"`
	}
	NormalUserLineUnitPrice struct {
		Name      string  `json:"name"`
		StartLng  float32 `json:"start_lng"`
		StartLat  float32 `json:"start_lat"`
		DestLng   float32 `json:"dest_lng"`
		DestLat   float32 `json:"dest_lat"`
		UnitPrice float32 `json:"unitprice"`
	}
)
