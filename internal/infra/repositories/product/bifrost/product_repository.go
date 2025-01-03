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

// ProductRepository define os métodos para interagir com a coleção de produtos.
type ProductRepository interface {
	// InsertProduct insere um novo produto no banco de dados.
	// Recebe um contexto e os dados do produto como parâmetros.
	// Retorna o ID do produto inserido ou um erro, caso ocorra.
	InsertProduct(ctx context.Context, product interface{}) (interface{}, error)

	// GetProduct busca um produto pelo seu id.
	// Recebe um contexto e o id do produto como parâmetros.
	// Retorna os dados do produto encontrado ou um erro, caso ocorra.
	GetProduct(ctx context.Context, id int, system string) (interface{}, error)

	// GetAllProducts retorna todos os produtos na coleção.
	// Recebe um contexto como parâmetro.
	// Retorna uma lista de produtos ou um erro, caso ocorra.
	GetAllProducts(ctx context.Context) ([]interface{}, error)

	// GetProductBySystem busca produtos pelo campo "system".
	// Recebe um contexto e o valor de "system" como parâmetros.
	// Retorna uma lista de produtos ou um erro, caso ocorra.
	GetProductBySystem(ctx context.Context, system string) ([]interface{}, error)

	// UpdateProduct atualiza os dados de um produto existente.
	// Recebe um contexto, o id do produto e os dados a serem atualizados como parâmetros.
	// Retorna um erro, caso ocorra.
	UpdateProduct(ctx context.Context, id int, updateData interface{}) error

	// DeleteProduct remove um produto do banco de dados.
	// Recebe um contexto e o id do produto como parâmetros.
	// Retorna um erro, caso ocorra.
	DeleteProduct(ctx context.Context, id int) error
}

type productRepository struct {
	collection *mongo.Collection
}

// NewProductRepository cria uma instância de repositório de produtos.
// Recebe uma configuração de conexão com MongoDB e um logger como parâmetros.
// Retorna uma instância de ProductRepository ou um erro, caso ocorra.
func NewProductRepository(client *mongo.Client) (ProductRepository, error) {
	// Obtém a coleção "customers" do banco de dados configurado.
	collection := client.Database("bifrost").Collection("products")

	// Retorna o repositório com a coleção já inicializada.
	return &productRepository{
		collection: collection,
	}, nil
}

// InsertProduct insere um novo produto no banco de dados.
func (r *productRepository) InsertProduct(ctx context.Context, product interface{}) (interface{}, error) {
	return nil, nil
}

// GetProduct busca um produto pelo seu ID.
func (r *productRepository) GetProduct(ctx context.Context, id int, system string) (interface{}, error) {

	filter := bson.D{
		{Key: "id", Value: id},
		{Key: "system", Value: system},
	}

	var product domain.Product

	err := r.collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Log.Error("não foi encontrado produto com o id informado", zap.Error(err), zap.Int("id", id), zap.String("system", system))
			return nil, err
		}
		logger.Log.Error("erro ao buscar o produto com o id informado", zap.Error(err))
		return nil, err
	}
	return nil, nil
}

// GetAllProducts retorna todos os produtos na coleção.
func (r *productRepository) GetAllProducts(ctx context.Context) ([]interface{}, error) {
	return nil, nil
}

// GetProductBySystem busca produtos pelo campo "system".
func (r *productRepository) GetProductBySystem(ctx context.Context, system string) ([]interface{}, error) {
	return nil, nil
}

// UpdateProduct atualiza os dados de um produto existente.
func (r *productRepository) UpdateProduct(ctx context.Context, id int, updateData interface{}) error {
	return nil
}

// DeleteProduct remove um produto do banco de dados.
func (r *productRepository) DeleteProduct(ctx context.Context, id int) error {
	return nil
}
