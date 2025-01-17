package usecase

import (
	"butler/application/domains/pick/models"
	binLocationModel "butler/application/domains/services/bin_location/models"
	binLocationSv "butler/application/domains/services/bin_location/service"
	initServices "butler/application/domains/services/init"
	invenoryModel "butler/application/domains/services/inventory/models"
	invenorySv "butler/application/domains/services/inventory/service"
	outboundModel "butler/application/domains/services/outbound_order/models"
	outboundOrderSv "butler/application/domains/services/outbound_order/service"
	pickingModel "butler/application/domains/services/picking/models"
	pickingSv "butler/application/domains/services/picking/service"
	pickingItemModel "butler/application/domains/services/picking_item/models"
	pickingItemSv "butler/application/domains/services/picking_item/service"
	"butler/application/lib"
	"butler/constants"
	"context"
	"fmt"
	"strings"
	"time"
)

type usecase struct {
	lib             *lib.Lib
	outboundOrderSv outboundOrderSv.IService
	pickingSv       pickingSv.IService
	pickingItemSv   pickingItemSv.IService
	invenorySv      invenorySv.IService
	binLocationSv   binLocationSv.IService
}

func InitUseCase(
	lib *lib.Lib,
	services *initServices.Services,
) IUseCase {
	return &usecase{
		lib:             lib,
		pickingSv:       services.PickingService,
		pickingItemSv:   services.PickingItemService,
		invenorySv:      services.InventoryService,
		outboundOrderSv: services.OutboundOrderService,
		binLocationSv:   services.BinLocationService,
	}
}

const DEFAULT_LOCATION = "F0-AP-00-00-00-01"

