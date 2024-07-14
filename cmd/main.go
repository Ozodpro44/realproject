package main

import (
	"fmt"
	"real_project/config"
	"real_project/pkg/db"
	log "real_project/pkg/logger"

	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {
	cfg := config.Load()

	log := log.NewLogger(cfg.GeneralConfig)

	pgxPool, err := db.ConnDB(cfg.PgConfig)
	if err != nil {
		log.Error("error on connecting with databse", logger.Error(err))
	}

	fmt.Println(pgxPool)
}