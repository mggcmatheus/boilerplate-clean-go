package databases

import (
	"bifrost-fr/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.uber.org/zap"
	"time"
)

// MongoDB define os métodos para interagir com o MongoDB.
type MongoDB interface {
	Database(name string) *mongo.Database
	Collection(dbName, collectionName string) (*mongo.Collection, error)
}

type mongoDB struct {
	client *mongo.Client
}

// NewMongoDB cria uma nova instância de conexão com o MongoDB.
// Retorna uma implementação da interface MongoDB.
func NewMongoDB(uri string, appName string) (MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri).SetAppName(appName)

	var (
		newClient *mongo.Client
		err       error
	)

	for {
		// Tenta conectar ao MongoDB
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		newClient, err = mongo.Connect(clientOptions)
		if err != nil {
			logger.Log.Error("Falha ao tentar conectar no MongoDB:", zap.Error(err))
			logger.Log.Info("Tentando conectar novamente em 5 minutos...")
			time.Sleep(5 * time.Minute)
			continue
		}

		// Testa a conexão com um ping
		err = newClient.Ping(ctx, readpref.Primary())
		if err != nil {
			logger.Log.Error("Erro ao tentar pingar no MongoDB:", zap.Error(err))
			errDisconnect := newClient.Disconnect(context.Background())
			if errDisconnect != nil {
				logger.Log.Error("Erro ao desconectar o cliente MongoDB:", zap.Error(err))
			}
			logger.Log.Info("Tentando pingar novamente em 5 minutos...")
			time.Sleep(5 * time.Minute)
			continue
		}

		// Conexão bem-sucedida
		break
	}

	return &mongoDB{
		client: newClient,
	}, nil
}

// Database retorna uma instância de um banco de dados do MongoDB.
func (r *mongoDB) Database(name string) *mongo.Database {
	return r.client.Database(name)
}

// Collection retorna uma instância de uma coleção do MongoDB.
func (r *mongoDB) Collection(dbName, collectionName string) (*mongo.Collection, error) {
	db := r.Database(dbName)
	if db == nil {
		return nil, mongo.ErrClientDisconnected
	}
	return db.Collection(collectionName), nil
}
