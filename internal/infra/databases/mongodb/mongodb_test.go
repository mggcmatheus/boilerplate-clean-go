package databases

import (
	"bifrost-fr/pkg/logger"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"testing"
)

type mockMongoClient struct{}

func (m *mockMongoClient) Database(name string) *mongo.Database {
	return &mongo.Database{}
}

func (m *mockMongoClient) Collection(dbName, collectionName string) (*mongo.Collection, error) {
	if dbName == "" || collectionName == "" {
		return nil, errors.New("invalid database or collection name")
	}
	return &mongo.Collection{}, nil
}

func TestNewMongoDB_Success(t *testing.T) {
	// Inicializa o logger
	logger.InitLogger()
	defer logger.SyncLogger()

	// Configurações de teste
	uri := "mongodb://localhost:27017"
	appName := "TestApp"

	// Executa a função
	mongoDB, err := NewMongoDB(uri, appName)

	// Valida os resultados
	assert.NoError(t, err, "Erro inesperado ao criar a instância do MongoDB")
	assert.NotNil(t, mongoDB, "A instância do MongoDB não deve ser nula")
}

func TestMongoDB_Database(t *testing.T) {
	mockClient := &mockMongoClient{}

	db := mockClient.Database("testDB")
	assert.NotNil(t, db, "Database retornado não deve ser nulo")
}

func TestMongoDB_Collection_Success(t *testing.T) {
	mockClient := &mockMongoClient{}

	collection, err := mockClient.Collection("testDB", "testCollection")
	assert.NoError(t, err, "Erro inesperado ao buscar a coleção")
	assert.NotNil(t, collection, "Collection retornada não deve ser nula")
}

func TestMongoDB_Collection_Error(t *testing.T) {
	mockClient := &mockMongoClient{}

	collection, err := mockClient.Collection("", "")
	assert.Error(t, err, "Esperava um erro ao passar nomes inválidos")
	assert.Nil(t, collection, "Collection retornada deve ser nula em caso de erro")
}
