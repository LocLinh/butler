package usecase

import (
	"butler/application/domains/cart/models"
	"context"
)

type IUseCase interface {
	ResetCart(ctx context.Context, params *models.ResetCartRequest) error
	ResetCartByUserId(ctx context.Context, params *models.ResetCartByUserIdRequest) (string, error)
	ResetCartByEmail(ctx context.Context, params *models.ResetCartByEmailRequest) (string, error)
}
