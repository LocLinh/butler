package repository

import (
	"butler/application/domains/services/bin_location_cart_mapping/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.BinLocationCartMapping) (*models.BinLocationCartMapping, error)
	UpdateMany(ctx context.Context, objs []*models.BinLocationCartMapping) error
	GetById(ctx context.Context, id int64) (*models.BinLocationCartMapping, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.BinLocationCartMapping, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.BinLocationCartMapping, error)
	Create(ctx context.Context, obj *models.BinLocationCartMapping) (*models.BinLocationCartMapping, error)
	Delete(ctx context.Context, id int64) error
}
