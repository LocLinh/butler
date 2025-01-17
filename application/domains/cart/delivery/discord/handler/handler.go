package handler

import (
	"butler/application/domains/cart/models"
	"butler/application/domains/cart/usecase"
	initServices "butler/application/domains/services/init"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	usecase usecase.IUseCase
}

func InitHandler(services *initServices.Services) Handler {
	usecase := usecase.InitUseCase(services)
	return Handler{
		usecase,
	}
}

func (h Handler) ResetCart(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// find cart code in message
	reg := regexp.MustCompile(`[0-9]+`)
	cartCode := reg.FindString(m.Content)
	logrus.Infof("Reset cart code: %s", cartCode)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	if err := h.usecase.ResetCart(ctx, &models.ResetCartRequest{
		CartCode: cartCode,
	}); err != nil {
		logrus.Errorf("Failed to reset cart: %v", err)
		return err
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Cart [%s] has been reset!", cartCode))
	return nil
}

func (h Handler) ResetCartByUserId(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// find user_id in message
	reg := regexp.MustCompile(`[0-9]+`)
	userIdStr := reg.FindString(m.Content)
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		logrus.Errorf("Failed to parse user ID: %v", err)
		return err
	}

	logrus.Infof("Reset cart mapping : %s", userIdStr)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	cartCode, err := h.usecase.ResetCartByUserId(ctx, &models.ResetCartByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		logrus.Errorf("Failed to reset cart mapping: %v", err)
		return err
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Cart [%s] - User [%s] has been reset!", cartCode, userIdStr))
	return nil
}

func (h Handler) ResetCartByEmail(s *discordgo.Session, m *discordgo.MessageCreate) error {
	parts := strings.Fields(m.Content)
	if len(parts) < 2 {
		return fmt.Errorf("Invalid command format. Usage: !reset_cart_email <username/email>")
	}
	input := parts[1]

	var email string
	emailReg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailReg.MatchString(input) {
		email = input
	} else {
		usernameReg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+$`)
		if usernameReg.MatchString(input) {
			email = input + "@hasaki.vn"
		} else {
			return fmt.Errorf("Invalid username/email format")
		}
	}

	logrus.Infof("Reset cart by email: %s", email)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	cartCode, err := h.usecase.ResetCartByEmail(ctx, &models.ResetCartByEmailRequest{
		Email: email,
	})
	if err != nil {
		logrus.Errorf("Failed to reset cart mapping: %v", err)
		return err
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Cart [%s] - Email [%s] has been reset!", cartCode, email))
	return nil
}
