package main

import (
	"bifrost-fr/pkg/logger"
	"gofr.dev/pkg/gofr"
)

func main() {
	// Inicializa o logger
	logger.InitLogger()
	defer logger.SyncLogger()

	// Inicializa o objeto do GoFr
	app := gofr.New()

	app.Run()
}
