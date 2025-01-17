package usecase

import (
	whModels "butler/application/domains/services/warehouse/models"
	"butler/application/domains/warehouse/models"
	"context"
)

type IUseCase interface {
	ShowWarehouse(ctx context.Context, params *models.ShowWarehouseRequest) error
	ResetShowWarehouse(ctx context.Context, warehouseId int64) (string, error)
	GetWarehouseById(ctx context.Context, warehouseId int64) (*whModels.Warehouse, error)
	RemoveConfigWarehouse(ctx context.Context, params *models.UpdateConfigWarehouseRequest) error
	AddConfigWarehouse(ctx context.Context, params *models.UpdateConfigWarehouseRequest) error
}
