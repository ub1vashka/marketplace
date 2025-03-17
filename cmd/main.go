package main

import (
	"github.com/ub1vashka/marketplace/internal/config"
	"github.com/ub1vashka/marketplace/internal/logger"
	"github.com/ub1vashka/marketplace/internal/server"
	"github.com/ub1vashka/marketplace/internal/service"
	"github.com/ub1vashka/marketplace/internal/storage"
)

func main() {
	cfg := config.ReadConfig()
	log := logger.Get(cfg.Debug)

	stor := storage.New()
	userService := service.NewUserService(stor)
	productService := service.NewProductService(stor)
	pusrchaseService := service.NewPurchaseService(stor)
	serve := server.New(cfg, userService, productService, pusrchaseService)
	if err := serve.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}

}