func (u *usecase) ReadyPickOutbound(ctx context.Context, params *models.ReadyPickOutboundRequest) error {
	outbound, err := u.outboundOrderSv.GetOne(ctx, &outboundModel.GetRequest{SalesOrderNumber: params.SalesOrderNumber})
	if err != nil {
		return err
	}
	if outbound == nil || outbound.OutboundOrderId == 0 {
		return fmt.Errorf("không tìm thấy mã outbound với source code [%s]", params.SalesOrderNumber)
	}

	if outbound.StatusId != 1 {
		return fmt.Errorf("outbound order [%s] không ở trạng thái picklisted", params.SalesOrderNumber)
	}
	if outbound.Config != 0 {
		if outbound.Config&constants.OUTBOUND_ORDER_CONFIG_NOT_ENOUGH_QTY != 0 {
			return fmt.Errorf("outbound order [%s] không đủ hàng đi pick", params.SalesOrderNumber)
		}

	}

	picking, err := u.pickingSv.GetOne(ctx, &pickingModel.GetRequest{OutboundOrderId: outbound.OutboundOrderId})
	if err != nil {
		return err
	}
	if picking == nil || picking.PickingId == 0 {
		return fmt.Errorf("outbound order [%s] không có picking", params.SalesOrderNumber)
	}
	if picking.StatusId != 1 {
		return fmt.Errorf("picking [%s] không ở trạng thái được pick", picking.PickingNumber)
	}

	defaultLocation := DEFAULT_LOCATION
	allowPickBinLocations, err := u.binLocationSv.GetList(ctx, &binLocationModel.GetRequest{
		WarehouseId:      outbound.WarehouseId,
		IsAllowPickOrder: "1",
	})
	if err != nil {
		return err
	}
	if len(allowPickBinLocations) == 0 {
		newLoc, err := u.binLocationSv.Create(ctx, &binLocationModel.BinLocation{
			WarehouseId:          outbound.WarehouseId,
			Code:                 defaultLocation,
			Zone:                 defaultLocation[:2],
			Area:                 defaultLocation[3:5],
			Aisle:                defaultLocation[6:8],
			Rack:                 defaultLocation[9:11],
			Shelf:                defaultLocation[12:14],
			Bin:                  defaultLocation[15:17],
			AllowPicklistedOrder: 1,
			AllowPickOrder:       1,
		})
		if err != nil {
			return err
		}
		allowPickBinLocations = append(allowPickBinLocations, newLoc)
		// return fmt.Errorf("warehouse [%d] không có bin location di pick", outbound.WarehouseId)
	}
	if outbound.OutboundOrderType == "ORDER" {
		defaultLocation = allowPickBinLocations[0].Code
		for _, loc := range allowPickBinLocations {
			if strings.Contains(loc.Code, "AP") {
				defaultLocation = loc.Code
				break
			}
		}
	}

	pickingItems, err := u.pickingItemSv.GetList(ctx, &pickingItemModel.GetRequest{PickingId: picking.PickingId})
	if err != nil {
		return err
	}
	if len(pickingItems) == 0 {
		return fmt.Errorf("outbound [%s] không có hàng để pick", picking.PickingNumber)
	}

	updateLocationPickingItemIds := make([]int64, 0)
	inventoryIds := make([]int64, 0)
	for _, item := range pickingItems {
		if outbound.OutboundOrderType == "INTERNAL_TRANSFER" {
			if !checkSuitableLocationForIt(item.LocationDescription) {
				updateLocationPickingItemIds = append(updateLocationPickingItemIds, item.PickingItemId)
			}
		} else if outbound.OutboundOrderType == "ORDER" {
			if !checkSuitableLocationForOrder(item.LocationDescription, allowPickBinLocations) {
				updateLocationPickingItemIds = append(updateLocationPickingItemIds, item.PickingItemId)
			}
		} else {
			return fmt.Errorf("loại outbound order [%s] không hợp lệ: %s ", params.SalesOrderNumber, outbound.OutboundOrderType)
		}
		if item.StatusId == 1 {
			inventoryIds = append(inventoryIds, item.InventoryId)
		}
	}
	if len(inventoryIds) == 0 {
		return fmt.Errorf("outbound [%s] không có item phù hợp để pick", picking.PickingNumber)
	}

	inventories, err := u.invenorySv.GetList(ctx, outbound.WarehouseId, &invenoryModel.GetRequest{InventoryIds: inventoryIds})
	if err != nil {
		return err
	}
	for _, pi := range updateLocationPickingItemIds {
		if _, err := u.pickingItemSv.Update(ctx, &pickingItemModel.PickingItem{
			PickingItemId:       pi,
			LocationDescription: defaultLocation,
		}); err != nil {
			return err
		}
	}

	// update

	if time.Until(outbound.CreatedAt).Abs().Minutes() < 30 {
		if _, err := u.outboundOrderSv.Update(ctx, &outboundModel.OutboundOrder{
			OutboundOrderId: outbound.OutboundOrderId,
			CreatedAt:       outbound.CreatedAt.Add(-30 * time.Minute),
		}); err != nil {
			return err
		}
	}

	for _, inv := range inventories {
		isUpdate := false
		updatedInventory := &invenoryModel.Inventory{
			InventoryId: inv.InventoryId,
		}
		if outbound.OutboundOrderType == "INTERNAL_TRANSFER" {
			if !checkSuitableLocationForIt(inv.LocationDescription) {
				isUpdate = true
				updatedInventory.LocationDescription = DEFAULT_LOCATION
			}
		} else if outbound.OutboundOrderType == "ORDER" {
			if !checkSuitableLocationForOrder(inv.LocationDescription, allowPickBinLocations) {
				isUpdate = true
				updatedInventory.LocationDescription = defaultLocation
			}
		}

		if inv.StatusId != 7 {
			isUpdate = true
			updatedInventory.StatusId = 7
		}

		if inv.AreaId != 22 {
			isUpdate = true
			updatedInventory.AreaId = 22
		}

		if isUpdate {
			if _, err := u.invenorySv.Update(ctx, outbound.WarehouseId, updatedInventory); err != nil {
				return err
			}
		}
	}

	return nil
}

func checkSuitableLocationForIt(location string) bool {
	if strings.Contains(location, "AP") {
		return true
	}
	if strings.Contains(location, "PO") {
		return true
	}
	if strings.Contains(location, "PG") {
		return true
	}
	if strings.Contains(location, "IT") {
		return true
	}
	if strings.Contains(location, "ST") {
		return true
	}
	if strings.Contains(location, "NG") {
		return true
	}

	return false
}
func checkSuitableLocationForOrder(location string, listSuitable []*binLocationModel.BinLocation) bool {
	for _, item := range listSuitable {
		if item.Code == location {
			return true
		}
	}
	return false
}
