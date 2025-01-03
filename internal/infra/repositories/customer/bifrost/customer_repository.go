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

// CustomerRepository define os métodos para interagir com a coleção de clientes.
type CustomerRepository interface {
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
type customerRepository struct {
	collection *mongo.Collection
}

// NewCustomerRepository cria uma instância de repositório de clientes.
// Recebe uma configuração de conexão com MongoDB e um logger como parâmetros.
// Retorna uma instância de CustomerRepository ou um erro, caso ocorra.
func NewCustomerRepository(client *mongo.Client) (CustomerRepository, error) {

	// Obtém a coleção "customers" do banco de dados configurado.
	collection := client.Database("bifrost").Collection("customers")

	// Retorna o repositório com a coleção já inicializada.
	return &customerRepository{
		collection: collection,
	}, nil
}

// Insert insere um novo cliente no banco de dados.
func (r *customerRepository) Insert(ctx context.Context, data interface{}) (interface{}, error) {
	return nil, nil
}

// Get busca um cliente pelo seu identifier.
func (r *customerRepository) Get(ctx context.Context, identifier, system string) (interface{}, error) {
	filter := bson.D{
		{Key: "identifier", Value: identifier},
		{Key: "system", Value: system},
	}

	var customer domain.Customer

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

// GetAll retorna todos os clientes na coleção.
func (r *customerRepository) GetAll(ctx context.Context) ([]interface{}, error) {
	return nil, nil
}

// GetBySystem busca clientes pelo campo "system".
func (r *customerRepository) GetBySystem(ctx context.Context, system string) ([]interface{}, error) {
	return nil, nil
}

// Update atualiza os dados de um cliente existente.
func (r *customerRepository) Update(ctx context.Context, identifier string, updateData interface{}) error {
	return nil
}

// Delete remove um cliente do banco de dados.
func (r *customerRepository) Delete(ctx context.Context, identifier string) error {
	return nil
}
