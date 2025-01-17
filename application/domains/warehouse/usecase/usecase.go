package usecase

import (
	initServices "butler/application/domains/services/init"
	whModels "butler/application/domains/services/warehouse/models"
	whSv "butler/application/domains/services/warehouse/service"
	"butler/application/domains/warehouse/models"
	"butler/application/lib"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type usecase struct {
	lib  *lib.Lib
	whSv whSv.IService
}

func InitUseCase(
	lib *lib.Lib,
	services *initServices.Services,
) IUseCase {
	return &usecase{
		lib:  lib,
		whSv: services.WarehouseService,
	}
}

const LOCATION_ID_555 = 12
const LOCATION_ID_29_HOANG_VIET = 382

func (u *usecase) ShowWarehouse(ctx context.Context, params *models.ShowWarehouseRequest) (err error) {
	var warehouse *whModels.Warehouse
	switch {
	case params.WarehouseName != "":
		warehouseNameKeyword := strings.ToUpper(params.WarehouseName)
		suggestedWarehouses, err := u.whSv.GetList(ctx, &whModels.GetRequest{
			WarehouseName: warehouseNameKeyword,
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
		if warehouse.LocationId == LOCATION_ID_555 {
			return fmt.Errorf("kho [%s] đã có thể đi pick ở vị trí kho 555 3/2", warehouse.WarehouseName)
		}
	case params.WarehouseId != 0:
		warehouse, err = u.whSv.GetOne(ctx, &whModels.GetRequest{
			WarehouseId: params.WarehouseId,
		})
		if err != nil {
			return err
		}
		if warehouse.WarehouseId == 0 {
			return fmt.Errorf("không tìm thấy kho")
		}
		if warehouse.LocationId == LOCATION_ID_555 {
			return fmt.Errorf("kho [%s] đã có thể đi pick ở vị trí kho 555 3/2", warehouse.WarehouseName)
		}
	default:
		return fmt.Errorf("Check lại lệnh, ví dụ: !showwarehouse <tên kho> hoặc !show_warehouse_by_id <warehouse_id>")
	}

	return u.updateLocationWarehouse(ctx, warehouse.WarehouseId, LOCATION_ID_555)

}

func (u *usecase) updateLocationWarehouse(ctx context.Context, warehouseId int64, locationId int64) error {
	if _, err := u.whSv.Update(ctx, &whModels.Warehouse{
		WarehouseId: warehouseId,
		LocationId:  locationId,
		Description: fmt.Sprintf("location-%d", locationId),
	}); err != nil {
		return err
	}
	return nil
}

func (u *usecase) ResetShowWarehouse(ctx context.Context, warehouseId int64) (string, error) {
	warehouses, err := u.whSv.GetList(ctx, &whModels.GetRequest{WarehouseId: warehouseId})
	if err != nil {
		return "", err
	}
	whResetedName := []string{}
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
					return "", err
				}
				whResetedName = append(whResetedName, fmt.Sprintf("%s,", warehouse.WarehouseName))
			}
		}
	}

	// reset redis
	currentIpAddr := "14.241.249.24"

	key := fmt.Sprintf("ipaddress:warehouse:%s", currentIpAddr)
	if value, err := u.lib.Rdb.Get(ctx, key).Result(); err == nil {
		if value == "" {
			return "", nil
		}

		if _, err := u.lib.Rdb.Del(ctx, key).Result(); err != nil {
			logrus.Errorf("Failed to delete key %s: %v", key, err)
		}
	}

	return strings.Join(whResetedName, "\n"), nil
}
