package handler

import (
	"butler/application/domains/pick_pack/models"
	"butler/application/domains/pick_pack/usecase"
	initServices "butler/application/domains/services/init"
	"butler/application/lib"
	"butler/config"
	"context"
	"errors"
	"regexp"
	"time"

	"strings"

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

func (h Handler) PickPackKafka(s *discordgo.Session, m *discordgo.MessageCreate) error {
	reg := regexp.MustCompile(`[0-9]+`)
	orderCode := strings.TrimSpace(reg.FindString(m.Content))

	if orderCode == "" {
		return errors.New("mã đơn không hợp lệ")
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	err := h.usecase.PickPackKafka(ctx, &models.AutoPickPackRequest{
		SalesOrderNumber: orderCode,
	})
	if err != nil {
		logrus.Errorf("Failed pickpack kafka: %v", err)
		return err
	}

	_, err = s.ChannelMessageSend(m.ChannelID, "DONE: Run PICK PACK")
	if err != nil {
		logrus.Errorf("Failed to send message: %v", err)
		return err
	}

	return nil
}

func (h Handler) SetOutboundOrderVoucherType(s *discordgo.Session, m *discordgo.MessageCreate) error {
	reg := regexp.MustCompile(`[0-9]+`)
	SalesOrderNumber := strings.TrimSpace(reg.FindString(m.Content))
	if SalesOrderNumber == "" {
		return errors.New("SalesOrderNumber không hợp lệ")
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	err := h.usecase.SetOutboundOrderVoucherType(ctx, &models.SetOutboundOrderVoucherTypeRequest{
		SalesOrderNumber: SalesOrderNumber,
		VoucherType:      1,
	})
	if err != nil {
		logrus.Errorf("Failed to set outbound order vouchertype: %v", err)
		return err
	}
	_, err = s.ChannelMessageSend(m.ChannelID, "DONE: Set Outbound Order Voucher Type = 1")
	if err != nil {
		logrus.Errorf("Failed to send message: %v", err)
		return err
	}

	return nil
}
