package usecase

import (
	"butler/application/domains/pick_pack/models"
	"context"
)

type IUseCase interface {
	PickPackKafka(ctx context.Context, params *models.AutoPickPackRequest) error
	SetOutboundOrderVoucherType(ctx context.Context, params *models.SetOutboundOrderVoucherTypeRequest) error
}
