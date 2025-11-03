package helper

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func HandleHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	logrus.Infof("User [%s] needs help", m.Author.Username)
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, HelpEmbed())
	if err != nil {
		logrus.Errorf("Error handle help command %v", err)
		return err
	}
	// s.MessageReactionAdd(m.ChannelID, helpMes.ID, constants.EMOJI_NUMBER_ONE)
	return nil
}

func HelpEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Command List",
		Description: `
			Here is the list of commands!
			`,
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Ask me anything",
				Value: `
				@butler <your request>
				Example: @butler write hello world app in go
				`,
			},
			{
				Name: "Chuyển trạng thái giỏ về available",
				Value: `
					!resetcart <mã giỏ>
					Example: !resetcart 160143
				`,
			},
			{
				Name: "reset tất cả giỏ của user_id",
				Value: `
                    !reset_cart_by_user_id <user_id>
                    Example: !reset_cart_by_user_id 1609
                `,
			},
			{
				Name: "reset tất cả giỏ của email",
				Value: `
                    !reset_cart_by_email <email>
                    Example: !reset_cart_by_email abc@gmail.com
                `,
			},
			{
				Name: "chuẩn bị cho đơn outbound có thể được đi pick",
				Value: `
					!readypick <mã source number1>, <mã source number2>
					Example: !readypick 100224050700001,100224050700002
				`,
			},
			{
				Name: "cho kho xuất hiện để đi pick ở vị trí kho 29 HOANG VIET",
				Value: `
					!showwarehouse <tên kho>
					Example: !showwarehouse SHOP - 29
				`,
			},
			{
				Name: "cho kho xuất hiện để đi pick ở vị trí kho 29 HOANG VIET",
				Value: `
					!show_warehouse_by_id <mã kho>
					Example: !show_warehouse_by_id 540
				`,
			},
			{
				Name: "reset vị trí của các kho bị đổi bởi lệnh !showwarehouse",
				Value: `
					!resetshowwarehouse
					Example: !resetshowwarehouse
				`,
			},
			{
				Name: "reset vị trí của một kho cụ thể bị đổi bởi lệnh !showwarehouse",
				Value: `
					!reset_show_warehouse_by_id <mã kho>
					Example: !reset_show_warehouse_by_id 540
				`,
			},
			{
				Name: "Cập nhật config kho",
				Value: `
					!whcfg add/sub <warehouse_id>
					Example: !whcfg add 14
							 !whcfg sub 14
					Or:
					!whcfg add/sub <warehouse_id> <config>
					Example: !whcfg add 14 1
							 !whcfg sub 14 1
				`,
			},
			{
				Name: "Cập nhật voucher_type của đơn hàng về  1",
				Value: `
					!set_voucher_type_outbound_order <sales_order_number>
					Example: !set_voucher_type_outbound_order 1234567890
				`,
			},
		},
	}
}

func HandleSendImage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	f, err := os.Open("some_img.jpg") // relative path to the main.go file
	if err != nil {
		logrus.Errorf("Error open image %v", err)
		return err
	}
	_, err = s.ChannelFileSend(m.ChannelID, "qwe.jpg", f)
	if err != nil {
		logrus.Errorf("Error handle send image %v", err)
		return err
	}
	return nil
}
