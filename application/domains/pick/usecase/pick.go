package usecase

import (
	"butler/application/domains/pick/models"
	outboundModel "butler/application/domains/services/outbound_order/models"
	"butler/constants"
	"context"
	"fmt"
)

func (u *usecase) PreparePick(ctx context.Context) error {
	return nil
}

func (u *usecase) Pick(ctx context.Context, params *models.PickRequest) error {
	if err := u.lib.Validator.Struct(params); err != nil {
		return err
	}

	outbound, err := u.outboundOrderSv.GetOne(ctx, &outboundModel.GetRequest{SalesOrderNumber: params.SalesOrderNumber})
	if err != nil {
		return err
	}
	if outbound.StatusId != constants.OUTBOUND_ORDER_STATUS_PICKLISTED {
		return fmt.Errorf("outbound order [%s] không ở trạng thái picklisted", params.SalesOrderNumber)
	}
	if outbound.Config != 0 {
		if outbound.Config&constants.OUTBOUND_ORDER_CONFIG_NOT_ENOUGH_QTY != 0 {
			return fmt.Errorf("outbound order [%s] không đủ hàng đi pick", params.SalesOrderNumber)
		}

		validConfigs := constants.OUTBOUND_ORDER_CONFIG_NOT_VERIFIED | constants.OUTBOUND_ORDER_CONFIG_HAS_REASON_DELAY | constants.OUTBOUND_ORDER_CONFIG_HAS_SALES_OFF | constants.OUTBOUND_ORDER_CONFIG_HAS_PENDING | constants.OUTBOUND_ORDER_CONFIG_HAS_FORCE_PICKING | constants.OUTBOUND_ORDER_CONFIG_REQUIRED_EXPORT_INSIDE | constants.OUTBOUND_ORDER_CONFIG_WRONG_COMBO | constants.OUTBOUND_ORDER_CONFIG_HAS_SALES_OFF_COMBO | constants.OUTBOUND_ORDER_CONFIG_WAITING_REPICK | constants.OUTBOUND_ORDER_CONFIG_NOT_PICKLISTED_UID | constants.OUTBOUND_ORDER_CONFIG_IT_TESTER | constants.OUTBOUND_ORDER_CONFIG_WAITING_SHOP_CONFIRM | constants.OUTBOUND_ORDER_CONFIG_WAITING_CS_CONFIRM | constants.OUTBOUND_ORDER_CONFIG_CS_PROCESSING | constants.OUTBOUND_ORDER_CONFIG_ALLOW_MISSING_PICKUP | constants.OUTBOUND_ORDER_CONFIG_COMPLETED
		if outbound.Config&validConfigs == 0 {
			return fmt.Errorf("outbound order [%s] không đủ điều kiện pick, order config %s",
				params.SalesOrderNumber,
				constants.OUTBOUND_ORDER_CONFIG_MAP_NAME[outbound.Config])
		}
	}

	if err := u.ReadyPickOutbound(ctx, &models.ReadyPickOutboundRequest{SalesOrderNumber: params.SalesOrderNumber}); err != nil {
		return err
	}

	if outbound.OutboundOrderType == constants.OUTBOUND_ORDER_TYPE_ORDER {
		return u.PickOrder(ctx, outbound)
	}
	if outbound.OutboundOrderType == constants.OUTBOUND_ORDER_TYPE_INTERNAL_TRANSFER {
		return u.PickIt(ctx, outbound)
	}

	return nil
}

func (u *usecase) PickOrder(ctx context.Context, outbound *outboundModel.OutboundOrder) error {
	return nil
}

func (u *usecase) PickIt(ctx context.Context, outbound *outboundModel.OutboundOrder) error {
	return nil
}
