package types

type Pagination struct {
	Page  int `form:"page" binding:"omitempty,gte=1"`
	Limit int `form:"limit" binding:"omitempty,gte=1"`
}

type (
	GetMineListReq struct {
		Pagination
	}
	GetMineListResp struct {
		List  []*Mine `json:"list"`
		Total int     `json:"total"`
	}
	Mine struct {
		ID        int32   `json:"id"`
		Name      string  `json:"name"`
		City      string  `json:"city"`
		Longitude float32 `json:"longitude"`
		Latitude  float32 `json:"latitude"`
	}
)

type (
	GetCarrierListReq struct {
		Pagination
	}
	GetCarrierListResp struct {
		List  []*Carrier `json:"list"`
		Total int        `json:"total"`
	}
	Carrier struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	}
)

type TimeRange struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Point struct {
	Time               string  `json:"time"`
	WeightedAvgPrice   float32 `json:"weighted_avg_price"`
	WeightedPriceIndex float32 `json:"weighted_price_index"`
}
