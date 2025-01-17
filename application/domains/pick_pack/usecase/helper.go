package usecase

import (
	"butler/application/domains/pick_pack/models"
	"butler/constants"
	"context"
	"encoding/json"
	"errors"

	"bitbucket.org/hasaki-tech/zeus/package/hrequest"
)

const NO_RETRY = 0
const NO_DELAY = 0
const TIMEOUT_20S = 20

func apiRetryCondition(status int, body []byte, errRequest error) bool {
	return false
}

func (u *usecase) loginWms(ctx context.Context, email, password string) (*models.LoginWmsResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	endpoint := u.cfg.ApiExternal.Wms.Url
	url := endpoint + constants.WMS_LOGIN

	_, body, err := hrequest.MakeRequest(headers, constants.METHOD_POST, url, &models.LoginWmsRequest{
		EmailWms:    email,
		PasswordWms: password,
	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
	if err != nil {
		return nil, err
	}
	response := &models.LoginWmsResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}
	if response.Message != "" {
		return nil, errors.New(response.Message)
	}
	return response, nil
}

func (u *usecase) loginDiscord(ctx context.Context, email, password string) (*models.LoginDiscordResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	endpoint := u.cfg.ApiExternal.Discord.Url
	url := endpoint + constants.DISCORD_LOGIN

	_, body, err := hrequest.MakeRequest(headers, constants.METHOD_POST, url, &models.LoginDiscordRequest{
		LoginDiscord:    email,
		PasswordDiscord: password,
		Undelete:        false,
	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
	if err != nil {
		return nil, err
	}
	response := &models.LoginDiscordResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}
	if response.Message != "" {
		return nil, errors.New(response.Message)
	}

	return response, nil

}

// func (u *usecase) resetCartByEmail(ctx context.Context, email string, token string) (*models.ResetCartByEmailResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Discord.Url
// 	url := endpoint + constants.DISCORD_RESET_CART_BY_EMAIL

// 	content := fmt.Sprintf("!reset_cart_by_email %s", email)
// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_POST, url, &models.ResetCartByEmailRequest{Content: content},
// 		10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.ResetCartByEmailResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	if response.Message != "" {
// 		return nil, errors.New(response.Message)
// 	}
// 	return response, nil
// }

// func (u *usecase) getOutboundOrder(ctx context.Context, SalesOrderNumber, token string) (*models.GetOrderResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_OUTBOUND_ORDER

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_GET, url, &models.GetOrderRequest{
// 		SalesOrderNumbers: SalesOrderNumber,
// 		Page:              1,
// 		Size:              1,
// 		ConfigQuery:       1,
// 	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.GetOrderResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) getOutboundOrderItem(ctx context.Context, outboundOrderId int64, token string) (*models.GetOrderItemResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_OUTBOUND_ORDER_ITEM

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_GET, url, &models.GetOrderItemRequest{
// 		OutboundOrderId: outboundOrderId,
// 		Page:            1,
// 		Size:            1,
// 	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.GetOrderItemResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) getCart(ctx context.Context, warehouseId int64, token string) (*models.GetCartResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_CART

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_GET, url, &models.GetCartRequest{
// 		WarehouseIds: strconv.FormatInt(warehouseId, 10),
// 		Page:         1,
// 		Size:         1,
// 		StatusIds:    strconv.FormatInt(int64(1), 10),
// 	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.GetCartResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) createPickingGroup(ctx context.Context, warehouseId, shippingUnitId, userId int64, cartCode string, numGroup, numOrder int64, token string) (*models.CreatePickingGroupResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_CREATE_PICKING_GROUP

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_POST, url, &models.CreatePickingGroupRequest{
// 		WarehouseId:    warehouseId,
// 		ShippingUnitId: shippingUnitId,
// 		UserId:         userId,
// 		CartCode:       cartCode,
// 		NumGroup:       numGroup,
// 		NumOrder:       numOrder,
// 	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.CreatePickingGroupResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) addOrderToPickingGroup(ctx context.Context, pickingGroupId, warehouseId int64, keyword string, token string) (*models.AddOrderToPickingGroupResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_ADD_ORDER_TO_PICKING_GROUP

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_GET, url, &models.AddOrderToPickingGroupRequest{
// 		PickingGroupId: pickingGroupId,
// 		WarehouseId:    warehouseId,
// 		Keyword:        keyword,
// 	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.AddOrderToPickingGroupResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) getMyPickingGroup(ctx context.Context, token string) (*models.GetMyPickingGroupResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_GET_MY_PICKING_GROUP

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_GET, url, &models.GetMyPickingGroupRequest{},
// 		10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.GetMyPickingGroupResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) startPickingGroup(ctx context.Context, pickingGroupId, outboundOrderId, userId int64, token string) (*models.StartPickingGroupResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_START_PICKING_GROUP

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_POST, url, &models.StartPickingGroupRequest{
// 		GroupId:          pickingGroupId,
// 		OutboundOrderIds: strconv.FormatInt(outboundOrderId, 10),
// 		UserId:           userId,
// 	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.StartPickingGroupResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) getPickingGroupById(ctx context.Context, pickingGroupId int64, token string) (*models.GetPickingGroupByIdResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_GET_PICKING_GROUP_BY_ID + strconv.FormatInt(pickingGroupId, 10)

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_GET, url, &models.GetPickingGroupByIdRequest{}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.GetPickingGroupByIdResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) picks(ctx context.Context, params models.PicksRequest, token string) (*models.PicksResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_PICKS

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_POST, url, params, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.PicksResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) getCameraCode(ctx context.Context, warehouseId int64, token string) (*models.GetCameraCodeResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_GET_CAMERA_CODE

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_GET, url, &models.GetCameraCodeRequest{
// 		WarehouseId: warehouseId,
// 		Page:        1,
// 		Size:        20,
// 		StatusIds:   "1",
// 	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.GetCameraCodeResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) startPicking(ctx context.Context, cartCode, cameraCode string, token string) (*models.StartPickingResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_START_PICKING

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_GET, url, &models.StartPickingRequest{
// 		CartCode:   cartCode,
// 		CameraCode: cameraCode,
// 	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.StartPickingResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) packed(ctx context.Context, outboundOrderId int64, token string) (*models.PackedResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_PACKED

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_GET, url, &models.PackedRequest{
// 		BoxType:         "string",
// 		OutboundOrderId: outboundOrderId,
// 	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.PackedResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (u *usecase) getReceipt(ctx context.Context, outboundOrderId int64, token string) (*models.GetReceiptResponse, error) {
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer " + token,
// 	}

// 	endpoint := u.cfg.ApiExternal.Wms.Url
// 	url := endpoint + constants.WMS_GET_RECEIPT

// 	_, body, err := hrequest.MakeRequest(header, constants.METHOD_GET, url, &models.GetReceiptRequest{
// 		OutboundOrderId: outboundOrderId,
// 	}, 10, NO_RETRY, NO_DELAY, apiRetryCondition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &models.GetReceiptResponse{}
// 	if err := json.Unmarshal(body, response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }
