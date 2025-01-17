package models

type ResetCartByEmailRequest struct {
	Content string `json:"content"`
}

type ResetCartByEmailResponse struct {
	Message string `json:"message"`
}

type ReadyPickRequest struct {
	MobileNetworkType string `json:"mobile_network_type"`
	Content           string `json:"content"`
	Tts               bool   `json:"tts"`
	Flags             int    `json:"flags"`
}
type ReadyPickResponse struct {
	Message string `json:"message"`
}

type GetOrderRequest struct {
	SalesOrderNumbers string `json:"sales_order_numbers"`
	Page              int64  `json:"page"`
	Size              int64  `json:"size"`
	ConfigQuery       int64  `json:"config_query"`
}

type GetOrderResponse struct {
	Page    int64 `json:"page"`
	Size    int64 `json:"size"`
	Records []struct {
		OutboundOrderId     int64  `json:"outbound_order_id"`
		OutboundOrderNumber string `json:"outbound_order_number"`
	} `json:"records"`
}

/*  */
type GetOrderItemRequest struct {
	OutboundOrderId int64 `json:"outbound_order_id"`
	Page            int64 `json:"page"`
	Size            int64 `json:"size"`
}

type GetOrderItemResponse struct {
	Count   int64 `json:"count"`
	Records []struct {
		OutboundOrderItemId int64  `json:"outbound_order_item_id"`
		OutboundOrderId     int64  `json:"outbound_order_id"`
		StatusId            int64  `json:"status_id"`
		ProductId           int64  `json:"product_id"`
		ProductName         string `json:"product_name"`
		Sku                 string `json:"sku"`
		Quantity            int64  `json:"quantity"`
	} `json:"records"`
}

type GetCartRequest struct {
	WarehouseIds string `json:"warehouse_ids"`
	Page         int64  `json:"page"`
	Size         int64  `json:"size"`
	StatusIds    string `json:"status_ids"`
}

type GetCartResponse struct {
	Count   int64 `json:"count"`
	Page    int64 `json:"page"`
	Size    int64 `json:"size"`
	Records []struct {
		WarehouseId       int64  `json:"warehouse_id"`
		WarehouseName     string `json:"warehouse_name"`
		CartCode          string `json:"cart_code"`
		Status            int64  `json:"status"`
		UpdatedAt         string `json:"updated_at"`
		UpdatedBy         int64  `json:"updated_by"`
		UpdatedByUserName string `json:"updated_by_user_name"`
		CartId            int64  `json:"cart_id"`
		StatusName        string `json:"status_name"`
	} `json:"records"`
}

type CreatePickingGroupRequest struct {
	WarehouseId    int64  `json:"warehouse_id"`
	ShippingUnitId int64  `json:"shipping_unit_id"`
	UserId         int64  `json:"user_id"`
	CartCode       string `json:"cart_code"`
	NumGroup       int64  `json:"num_group"`
	NumOrder       int64  `json:"num_order"`
}

type CreatePickingGroupResponse struct {
	PickingGroupId int64 `json:"picking_group_id"`
}

type AddOrderToPickingGroupRequest struct {
	Keyword        string `json:"keyword"`
	WarehouseId    int64  `json:"warehouse_id"`
	PickingGroupId int64  `json:"picking_group_id"`
}

type AddOrderToPickingGroupResponse struct {
	CreatedAt           string `json:"created_at"`
	CreatedBy           int64  `json:"created_by"`
	ErrorMessage        string `json:"error_message"`
	OutboundOrderId     int64  `json:"outbound_order_id"`
	OutboundOrderNumber string `json:"outbound_order_number"`
	OutboundOrderType   string `json:"outbound_order_type"`
	SalesOrderId        string `json:"sales_order_id"`
	SalesOrderNumber    string `json:"sales_order_number"`
	StatusId            int64  `json:"status_id"`
}

/* */
type GetMyPickingGroupRequest struct {
}

