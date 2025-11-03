package delivery

import (
	"butler/application/commands/helper"
	"butler/application/lib"
	"butler/constants"
	"fmt"
	"strings"

	cartHandler "butler/application/domains/cart/delivery/discord/handler"
	kpiHandler "butler/application/domains/kpi/delivery/discord/handler"
	pickHandler "butler/application/domains/pick/delivery/discord/handler"
	pickPackHandler "butler/application/domains/pick_pack/delivery/discord/handler"
	makersuiteHandler "butler/application/domains/promt_ai/makersuite/handler"
	warehouseHandler "butler/application/domains/warehouse/delivery/discord/handler"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	GetCommandsHandler(*discordgo.Session, *discordgo.MessageCreate)
	GetReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd)
}

type commandHandler struct {
	lib               *lib.Lib
	discord           *discordgo.Session
	makersuiteHandler makersuiteHandler.Handler
	cartHandler       cartHandler.Handler
	pickHandler       pickHandler.Handler
	whHandler         warehouseHandler.Handler
	kpiHandler        kpiHandler.Handler
	pickPackHandler   pickPackHandler.Handler
}

func NewCommandHandler(
	lib *lib.Lib,
	discord *discordgo.Session,
	makersuiteHandler makersuiteHandler.Handler,
	cartHandler cartHandler.Handler,
	pickHandler pickHandler.Handler,
	whHandler warehouseHandler.Handler,
	pickPackHandler pickPackHandler.Handler,
) Handler {
	return &commandHandler{
		discord:           discord,
		makersuiteHandler: makersuiteHandler,
		cartHandler:       cartHandler,
		pickHandler:       pickHandler,
		whHandler:         whHandler,
		lib:               lib,
		pickPackHandler:   pickPackHandler,
	}
}

func (c *commandHandler) GetCommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("runtime error: %v", err)
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Something went wrong: %v", err))
		}
	}()

	if m.Author.ID == s.State.User.ID {
		return
	}
	if !strings.HasPrefix(m.Content, constants.BOT_COMMAND_PREFIX) && !helper.CheckMention(m, s.State.User) {
		return
	}

	var err error
	switch {
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_HELP):
		err = helper.HandleHelpCommand(s, m)
	case helper.CheckMention(m, s.State.User):
		err = c.makersuiteHandler.Ask(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_RESET_CART):
		err = c.cartHandler.ResetCart(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_RESET_CART_BY_USER_ID):
		err = c.cartHandler.ResetCartByUserId(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_RESET_CART_BY_EMAIL):
		err = c.cartHandler.ResetCartByEmail(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_READY_PICK):
		err = c.pickHandler.ReadyPickOutbound(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_PICK):
		err = c.pickHandler.PreparePick(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_SHOW_WAREHOUSE):
		err = c.whHandler.ShowWarehouse(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_SHOW_WAREHOUSE_BY_ID):
		err = c.whHandler.ShowWarehouseById(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_RESET_SHOW_WAREHOUSE):
		err = c.whHandler.ResetShowWarehouse(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_RESET_SHOW_WAREHOUSE_BY_ID):
		err = c.whHandler.ResetShowWarehouseById(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_COUNT_KPI),
		helper.CheckPrefixCommand(m.Content, constants.COMMAND_COUNT_PROD_KPI):
		err = c.kpiHandler.CountKpi(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_WH_CONFIG):
		err = c.whHandler.ShowConfigWarehouse(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_PICK_PACK_KAFKA):
		err = c.pickPackHandler.PickPackKafka(s, m)
	case helper.CheckPrefixCommand(m.Content, constants.COMMAND_SET_VOUCHER_TYPE_OUTBOUND_ORDER):
		err = c.pickPackHandler.SetOutboundOrderVoucherType(s, m)
	}
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
	}
}
