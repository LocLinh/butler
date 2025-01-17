package models

import (
	"time"
)

type BinLocationCartMapping struct {
	Id        int64     `gorm:"primaryKey;column:id"`
	Location  string    `gorm:"column:location"`
	CartCode  string    `gorm:"column:cart_code"`
	UsedBy    int64     `gorm:"column:used_by"`
	Purpose   string    `gorm:"column:purpose_of_use"`
	AreaCode  string    `gorm:"column:area_code"`
	Status    int64     `gorm:"column:status"`
	CreatedAt time.Time `gorm:"autoCreateTime" jenk:"filter:-;range:from-to;type:datetime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:(-)" jenk:"filter:-"`
}

func (t *BinLocationCartMapping) TableName() string {
	return "wms_bin_location_cart_mapping"
}

type GetRequest struct {
	Id          int64
	CartCode    string
	WarehouseId int64
	Status      int64
	UsedBy      int64
}