type GetMyPickingGroupResponse struct {
	Count   int64 `json:"count"`
	Page    int64 `json:"page"`
	Records []struct {
		CartCode        string `json:"cart_code"`
		CartCodeRepick  string `json:"cart_code_repick"`
		CartCodeReturn  string `json:"cart_code_return"`
		CreatedAt       string `json:"created_at"`
		CreatedBy       int64  `json:"created_by"`
		CreatedUserName string `json:"created_user_name"`
		Description     string `json:"description"`
		GroupId         int64  `json:"group_id"`
	} `json:"records"`
}

/* */
type StartPickingGroupRequest struct {
	GroupId          int64  `json:"group_id"`
	OutboundOrderIds string `json:"outbound_order_ids"`
	UserId           int64  `json:"user_id"`
}

type StartPickingGroupResponse struct {
	Message string `json:"message"`
}

/* */
type GetPickingGroupByIdRequest struct {
	GroupId int64 `json:"group_id"`
}

type GetPickingGroupByIdResponse struct {
	Message        string `json:"message"`
	IsForceSuggest bool   `json:"is_force_suggest"`
	NotPickItems   []struct {
		Barcode          string `json:"barcode"`
		BrandId          int64  `json:"brand_id"`
		CartCode         string `json:"cart_code"`
		Expiration       string `json:"expiration"`
		GroupId          int64  `json:"group_id"`
		Location         string `json:"location"`
		NextPgCode       string `json:"next_pg_code"`
		PgCode           string `json:"pg_code"`
		PgTrackingId     int64  `json:"pg_tracking_id"`
		Priority         int64  `json:"priority"`
		ProductName      string `json:"product_name"`
		Quantity         int64  `json:"quantity"`
		QuantityNotFound int64  `json:"quantity_notfound"`
		QuantityPicked   int64  `json:"quantity_picked"`
		RequireExpDate   string `json:"require_exp_date"`
		SerialNumber     string `json:"serial_number"`
		Sku              string `json:"sku"`
		Status           int64  `json:"status"`
	} `json:"not_pick_items"`
	PickedItems []struct {
		Barcode          string `json:"barcode"`
		BrandId          int64  `json:"brand_id"`
		CartCode         string `json:"cart_code"`
		Expiration       string `json:"expiration"`
		GroupId          int64  `json:"group_id"`
		Location         string `json:"location"`
		NextPgCode       string `json:"next_pg_code"`
		PgCode           string `json:"pg_code"`
		PgTrackingId     int64  `json:"pg_tracking_id"`
		Priority         int64  `json:"priority"`
		ProductName      string `json:"product_name"`
		Quantity         int64  `json:"quantity"`
		QuantityNotFound int64  `json:"quantity_notfound"`
		QuantityPicked   int64  `json:"quantity_picked"`
		RequireExpDate   string `json:"require_exp_date"`
		SerialNumber     string `json:"serial_number"`
		Sku              string `json:"sku"`
		Status           int64  `json:"status"`
	} `json:"picked_items"`
}

/* */
type PicksRequest struct {
	Items []struct {
		ExpirationDate []struct {
			Exp string `json:"exp"`
			Qty int64  `json:"qty"`
		} `json:"expiration_date"`
		PgCode       string   `json:"pg_code"`
		Quantity     int64    `json:"quantity"`
		ReasonCode   string   `json:"reason_code"`
		SerialNumber []string `json:"serial_number"`
		TempCode     string   `json:"temp_code"`
	} `json:"items"`
}

type PicksResponse struct {
	Message string `json:"message"`
}

/* */
type GetCameraCodeRequest struct {
	WarehouseId int64  `json:"warehouse_id"`
	Page        int64  `json:"page"`
	Size        int64  `json:"size"`
	StatusIds   string `json:"status_ids"`
}

type GetCameraCodeResponse struct {
	Message string `json:"message"`
}

/* */
type StartPickingRequest struct {
	CartCode   string `json:"cart_code"`
	CameraCode string `json:"camera_code"`
}

type StartPickingResponse struct {
	Message string `json:"message"`
}

/* */
type PackedRequest struct {
	BoxType         string `json:"box_type"`
	OutboundOrderId int64  `json:"outbound_order_id"`
}

type PackedResponse struct {
	Message string `json:"message"`
}

/* */
type GetReceiptRequest struct {
	OutboundOrderId int64 `json:"outbound_order_id"`
}

type GetReceiptResponse struct {
	Message string `json:"message"`
}
