package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var Log *zap.Logger

// InitLogger inicializa o Zap logger e atribui à variável global Log
func InitLogger() {
	// Configuração padrão do Zap (desenvolvimento ou produção)
	var config zap.Config
	if os.Getenv("APP_ENV") == "production" {
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	} else {
		config = zap.NewDevelopmentConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	// Saída no console
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Inicializa o logger
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Erro ao iniciar o logger: %v", err)
	}

	Log = logger
}

// SyncLogger garante que todos os logs sejam descarregados
func SyncLogger() {
	if Log != nil {
		_ = Log.Sync()
	}
}
