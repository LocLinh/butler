package main

import (
	"butler/application/lib"
	"butler/config"
	_ "butler/pkg/log"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"butler/application/domains/pick_pack/models"
	pickPackUc "butler/application/domains/pick_pack/usecase"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		logrus.Fatalf("Err get config: %v\n", err)
	}
	lib := lib.InitLib(cfg)

	runTestPickPack(lib, cfg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Ctrl + C to exit program")
	<-quit
}

func runTestPickPack(lib *lib.Lib, cfg *config.Config) {
	ctx := context.Background()
	initPickPackUc := pickPackUc.InitUseCase(lib, cfg, nil)

	initPickPackUc.AutoPickPack(ctx, models.AutoPickPackRequest{
		LoginRequest: models.LoginRequest{
			LoginDiscordRequest: models.LoginDiscordRequest{
				Login:    "sonplh@hasaki.vn",
				Password: "12345a@A",
			},
		},
		SalesOrderNumber: "123456",
		WarehouseId:      14,
		ShippingUnitId:   104, //
	})

}
