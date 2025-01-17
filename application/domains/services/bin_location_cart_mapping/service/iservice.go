package service

import (
	"butler/application/domains/services/bin_location_cart_mapping/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, params *models.GetRequest) (*models.BinLocationCartMapping, error)
	Create(ctx context.Context, obj *models.BinLocationCartMapping) (*models.BinLocationCartMapping, error)
	GetById(ctx context.Context, id int64) (*models.BinLocationCartMapping, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.BinLocationCartMapping, error)
	Update(ctx context.Context, obj *models.BinLocationCartMapping) (*models.BinLocationCartMapping, error)
	Delete(ctx context.Context, id int64) error
}
