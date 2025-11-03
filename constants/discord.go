package constants

const BOT_COMMAND_PREFIX = "!"

const (
	COMMAND_HELP                            = "help"
	COMMAND_REMIND                          = "remind"
	COMMAND_RESET_CART                      = "resetcart"
	COMMAND_RESET_CART_BY_USER_ID           = "reset_cart_by_user_id"
	COMMAND_RESET_CART_BY_EMAIL             = "reset_cart_by_email"
	COMMAND_READY_PICK                      = "readypick"
	COMMAND_SHOW_WAREHOUSE                  = "showwarehouse"
	COMMAND_SHOW_WAREHOUSE_BY_ID            = "show_warehouse_by_id"
	COMMAND_RESET_SHOW_WAREHOUSE            = "resetshowwarehouse"
	COMMAND_RESET_SHOW_WAREHOUSE_BY_ID      = "reset_show_warehouse_by_id"
	COMMAND_PICK                            = "pick"
	COMMAND_COUNT_KPI                       = "kpi"
	COMMAND_COUNT_PROD_KPI                  = "prod kpi"
	COMMAND_WH_CONFIG                       = "whcfg"
	COMMAND_PICK_PACK_KAFKA                 = "ppkafka"
	COMMAND_SET_VOUCHER_TYPE_OUTBOUND_ORDER = "set_voucher_type_outbound_order"
)

const USER_ID_REGEX = `<@(\d+)>`
