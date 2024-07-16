package main

import (
	"context"
	"fmt"
	"real_project/api"
	"real_project/config"
	"real_project/pkg/db"
	log "real_project/pkg/logger"
	"real_project/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

var ctx = context.Background()

func main() {
	cfg := config.Load()

	log := log.NewLogger(cfg.GeneralConfig)

	pgxPool, err := db.ConnDB(cfg.PgConfig)
	if err != nil {
		log.Error("error on connecting with databse", logger.Error(err))
		return
	}
	log.Debug("successfully connected to db.")

	fmt.Println(pgxPool)

	redisCli, err := db.ConnRedis(log, ctx, cfg.RedisConfig)
	if err != nil {
		log.Error("error on connecting with redis", logger.Error(err))
		return
	}
	log.Debug("successfully connected to redis.")

	fmt.Println(redisCli)

	storage := storage.NewStorage(pgxPool, log)

	engine := api.Api(api.Options{
		Storage: storage,
		Log: log,
	})

	log.Debug("server is running on",logger.String("port",cfg.GeneralConfig.HTTPPort))
	engine.Run(cfg.GeneralConfig.HTTPPort)
}
