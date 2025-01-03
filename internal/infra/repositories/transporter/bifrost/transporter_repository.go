package bifrost

import (
	domain "bifrost-fr/internal/domain/entities"
	"bifrost-fr/pkg/logger"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

// TransporterRepository define os métodos para interagir com a coleção do banco de dados.
type TransporterRepository interface {
	// Insert insere um novo registro no banco de dados.
	// Recebe um contexto e os dados do registro como parâmetros.
	// Retorna o ID do registro inserido ou um erro, caso ocorra.
	Insert(ctx context.Context, data interface{}) (interface{}, error)

	// Get busca um registro pelo seu identificador.
	// Recebe um contexto e o identificador do registro como parâmetros.
	// Retorna os dados do registro encontrado ou um erro, caso ocorra.
	Get(ctx context.Context, identifier, system string) (interface{}, error)

	// GetAll retorna todos os registros na coleção.
	// Recebe um contexto como parâmetro.
	// Retorna uma lista de registros ou um erro, caso ocorra.
	GetAll(ctx context.Context) ([]interface{}, error)

	// GetBySystem busca registros pelo campo "system".
	// Recebe um contexto e o valor de "system" como parâmetros.
	// Retorna uma lista de registros ou um erro, caso ocorra.
	GetBySystem(ctx context.Context, system string) ([]interface{}, error)

	// Update atualiza os dados de um registro existente.
	// Recebe um contexto, o identifier do registro e os dados a serem atualizados como parâmetros.
	// Retorna um erro, caso ocorra.
	Update(ctx context.Context, identifier string, updateData interface{}) error

	// Delete remove um registro do banco de dados.
	// Recebe um contexto e o identifier do registro como parâmetros.
	// Retorna um erro, caso ocorra.
	Delete(ctx context.Context, identifier string) error
}

type transporterRepository struct {
	collection *mongo.Collection
}

// NewTransporterRepository cria uma instância de repositório de clientes.
// Recebe uma configuração de conexão com MongoDB e um logger como parâmetros.
// Retorna uma instância de TransporterRepository ou um erro, caso ocorra.
func NewTransporterRepository(client *mongo.Client) (TransporterRepository, error) {

	// Obtém a coleção especificada do banco de dados configurado.
	collection := client.Database("bifrost").Collection("transporters")

	// Retorna o repositório com a coleção já inicializada.
	return &transporterRepository{
		collection: collection,
	}, nil
}

// Insert insere um novo registro no banco de dados.
func (r *transporterRepository) Insert(ctx context.Context, data interface{}) (interface{}, error) {
	return nil, nil
}

// Get busca um registro pelo seu identifier.
func (r *transporterRepository) Get(ctx context.Context, identifier, system string) (interface{}, error) {
	filter := bson.D{
		{Key: "identifier", Value: identifier},
		{Key: "system", Value: system},
	}

	var customer domain.Transporter

	err := r.collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Log.Error("não foi encontrado cliente com o identificador informado", zap.Error(err), zap.String("identifier", identifier), zap.String("system", system))
			return nil, err
		}
		logger.Log.Error("erro ao buscar o cliente com o identificador informado", zap.Error(err))
		return nil, err
	}
	return nil, nil
}

// GetAll retorna todos os registros na coleção.
func (r *transporterRepository) GetAll(ctx context.Context) ([]interface{}, error) {
	return nil, nil
}

// GetBySystem busca registros pelo campo "system".
func (r *transporterRepository) GetBySystem(ctx context.Context, system string) ([]interface{}, error) {
	return nil, nil
}

// Update atualiza os dados de um registro existente.
func (r *transporterRepository) Update(ctx context.Context, identifier string, updateData interface{}) error {
	return nil
}

// Delete remove um registro do banco de dados.
func (r *transporterRepository) Delete(ctx context.Context, identifier string) error {
	return nil
}
