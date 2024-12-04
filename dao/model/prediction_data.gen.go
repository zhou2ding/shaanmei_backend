// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePredictionDatum = "prediction_data"

// PredictionDatum mapped from table <prediction_data>
type PredictionDatum struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CarrierName     string    `gorm:"column:carrier_name" json:"carrier_name"`
	RecipientName   string    `gorm:"column:recipient_name" json:"recipient_name"`
	MineName        string    `gorm:"column:mine_name" json:"mine_name"`
	MineCounty      string    `gorm:"column:mine_county" json:"mine_county"`
	RecipientCounty string    `gorm:"column:recipient_county" json:"recipient_county"`
	ProductType     string    `gorm:"column:product_type" json:"product_type"`
	VehicleType     string    `gorm:"column:vehicle_type" json:"vehicle_type"`
	LoadWeight      float32   `gorm:"column:load_weight" json:"load_weight"`
	Distance        float32   `gorm:"column:distance" json:"distance"`
	ShipmentCost    float32   `gorm:"column:shipment_cost" json:"shipment_cost"`
	ArrivalTime     time.Time `gorm:"column:arrival_time" json:"arrival_time"`
}

// TableName PredictionDatum's table name
func (*PredictionDatum) TableName() string {
	return TableNamePredictionDatum
}