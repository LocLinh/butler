package init

import (
	binLocationInit "butler/application/domains/services/bin_location/init"
	binLocationSv "butler/application/domains/services/bin_location/service"
	cartMappingInit "butler/application/domains/services/bin_location_cart_mapping/init"
	cartMappingSv "butler/application/domains/services/bin_location_cart_mapping/service"
	cartInit "butler/application/domains/services/cart/init"
	cartSv "butler/application/domains/services/cart/service"
	inventoryInit "butler/application/domains/services/inventory/init"
	inventorySv "butler/application/domains/services/inventory/service"
	outboundOrderInit "butler/application/domains/services/outbound_order/init"
	outboundOrderSv "butler/application/domains/services/outbound_order/service"
	packingInit "butler/application/domains/services/packing/init"
	packingSv "butler/application/domains/services/packing/service"
	pickingInit "butler/application/domains/services/picking/init"
	pickingSv "butler/application/domains/services/picking/service"
	pickingGroupInit "butler/application/domains/services/picking_group/init"
	pickingGroupSv "butler/application/domains/services/picking_group/service"
	pickingItemInit "butler/application/domains/services/picking_item/init"
	pickingItemSv "butler/application/domains/services/picking_item/service"
	promtAiSv "butler/application/domains/services/promt_ai/service"
	userInit "butler/application/domains/services/user/init"
	userSv "butler/application/domains/services/user/service"
	warehouseInit "butler/application/domains/services/warehouse/init"
	warehouseSv "butler/application/domains/services/warehouse/service"
	"butler/config"

	"github.com/google/generative-ai-go/genai"
	"gorm.io/gorm"
)

type Services struct {
	PromtAiSv            promtAiSv.IService
	CartService          cartSv.IService
	CartMappingService   cartMappingSv.IService
	PackingService       packingSv.IService
	PickingGroupService  pickingGroupSv.IService
	OutboundOrderService outboundOrderSv.IService
	PickingService       pickingSv.IService
	PickingItemService   pickingItemSv.IService
	InventoryService     inventorySv.IService
	BinLocationService   binLocationSv.IService
	WarehouseService     warehouseSv.IService
	UserService          userSv.IService
}

func InitService(cfg *config.Config, db *gorm.DB, genaiClient *genai.Client) *Services {
	initPromtAiSv := promtAiSv.InitService(cfg, genaiClient)
	cart := cartInit.NewInit(db, cfg)
	cartMapping := cartMappingInit.NewInit(db, cfg)
	packing := packingInit.NewInit(db, cfg)
	pickingGroup := pickingGroupInit.NewInit(db, cfg)
	outboundOrder := outboundOrderInit.NewInit(db, cfg)
	picking := pickingInit.NewInit(db, cfg)
	pickingItem := pickingItemInit.NewInit(db, cfg)
	inventory := inventoryInit.NewInit(db, cfg)
	binLocation := binLocationInit.NewInit(db, cfg)
	warehouse := warehouseInit.NewInit(db, cfg)
	user := userInit.NewInit(db, cfg)
	return &Services{
		PromtAiSv:            initPromtAiSv,
		CartService:          cart.Service,
		CartMappingService:   cartMapping.Service,
		PackingService:       packing.Service,
		PickingGroupService:  pickingGroup.Service,
		OutboundOrderService: outboundOrder.Service,
		PickingService:       picking.Service,
		PickingItemService:   pickingItem.Service,
		InventoryService:     inventory.Service,
		BinLocationService:   binLocation.Service,
		WarehouseService:     warehouse.Service,
		UserService:          user.Service,
	}
}
