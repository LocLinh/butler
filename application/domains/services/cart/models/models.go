package models

import "time"

type Cart struct {
	CartId      int64     `gorm:"primaryKey;column:cart_id" json:"cart_id"`
	CartCode    string    `gorm:"column:cart_code" json:"cart_code"`
	WarehouseId int64     `gorm:"column:warehouse_id" json:"warehouse_id"`
	Status      int64     `gorm:"column:status" json:"status"`
	CartStt     int64     `gorm:"column:cart_stt" json:"cart_stt"`
	CartDesc    string    `gorm:"column:cart_desc" json:"cart_desc"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy   int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at"`
	UpdatedBy   int64     `gorm:"column:updated_by" json:"updated_by"`

	// Custom fields
	UpdatedByUserName string `gorm:"->" json:"updated_by_user_name"`
	OrderQuantity     int64  `gorm:"->" json:"order_quantity"`
	WarehouseName     string `gorm:"->" json:"warehouse_name"`
	StatusName        string `gorm:"->" json:"status_name"`
}

func (c *Cart) TableName() string {
	return "wms_cart"
}

type GetRequest struct {
	CartId      int64
	CartCode    string
	WarehouseId int64
	Status      int64
	Statuses    []int64
	UpdatedBy   int64
}
