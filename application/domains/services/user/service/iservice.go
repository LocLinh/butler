package service

import (
	"butler/application/domains/services/user/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, params *models.GetRequest) (*models.User, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.User, error)
	Update(ctx context.Context, obj *models.User) (*models.User, error)
}
