package usecase

import (
	initServices "butler/application/domains/services/init"
	outboundOrderSv "butler/application/domains/services/outbound_order/service"
	outboundOrderExtendSv "butler/application/domains/services/outbound_order_extend/service"
	"butler/application/lib"
	"butler/config"
)

type usecase struct {
	lib                   *lib.Lib
	cfg                   *config.Config
	outboundOrderSv       outboundOrderSv.IService
	outboundOrderExtendSv outboundOrderExtendSv.IService
}

func InitUseCase(
	lib *lib.Lib,
	cfg *config.Config,
	services *initServices.Services,
) IUseCase {
	return &usecase{
		lib:                   lib,
		cfg:                   cfg,
		outboundOrderSv:       services.OutboundOrderService,
		outboundOrderExtendSv: services.OutboundOrderExtendService,
	}
}
