package databases

import (
	"bifrost-fr/pkg/logger"
	"database/sql"
	"fmt"
	_ "github.com/sijms/go-ora/v2"
	"go.uber.org/zap"
	"sync"
)

type Oracle interface{}

type oracle struct {
	client *sql.DB
}

var (
	once      sync.Once
	client    *sql.DB
	errClient error
)

func NewOracle(connectionString string) (Oracle, error) {
	once.Do(func() {
		// Interpolar os valores na string DSN
		dsn := fmt.Sprintf("oracle://%s", connectionString)

		// Conecta ao banco de dados Oracle
		client, errClient = sql.Open("oracle", dsn)
		if errClient != nil {
			logger.Log.Error("Falha ao tentar conectar no Oracle:", zap.Error(errClient))
			return
		}

		// Define o número máximo de conexões no pool
		client.SetMaxOpenConns(10) // Número máximo de conexões abertas simultaneamente

		// Testa a conexão
		errClient = client.Ping()
		if errClient != nil {
			client.Close()
			logger.Log.Error("Falha ao testar a conectar no Oracle:", zap.Error(errClient))
			return
		}

		logger.Log.Info("Conexão com Oracle estabelecida com sucesso")
	})

	return &oracle{
		client: client,
	}, nil
}

func (r *oracle) Close() {
	if r.client != nil {
		errClient = r.client.Close()
		if errClient != nil {
			logger.Log.Error("Erro ao fechar a conexão com Oracle:", zap.Error(errClient))
		} else {
			logger.Log.Info("Conexão com o Oracle encerrada com sucesso")
		}
	} else {
		logger.Log.Warn("Tentativa de fechar conexão Oracle nula.")
	}
}
