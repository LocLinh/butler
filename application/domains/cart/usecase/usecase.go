package usecase

import (
	"butler/application/domains/cart/models"
	cartMappingModels "butler/application/domains/services/bin_location_cart_mapping/models"
	cartMappingSv "butler/application/domains/services/bin_location_cart_mapping/service"
	cartModels "butler/application/domains/services/cart/models"
	cartSv "butler/application/domains/services/cart/service"
	initServices "butler/application/domains/services/init"
	packingModels "butler/application/domains/services/packing/models"
	packingSv "butler/application/domains/services/packing/service"
	pgModels "butler/application/domains/services/picking_group/models"
	pgSv "butler/application/domains/services/picking_group/service"
	userModels "butler/application/domains/services/user/models"
	userSv "butler/application/domains/services/user/service"
	"butler/constants"
	"context"
	"fmt"
)

type usecase struct {
	cartSv         cartSv.IService
	cartMappingSv  cartMappingSv.IService
	pickingGroupSv pgSv.IService
	packingSv      packingSv.IService
	userSv         userSv.IService
}

func InitUseCase(
	services *initServices.Services,
) IUseCase {
	return &usecase{
		cartSv:         services.CartService,
		cartMappingSv:  services.CartMappingService,
		pickingGroupSv: services.PickingGroupService,
		packingSv:      services.PackingService,
		userSv:         services.UserService,
	}
}

