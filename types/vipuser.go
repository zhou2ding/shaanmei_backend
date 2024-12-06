package types

type (
	VipUserGetPriceIndexReq struct {
		Date    string  `form:"date"`
		Mine    int32   `form:"mine"`
		DestLng float32 `form:"dest_lng"`
		DestLat float32 `form:"dest_lat"`
	}
	VipUserGetPriceIndexRes struct {
		Date       string  `form:"date"`
		PriceIndex float32 `form:"price_index"`
		Mine       int32   `form:"mine"`
		City       string  `form:"city"`
		MineLng    float32 `form:"mine_lng"`
		MineLat    float32 `form:"mine_lat"`
	}
)

type (
	VipUserGetPriceTimelineReq struct {
		Type     string   `form:"type" binding:"required,oneof=common carrier mine"`
		Carriers []string `form:"carriers"`
		Mines    []string `form:"mines"`
	}
	VipUserGetPriceTimelineRes struct {
		TimeRange *TimeRange `form:"time_range"`
		Points    []*Point   `form:"points"`
	}
)

type (
	VipUserGetLineUnitPriceReq struct {
		Mine     int32   `form:"mine"`
		EndLng   float32 `form:"end_lng"`
		EndLat   float32 `form:"end_lat"`
		Distance float32 `form:"distance"`
	}
	VipUserGetLineUnitPriceRes struct {
		ID        int32   `json:"id"`
		Name      string  `json:"name"`
		UnitPrice float32 `json:"unitprice"`
	}
)

type (
	VipUserGetPriceRadarRes struct {
		List []*VipUserRadar `json:"list"`
	}
	VipUserRadar struct {
		Title  string               `json:"title"`
		Points []*VipUserRadarPoint `json:"points"`
	}
	VipUserRadarPoint struct {
		Indicator string  `json:"indicator"`
		Value     float32 `json:"value"`
	}
)

type (
	VipUserGetPriceHeatmapRes struct {
		List []*VipUserHeatmap `json:"list"`
	}
	VipUserHeatmap struct {
		Title string                `json:"title"`
		Grids []*VipUserHeatmapGrid `json:"grids"`
	}
	VipUserHeatmapGrid struct {
		XName        string  `json:"x_name"`
		YName        string  `json:"y_name"`
		AvgUnitprice float32 `json:"avg_unitprice"`
		AvgDistance  float32 `json:"avg_distance"`
	}
)
