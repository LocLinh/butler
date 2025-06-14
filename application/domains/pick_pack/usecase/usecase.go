package usecase

import (
	"butler/application/domains/pick_pack/models"
	initServices "butler/application/domains/services/init"
	outboundOrderSv "butler/application/domains/services/outbound_order/service"
	outboundOrderExtendSv "butler/application/domains/services/outbound_order_extend/service"
	"butler/application/lib"
	"butler/config"
	"butler/constants"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"
)

type usecase struct {
	lib                   *lib.Lib
	cfg                   *config.Config
	outboundOrderSv       outboundOrderSv.IService
	outboundOrderExtendSv outboundOrderExtendSv.IService
}

func InitUseCase(
	lib *lib.Lib,
	cfg *config.Config,
	services *initServices.Services,
) IUseCase {
	return &usecase{
		lib:                   lib,
		cfg:                   cfg,
		outboundOrderSv:       services.OutboundOrderService,
		outboundOrderExtendSv: services.OutboundOrderExtendService,
	}
}

func (u *usecase) AutoPickPack(ctx context.Context, params models.AutoPickPackRequest) (string, error) {
	// Login
	login, err := u.Login(ctx, &params.LoginRequest)
	if err != nil {
		return "", err
	}

	tokenWms := login.LoginWmsResponse.Token
	emailWms := login.LoginWmsResponse.User.Email
	userId := login.LoginWmsResponse.User.UserId
	tokenDiscord := login.LoginDiscordResponse.Token

	// Run newman json
	result, err := u.runNewman(ctx, params.SalesOrderNumber, params.ShippingUnitId, emailWms, int64(userId), tokenWms, tokenDiscord)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (u *usecase) Login(ctx context.Context, params *models.LoginRequest) (*models.LoginResponse, error) {
	var wms *models.LoginWmsResponse
	var discord *models.LoginDiscordResponse

	wmsDataStr, err := u.lib.Rdb.Get(ctx, fmt.Sprintf("%s:%s", constants.WMS_DATA, params.LoginWmsRequest.EmailWms)).Result()
	if err != nil && err.Error() != "redis: nil" {
		return nil, err
	}

	var wmsData map[string]string
	if wmsDataStr != "" {
		if err := json.Unmarshal([]byte(wmsDataStr), &wmsData); err != nil {
			return nil, err
		}
	}

	if wmsData == nil || wmsData["token"] == "" {
		wms, err = u.loginWms(ctx, params.LoginWmsRequest.EmailWms, params.LoginWmsRequest.PasswordWms)
		if err != nil {
			return nil, err
		}

		wmsData = map[string]string{
			"token":  wms.Token,
			"email":  wms.User.Email,
			"userId": strconv.FormatInt(int64(wms.User.UserId), 10),
		}

		wmsDataBytes, err := json.Marshal(wmsData)
		if err != nil {
			return nil, err
		}

		if err := u.lib.Rdb.Set(ctx, fmt.Sprintf("%s:%s", constants.WMS_DATA, wms.User.Email),
			string(wmsDataBytes), 8*time.Hour).Err(); err != nil {
			return nil, err
		}
	} else {
		userId, _ := strconv.Atoi(wmsData["userId"])
		wms = &models.LoginWmsResponse{
			Token: wmsData["token"],
			User: struct {
				UserId   int    `json:"user_id"`
				LastName string `json:"last_name"`
				Email    string `json:"email"`
				Status   string `json:"status"`
			}{
				Email:  wmsData["email"],
				UserId: userId,
			},
		}
	}

	tokenDiscord, err := u.lib.Rdb.Get(ctx, constants.TOKEN_DISCORD).Result()
	if err != nil && err.Error() != "redis: nil" {
		return nil, err
	}

	if tokenDiscord == "" {
		discord, err = u.loginDiscord(ctx, params.LoginDiscordRequest.LoginDiscord, params.LoginDiscordRequest.PasswordDiscord)
		if err != nil {
			return nil, err
		}
		tokenDiscord = discord.Token

		if err := u.lib.Rdb.Set(ctx, constants.TOKEN_DISCORD, tokenDiscord, 24*time.Hour).Err(); err != nil {
			return nil, err
		}
	} else {
		discord = &models.LoginDiscordResponse{
			Token: tokenDiscord,
		}
	}

	if wms.Token == "" || discord.Token == "" {
		return nil, errors.New("token wms or token discord is empty")
	}

	return &models.LoginResponse{
		LoginWmsResponse:     *wms,
		LoginDiscordResponse: *discord,
	}, nil
}

func (u *usecase) runNewman(ctx context.Context, shipmentNumber string, shippingUnitId int64, emailWms string, userId int64, tokenWms string, tokenDiscord string) (string, error) {
	cmd := exec.Command("newman",
		"run",
		"collection1.json",
		"--env-var", fmt.Sprintf("token_discord=%s", tokenDiscord),
		"--env-var", fmt.Sprintf("token=%s", tokenWms),
		"--env-var", fmt.Sprintf("email=%s", emailWms),
		"--env-var", fmt.Sprintf("user_id=%d", userId),
		"--env-var", fmt.Sprintf("shipment_number=%s", shipmentNumber),
		"--env-var", fmt.Sprintf("shipping_unit_id=%d", shippingUnitId),
	)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run newman
	err := cmd.Run()
	if err != nil {
		log.Printf("Lỗi stderr: %s", stderr.String())
		return "", fmt.Errorf("error running newman: %v\nStderr: %s", err, stderr.String())
	}

	return fmt.Sprintf("Stdout: %s\nStderr: %s", stdout.String(), stderr.String()), nil
}
