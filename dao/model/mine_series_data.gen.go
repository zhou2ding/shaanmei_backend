// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMineSeriesDatum = "mine_series_data"

// MineSeriesDatum mapped from table <mine_series_data>
type MineSeriesDatum struct {
	ID                  int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name                string    `gorm:"column:name" json:"name"`
	Longitude           float32   `gorm:"column:longitude" json:"longitude"`
	Latitude            float32   `gorm:"column:latitude" json:"latitude"`
	AvgShipmentPrice    float32   `gorm:"column:avg_shipment_price" json:"avg_shipment_price"`
	TotalShipmentVolume float32   `gorm:"column:total_shipment_volume" json:"total_shipment_volume"`
	Timestamp           time.Time `gorm:"column:timestamp" json:"timestamp"`
}

// TableName MineSeriesDatum's table name
func (*MineSeriesDatum) TableName() string {
	return TableNameMineSeriesDatum
}
