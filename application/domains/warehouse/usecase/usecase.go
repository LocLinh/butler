package usecase

import (
	initServices "butler/application/domains/services/init"
	whModels "butler/application/domains/services/warehouse/models"
	whSv "butler/application/domains/services/warehouse/service"
	"butler/application/domains/warehouse/models"
	"context"
	"fmt"
	"strconv"
	"strings"
)

type usecase struct {
	whSv whSv.IService
}

func InitUseCase(
	services *initServices.Services,
) IUseCase {
	return &usecase{
		whSv: services.WarehouseService,
	}
}

const LOCATION_ID_555 = 12
const LOCATION_ID_29_HOANG_VIET = 382

func (u *usecase) ShowWarehouse(ctx context.Context, params *models.ShowWarehouseRequest) error {
	if params.WarehouseName == "" {
		return fmt.Errorf("warehouse name is required")
	}
	warehouseNameKeyword := strings.ToUpper(params.WarehouseName)

	suggestedWarehouses, err := u.whSv.GetList(ctx, &whModels.GetRequest{
		WarehouseNameSimilar: warehouseNameKeyword,
	})
	if err != nil {
		return err
	}
	if len(suggestedWarehouses) == 0 {
		return fmt.Errorf("không có kho nào có tên giống [%s]", warehouseNameKeyword)
	}
	if len(suggestedWarehouses) > 1 {
		var warehouseNames []string
		for _, wh := range suggestedWarehouses {
			warehouseNames = append(warehouseNames, wh.WarehouseName)
		}
		return fmt.Errorf("có nhiều kho có tên giống [%s], vui lòng nhập đúng tên: \n - %s", params.WarehouseName, strings.Join(warehouseNames, "\n- "))
	}
	warehouse := suggestedWarehouses[0]
	if warehouse.LocationId == LOCATION_ID_29_HOANG_VIET {
		return fmt.Errorf("kho [%s] đã có thể đi pick ở vị trí kho 29 hoang viet", warehouse.WarehouseName)
	}

	if _, err := u.whSv.Update(ctx, &whModels.Warehouse{
		WarehouseId: warehouse.WarehouseId,
		LocationId:  LOCATION_ID_29_HOANG_VIET,
		Description: fmt.Sprintf("location-%d", warehouse.LocationId),
	}); err != nil {
		return err
	}

	return nil
}

func (u *usecase) ResetShowWarehouse(ctx context.Context) error {
	warehouses, err := u.whSv.GetList(ctx, &whModels.GetRequest{})
	if err != nil {
		return err
	}
	for _, warehouse := range warehouses {
		if strings.Contains(warehouse.Description, "location-") {
			subDesc := strings.Split(warehouse.Description, "-")
			if len(subDesc) > 1 {
				locationIdStr := subDesc[1]
				location, _ := strconv.ParseInt(locationIdStr, 10, 64)
				if location == 0 {
					continue
				}
				if err := u.whSv.UpdateWithMap(ctx, warehouse.WarehouseId, map[string]any{
					"location_id": location,
					"description": "",
				}); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
