package usecase

import (
	"butler/application/domains/pick_pack/models"
	"context"
)

type IUseCase interface {
	AutoPickPack(ctx context.Context, params models.AutoPickPackRequest) (string, error)
}
