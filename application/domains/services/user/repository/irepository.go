package repository

import (
	"butler/application/domains/services/user/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.User) (*models.User, error)
	GetById(ctx context.Context, id int64) (*models.User, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.User, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.User, error)
}
