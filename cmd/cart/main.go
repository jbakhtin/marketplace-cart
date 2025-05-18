package main

import (
	"context"
	"fmt"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/config"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/logger/zap"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/server/rest"
	"github.com/jbakhtin/marketplace-cart/internal/infrastucture/storage/postgres"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart"
	"github.com/jbakhtin/marketplace-cart/pkg/closer"
	"github.com/jbakhtin/marketplace-cart/pkg/starter"
	"log"
	"os/signal"
	"syscall"
)

var err error
var logger zap.Logger
var str starter.Starter
var clr closer.Closer
var cfg config.Config
var restServer rest.Server

func init() {
	cfg, err = config.NewConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	logger, err = zap.NewLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}

	starterBuilder := starter.New()
	closerBuilder := closer.New()

	orderRepository, err := postgres.NewOrderStorage()
	stockRepository, err := postgres.NewStockStorage()

	lomsModule, err := cart.InitModule(logger, orderRepository, stockRepository)

	restServer, err = rest.NewWebServer(&cfg, logger, lomsModule)
	if err != nil {
		log.Fatal(err)
	}
	starterBuilder.Add(restServer.Start)
	closerBuilder.Add(restServer.Shutdown)

	str = starterBuilder.Build()
	clr = closerBuilder.Build()
}

func main() {
	osCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	err = str.Start(osCtx)
	if err != nil {
		fmt.Println(err.Error())
	}

	<-osCtx.Done()

	err = clr.Close(osCtx)
	if err != nil {
		fmt.Println(err.Error())
	}
}
