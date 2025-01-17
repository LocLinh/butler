package constants

const (
	WMS_LOGIN                      = "/api/v1/auth/user/login"
	WMS_OUTBOUND_ORDER             = "/api/v1/wms/outbound/outbound-orders"
	WMS_OUTBOUND_ORDER_ITEM        = "/api/v1/wms/outbound/outbound-order-items/detail-by-outbound-order"
	WMS_CART                       = "api/v1/wms/outbound-orders/cart"
	WMS_CREATE_PICKING_GROUP       = "/api/v1/wms/pick-pack/orders/picking-group/create"
	WMS_ADD_ORDER_TO_PICKING_GROUP = "/api/v1/wms/pick-pack/orders/picking-group/validate-outbound"
	WMS_GET_MY_PICKING_GROUP       = "/api/v1/wms/outbound-orders/picking-group/group/me"
	WMS_START_PICKING_GROUP        = "/api/v1/wms/outbound-orders/picking-group/start"
	WMS_GET_PICKING_GROUP_BY_ID    = "/api/v1/wms/outbound-orders/picking-group/tracking/by-group/"
	WMS_PICKS                      = "/api/v1/wms/outbound-orders/picking-group/tracking/pick-many-v1"
	WMS_GET_CAMERA_CODE            = "/api/v1/wms/cameras"
	WMS_START_PICKING              = "/api/v1/wms/outbound-orders/packing/by-cart/v2"
	WMS_PACKED                     = "/api/v1/wms/outbound-orders/packing/complete"
	WMS_GET_RECEIPT                = "/api/v1/wms/outbound-orders/packing/receipt"
)

const (
	DISCORD_LOGIN               = "/api/v9/auth/login"
	DISCORD_RESET_CART_BY_EMAIL = "/api/v9/channels/1134696880668942447/messages"
	DISCORD_READY_PICK          = "/api/v9/channels/1134696880668942447/messages"
)
