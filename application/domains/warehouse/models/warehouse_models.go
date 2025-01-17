package models

type ShowWarehouseRequest struct {
	WarehouseName string
	WarehouseId   int64
}

type UpdateConfigWarehouseRequest struct {
	WarehouseId int64
	Config      int
	Operation   string
}
