package constants

const (
	OUTBOUND_ORDER_STATUS_PICKLISTED int64 = iota + 1 // get from table wms_status
	OUTBOUND_ORDER_STATUS_CANCELED
	OUTBOUND_ORDER_STATUS_PENDING
	OUTBOUND_ORDER_STATUS_PICKED
	OUTBOUND_ORDER_STATUS_PACKED
	OUTBOUND_ORDER_STATUS_SHIPPED
	OUTBOUND_ORDER_STATUS_DELIVERY
	OUTBOUND_ORDER_STATUS_DELIVERED
	OUTBOUND_ORDER_STATUS_RETURNED
	OUTBOUND_ORDER_STATUS_COMPLETE_REPUSH
)

const (
	OUTBOUND_ORDER_TYPE_RECEIPT           string = "RECEIPT"
	OUTBOUND_ORDER_TYPE_ORDER             string = "ORDER"
	OUTBOUND_ORDER_TYPE_INTERNAL_TRANSFER string = "INTERNAL_TRANSFER"
	OUTBOUND_ORDER_TYPE_ADJUSTMENT        string = "ADJUSTMENT"
	OUTBOUND_ORDER_TYPE_VENDOR_RETURN     string = "VENDOR_RETURN"
)

const (
	OUTBOUND_ORDER_CONFIG_DESIGNATE_DATE         = 2097152
	OUTBOUND_ORDER_CONFIG_NOT_ENOUGH_QTY         = 1     // Not enough quantity
	OUTBOUND_ORDER_CONFIG_NOT_VERIFIED           = 2     // Đơn chưa được xác nhận
	OUTBOUND_ORDER_CONFIG_HAS_REASON_DELAY       = 4     // Đơn bị delay
	OUTBOUND_ORDER_CONFIG_HAS_SALES_OFF          = 8     // Đơn vào ngày sale off
	OUTBOUND_ORDER_CONFIG_HAS_PENDING            = 16    // Đơn tạm khóa
	OUTBOUND_ORDER_CONFIG_HAS_FORCE_PICKING      = 32    // Đơn cho đi pick khi thiếu uid
	OUTBOUND_ORDER_CONFIG_REQUIRED_EXPORT_INSIDE = 64    // Đơn yêu cầu xuất ở inside
	OUTBOUND_ORDER_CONFIG_WRONG_COMBO            = 128   // Đơn chứa combo empty
	OUTBOUND_ORDER_CONFIG_HAS_SALES_OFF_COMBO    = 256   // Đơn vào ngày sale off combo
	OUTBOUND_ORDER_CONFIG_WAITING_REPICK         = 512   // Đơn waiting repick
	OUTBOUND_ORDER_CONFIG_NOT_PICKLISTED_UID     = 1024  // Đơn không picklisted uid trước
	OUTBOUND_ORDER_CONFIG_IT_TESTER              = 2048  // Đơn không picklisted uid trước
	OUTBOUND_ORDER_CONFIG_WAITING_SHOP_CONFIRM   = 4096  // Đơn chờ shop xác nhận (thiếu hàng)
	OUTBOUND_ORDER_CONFIG_WAITING_CS_CONFIRM     = 8192  // Đơn chờ CS xác nhận (thiếu hàng)
	OUTBOUND_ORDER_CONFIG_CS_PROCESSING          = 16384 // Đơn CS đang xử lý (thiếu hàng)
	OUTBOUND_ORDER_CONFIG_ALLOW_MISSING_PICKUP   = 32768 // CS cho phép pick thiếu
	OUTBOUND_ORDER_CONFIG_COMPLETED              = 65536 // Đơn Hoàn thành xử lý (thiếu hàng)
)

var OUTBOUND_ORDER_CONFIG_MAP_NAME = map[int]string{
	OUTBOUND_ORDER_CONFIG_NOT_ENOUGH_QTY:         "Đơn không đủ hàng đi pick",
	OUTBOUND_ORDER_CONFIG_NOT_VERIFIED:           "Đơn chưa được xác nhận",
	OUTBOUND_ORDER_CONFIG_HAS_REASON_DELAY:       "Đơn bị delay",
	OUTBOUND_ORDER_CONFIG_HAS_SALES_OFF:          "Đơn vào ngày sale off",
	OUTBOUND_ORDER_CONFIG_HAS_PENDING:            "Đơn tạm khóa",
	OUTBOUND_ORDER_CONFIG_HAS_FORCE_PICKING:      "Đơn cho đi pick khi thiếu uid",
	OUTBOUND_ORDER_CONFIG_REQUIRED_EXPORT_INSIDE: "Đơn yêu cầu xuất ở inside",
	OUTBOUND_ORDER_CONFIG_WRONG_COMBO:            "Đơn chứa combo empty",
	OUTBOUND_ORDER_CONFIG_HAS_SALES_OFF_COMBO:    "Đơn vào ngày sale off combo",
	OUTBOUND_ORDER_CONFIG_WAITING_REPICK:         "Đơn waiting repick",
	OUTBOUND_ORDER_CONFIG_NOT_PICKLISTED_UID:     "Đơn không picklisted uid trước",
	OUTBOUND_ORDER_CONFIG_IT_TESTER:              "Đơn không picklisted uid trước",
	OUTBOUND_ORDER_CONFIG_WAITING_SHOP_CONFIRM:   "Đơn chờ shop xác nhận (thiếu hàng)",
	OUTBOUND_ORDER_CONFIG_WAITING_CS_CONFIRM:     "Đơn chờ CS xác nhận (thiếu hàng)",
	OUTBOUND_ORDER_CONFIG_CS_PROCESSING:          "Đơn CS đang xử lý (thiếu hàng)",
	OUTBOUND_ORDER_CONFIG_ALLOW_MISSING_PICKUP:   "CS cho phép pick thiếu",
	OUTBOUND_ORDER_CONFIG_COMPLETED:              "Đơn Hoàn thành xử lý (thiếu hàng)",
}
