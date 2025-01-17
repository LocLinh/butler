package handler

import (
	"butler/application/domains/pick_pack/models"
	"butler/application/domains/pick_pack/usecase"
	initServices "butler/application/domains/services/init"
	"butler/application/lib"
	"butler/config"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	lib     *lib.Lib
	usecase usecase.IUseCase
}

func InitHandler(lib *lib.Lib, cfg *config.Config, services *initServices.Services) Handler {
	usecase := usecase.InitUseCase(lib, cfg, services)
	return Handler{
		lib:     lib,
		usecase: usecase,
	}
}

func (h Handler) ReadyPickPack(s *discordgo.Session, m *discordgo.MessageCreate) error {
	re := regexp.MustCompile(`\[([^\]]+)\]`)
	matches := re.FindAllStringSubmatch(m.Content, -1)

	if len(matches) < 4 {
		return fmt.Errorf("Thiếu tham số. Vui lòng sử dụng: !runpickpack [email][password][mã đơn][mã vận chuyển]")
	}

	// Lấy email và password từ tham số đầu vào
	emailWms := matches[0][1]                                      // email ở vị trí đầu tiên
	passwordWms := matches[1][1]                                   // password ở vị trí thứ hai
	shipmentNumber := matches[2][1]                                // mã đơn ở vị trí thứ ba
	shippingUnitId, err := strconv.ParseInt(matches[3][1], 10, 64) // mã vận chuyển ở vị trí thứ tư
	if err != nil {
		return fmt.Errorf("Mã vận chuyển không hợp lệ: %v", err)
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()
	requestParams := &models.AutoPickPackRequest{
		LoginRequest: models.LoginRequest{
			LoginWmsRequest: models.LoginWmsRequest{
				EmailWms:    emailWms,
				PasswordWms: passwordWms,
			},
			LoginDiscordRequest: models.LoginDiscordRequest{
				LoginDiscord:    "trieuld@hasaki.vn",
				PasswordDiscord: "123456aH@",
				Undelete:        false,
			},
		},
		SalesOrderNumber: shipmentNumber,
		ShippingUnitId:   shippingUnitId,
	}

	_, err = h.usecase.AutoPickPack(ctx, *requestParams)
	if err != nil {
		logrus.Errorf("Failed ready pickpack: %v", err)
		return err
	}

	// Chia nhỏ kết quả thành nhiều phần nếu quá dài
	// const maxLength = 1990 // Để lại một chút dư cho các ký tự định dạng
	// for i := 0; i < len(result); i += maxLength {
	// 	end := i + maxLength
	// 	if end > len(result) {
	// 		end = len(result)
	// 	}
	// 	chunk := result[i:end]
	// 	_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("```%s```", chunk))
	// 	if err != nil {
	// 		logrus.Errorf("Failed to send message: %v", err)
	// 		return err
	// 	}
	// }
	_, err = s.ChannelMessageSend(m.ChannelID, "DONE: Run PICK PACK")
	if err != nil {
		logrus.Errorf("Failed to send message: %v", err)
		return err
	}

	return nil
}