func (u *usecase) ResetCart(ctx context.Context, params *models.ResetCartRequest) error {
	cart, err := u.cartSv.GetOne(ctx, &cartModels.GetRequest{CartCode: params.CartCode})
	if err != nil {
		return err
	}
	if cart == nil || cart.CartId == 0 {
		return fmt.Errorf("cart with code [%s] not found", params.CartCode)
	}
	cartMapping, err := u.cartMappingSv.GetOne(ctx, &cartMappingModels.GetRequest{
		CartCode: cart.CartCode,
		UsedBy:   cart.UpdatedBy,
	})
	if err != nil {
		return err
	}

	pickingGroups, err := u.pickingGroupSv.GetList(ctx, &pgModels.GetRequest{
		CartCode:  cart.CartCode,
		StatusIds: []int64{constants.PICKING_GROUP_STATUS_NEW, constants.PICKING_GROUP_STATUS_PICKING}},
	)
	if err != nil {
		return err
	}

	packings, err := u.packingSv.GetList(ctx, &packingModels.GetRequest{
		CartCode: cart.CartCode,
		StatusIds: []int64{
			constants.PACKING_STATUS_OPEN, constants.PACKING_STATUS_PACKING,
		},
	})
	if err != nil {
		return err
	}

	_, err = u.cartSv.Update(ctx, &cartModels.Cart{
		CartId: cart.CartId,
		Status: constants.CART_STATUS_AVAILABLE,
	})
	if err != nil {
		return err
	}
	if cartMapping != nil {
		_, err = u.cartMappingSv.Update(ctx, &cartMappingModels.BinLocationCartMapping{
			Id:     cartMapping.Id,
			Status: constants.BIN_LOCATION_CART_MAPPING_STATUS_USING,
		})
		if err != nil {
			return err
		}
	} else {
		_, err := u.cartMappingSv.Create(ctx, &cartMappingModels.BinLocationCartMapping{
			CartCode: cart.CartCode,
			UsedBy:   cart.UpdatedBy,
			Status:   constants.BIN_LOCATION_CART_MAPPING_STATUS_USING,
		})
		if err != nil {
			return err
		}
	}

	for _, pg := range pickingGroups {
		_, err := u.pickingGroupSv.Update(ctx, &pgModels.PickingGroup{
			PickingGroupId: pg.PickingGroupId,
			Status:         constants.PICKING_GROUP_STATUS_CANCELED,
		})
		if err != nil {
			return err
		}
	}
	for _, packing := range packings {
		_, err := u.packingSv.Update(ctx, &packingModels.Packing{
			PackingId: packing.PackingId,
			StatusId:  constants.PACKING_STATUS_CANCELED,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *usecase) ResetCartByUserId(ctx context.Context, param *models.ResetCartByUserIdRequest) (string, error) {
	user, err := u.userSv.GetOne(ctx, &userModels.GetRequest{UserId: param.UserId})
	if err != nil {
		return "", err
	}
	if user == nil || user.Id == 0 {
		return "", fmt.Errorf("user with id [%d] not found", param.UserId)
	}
	userId := user.Id

	cart, err := u.cartSv.GetOne(ctx, &cartModels.GetRequest{UpdatedBy: userId})
	if err != nil {
		return "", err
	}
	if cart == nil || cart.CartId == 0 {
		return "", fmt.Errorf("cart using by user_id [%d] not found", userId)
	}
	cartMapping, err := u.cartMappingSv.GetOne(ctx, &cartMappingModels.GetRequest{
		CartCode: cart.CartCode,
	})
	if err != nil {
		return "", err
	}

	pickingGroups, err := u.pickingGroupSv.GetList(ctx, &pgModels.GetRequest{
		CartCode:  cart.CartCode,
		StatusIds: []int64{constants.PICKING_GROUP_STATUS_NEW, constants.PICKING_GROUP_STATUS_PICKING}},
	)
	if err != nil {
		return "", err
	}

	packings, err := u.packingSv.GetList(ctx, &packingModels.GetRequest{
		CartCode: cart.CartCode,
		StatusIds: []int64{
			constants.PACKING_STATUS_OPEN, constants.PACKING_STATUS_PACKING,
		},
	})
	if err != nil {
		return "", err
	}

	_, err = u.cartSv.Update(ctx, &cartModels.Cart{
		CartId: cart.CartId,
		Status: constants.CART_STATUS_AVAILABLE,
	})
	if err != nil {
		return "", err
	}
	if cartMapping != nil {
		_, err = u.cartMappingSv.Update(ctx, &cartMappingModels.BinLocationCartMapping{
			Id:     cartMapping.Id,
			Status: constants.BIN_LOCATION_CART_MAPPING_STATUS_USING,
		})
		if err != nil {
			return "", err
		}
	} else {
		_, err := u.cartMappingSv.Create(ctx, &cartMappingModels.BinLocationCartMapping{
			CartCode: cart.CartCode,
			UsedBy:   cart.UpdatedBy,
			Status:   constants.BIN_LOCATION_CART_MAPPING_STATUS_USING,
		})
		if err != nil {
			return "", err
		}
	}
	for _, pg := range pickingGroups {
		_, err := u.pickingGroupSv.Update(ctx, &pgModels.PickingGroup{
			PickingGroupId: pg.PickingGroupId,
			Status:         constants.PICKING_GROUP_STATUS_CANCELED,
		})
		if err != nil {
			return "", err
		}
	}
	for _, packing := range packings {
		_, err := u.packingSv.Update(ctx, &packingModels.Packing{
			PackingId: packing.PackingId,
			StatusId:  constants.PACKING_STATUS_CANCELED,
		})
		if err != nil {
			return "", err
		}
	}

	return cart.CartCode, nil
}

func (u *usecase) ResetCartByEmail(ctx context.Context, param *models.ResetCartByEmailRequest) (string, error) {
	user, err := u.userSv.GetOne(ctx, &userModels.GetRequest{Email: param.Email})
	if err != nil {
		return "", err
	}
	if user == nil || user.Id == 0 {
		return "", fmt.Errorf("user with email [%s] not found", param.Email)
	}
	userId := user.Id

	carts, err := u.cartSv.GetList(ctx, &cartModels.GetRequest{UpdatedBy: userId, Statuses: constants.CART_STATUS_NOT_AVAILABLE})
	if err != nil {
		return "", err
	}
	if len(carts) == 0 {
		return "", fmt.Errorf("cart using by user_id [%d] not found", userId)
	}

	cartCodeSring := ""
	for _, cart := range carts {
		cartCodeSring += cart.CartCode + ", "
		_, err = u.cartSv.Update(ctx, &cartModels.Cart{
			CartId: cart.CartId,
			Status: constants.CART_STATUS_AVAILABLE,
		})
		if err != nil {
			return "", err
		}
		cartMapping, err := u.cartMappingSv.GetOne(ctx, &cartMappingModels.GetRequest{CartCode: cart.CartCode, UsedBy: cart.UpdatedBy, Status: constants.BIN_LOCATION_CART_MAPPING_STATUS_FINISH})
		if err != nil {
			return "", err
		}
		if cartMapping != nil {
			err := u.cartMappingSv.Delete(ctx, cartMapping.Id)
			if err != nil {
				return "", err
			}
		}
		// else {
		// 	_, err := u.cartMappingSv.Create(ctx, &cartMappingModels.BinLocationCartMapping{
		// 		CartCode: cart.CartCode,
		// 		UsedBy:   cart.UpdatedBy,
		// 		Status:   constants.BIN_LOCATION_CART_MAPPING_STATUS_FINISH,
		// 	})
		// 	if err != nil {
		// 		return "", err
		// 	}
		// }
		pickingGroups, err := u.pickingGroupSv.GetList(ctx, &pgModels.GetRequest{
			CartCode:  cart.CartCode,
			StatusIds: []int64{constants.PICKING_GROUP_STATUS_NEW, constants.PICKING_GROUP_STATUS_PICKING}},
		)
		if err != nil {
			return "", err
		}

		packings, err := u.packingSv.GetList(ctx, &packingModels.GetRequest{
			CartCode: cart.CartCode,
			StatusIds: []int64{
				constants.PACKING_STATUS_OPEN, constants.PACKING_STATUS_PACKING,
			},
		})
		if err != nil {
			return "", err
		}

		for _, pg := range pickingGroups {
			_, err := u.pickingGroupSv.Update(ctx, &pgModels.PickingGroup{
				PickingGroupId: pg.PickingGroupId,
				Status:         constants.PICKING_GROUP_STATUS_CANCELED,
			})
			if err != nil {
				return "", err
			}
		}
		for _, packing := range packings {
			_, err := u.packingSv.Update(ctx, &packingModels.Packing{
				PackingId: packing.PackingId,
				StatusId:  constants.PACKING_STATUS_CANCELED,
			})
			if err != nil {
				return "", err
			}
		}
	}

	return cartCodeSring, nil
}
