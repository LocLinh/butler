package server

import (
	commandHandler "butler/application/commands/handler"
	initCartHandler "butler/application/domains/cart/delivery/discord/handler"
	initPickHandler "butler/application/domains/pick/delivery/discord/handler"
	initPickPackHandler "butler/application/domains/pick_pack/delivery/discord/handler"
	initPromtAiHandler "butler/application/domains/promt_ai/makersuite/handler"
	initServices "butler/application/domains/services/init"
	initWarehouseHandler "butler/application/domains/warehouse/delivery/discord/handler"
	"butler/application/lib"
	"context"

	"butler/config"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
	"github.com/sirupsen/logrus"

	"google.golang.org/api/option"
)

type Server struct {
	cfg         *config.Config
	discordBot  *discordgo.Session
	genaiClient *genai.Client
	lib         *lib.Lib
}

func NewServer(cfg *config.Config) *Server {
	// discord
	dg, err := discordgo.New("Bot " + cfg.DiscordBot.Butler.Token)
	if err != nil {
		logrus.Fatalf("init discord bot err: %v", err)
	}
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// genai client
	genaiClient, err := genai.NewClient(context.Background(), option.WithAPIKey(cfg.Makersuite.ApiKey))
	if err != nil {
		//logrus.Fatalf("init genai client err: %s", err)
	}

	return &Server{
		cfg:         cfg,
		discordBot:  dg,
		genaiClient: genaiClient,
		lib:         lib.InitLib(cfg),
	}
}

func (s *Server) Start() {
	s.run()
}

func (s *Server) Stop() {
	s.discordBot.Close()
	s.genaiClient.Close()
}

func (s *Server) run() {
	err := s.discordBot.Open()
	if err != nil {
		logrus.Fatalf("opening connection discord err: %v", err)
		return
	}
	// init services
	services := initServices.InitService(s.cfg, s.lib.Db, s.genaiClient)

	// init external
	promtAiHandler := initPromtAiHandler.InitHandler(s.cfg, services)

	// init cart handler
	cartHandler := initCartHandler.InitHandler(services)

	// init pick handler
	pickHandler := initPickHandler.InitHandler(s.lib, services)

	// init warehouse handler
	warehouseHandler := initWarehouseHandler.InitHandler(s.lib, services)

	// init pick pack handler
	pickPackHandler := initPickPackHandler.InitHandler(s.lib, s.cfg, services)

	// register handler for discord command
	commandHandler := commandHandler.NewCommandHandler(s.lib, s.discordBot, promtAiHandler, cartHandler, pickHandler, warehouseHandler, pickPackHandler)
	s.discordBot.AddHandler(commandHandler.GetCommandsHandler)
	s.discordBot.AddHandler(commandHandler.GetReactionHandler)

	logrus.Infof("start server success")
}
