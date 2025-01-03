package databases

import (
	"bifrost-fr/pkg/logger"
	"database/sql"
	"go.uber.org/zap"
	"sync"
)

type PostgreSQL interface{}

type postgreSQL struct {
	client *sql.DB
}

var (
	once      sync.Once
	client    *sql.DB
	errClient error
)

func NewPostgreSQL(connectionString string) (PostgreSQL, error) {
	once.Do(func() {
		client, errClient = sql.Open("postgres", connectionString)
		if errClient != nil {
			logger.Log.Error("Falha ao tentar conectar o PostgreSQL:", zap.Error(errClient))
			return
		}

		// Define o número máximo de conexões no pool
		client.SetMaxOpenConns(5) // Número máximo de conexões abertas simultaneamente
		client.SetMaxIdleConns(2)

		// Testa a conexão
		errClient = client.Ping()
		if errClient != nil {
			client.Close()
			logger.Log.Error("Falha ao testar a conectar no PostgreSQL:", zap.Error(errClient))
			return
		}
		logger.Log.Info("Conexão com PostgreSQL estabelecida com sucesso")
	})

	// Retorna o banco de dados e o erro acumulado
	if errClient != nil {
		return nil, errClient
	}

	return errClient, nil
}

func (r *postgreSQL) Close() {
	if r.client != nil {
		errClient = r.client.Close()
		if errClient != nil {
			logger.Log.Error("Erro ao fechar a conexão com PostgreSQL:", zap.Error(errClient))
		} else {
			logger.Log.Info("Conexão com o PostgreSQL encerrada com sucesso")
		}
	} else {
		logger.Log.Warn("Tentativa de fechar conexão PostgreSQL nula.")
	}
}
