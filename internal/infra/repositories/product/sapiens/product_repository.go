package sapiens

import (
	"context"
	"database/sql"
)

// ProductRepository define os métodos para interagir com a coleção de produtos.
type ProductRepository interface {
	// Insert insere um novo registro no banco de dados.
	// Recebe um contexto e os dados do registro como parâmetros.
	// Retorna o ID do registro inserido ou um erro, caso ocorra.
	Insert(ctx context.Context, data interface{}) (interface{}, error)

	// Get busca um registro pelo seu identificador.
	// Recebe um contexto e o identificador do registro como parâmetros.
	// Retorna os dados do registro encontrado ou um erro, caso ocorra.
	Get(ctx context.Context, identifier int, system string) (interface{}, error)

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
	Update(ctx context.Context, id int, updateData interface{}) error

	// Delete remove um registro do banco de dados.
	// Recebe um contexto e o identifier do registro como parâmetros.
	// Retorna um erro, caso ocorra.
	Delete(ctx context.Context, id int) error
}

type productRepository struct {
	db *sql.DB
}

// NewProductRepository cria uma instância de repositório de produtos.
// Recebe uma configuração de conexão com MongoDB e um logger como parâmetros.
// Retorna uma instância de ProductRepository ou um erro, caso ocorra.
func NewProductRepository(db *sql.DB) (ProductRepository, error) {

	// Retorna o repositório com a coleção já inicializada.
	return &productRepository{
		db: db,
	}, nil
}

// Insert insere um novo registro no banco de dados.
func (r *productRepository) Insert(ctx context.Context, product interface{}) (interface{}, error) {
	return nil, nil
}

// Get busca um registro pelo seu Id.
func (r *productRepository) Get(ctx context.Context, id int, system string) (interface{}, error) {
	return nil, nil
}

// GetAll retorna todos os registros na coleção.
func (r *productRepository) GetAll(ctx context.Context) ([]interface{}, error) {
	return nil, nil
}

// GetBySystem busca registros pelo campo "system".
func (r *productRepository) GetBySystem(ctx context.Context, system string) ([]interface{}, error) {
	return nil, nil
}

// Update atualiza os dados de um registro existente.
func (r *productRepository) Update(ctx context.Context, id int, updateData interface{}) error {
	return nil
}

// Delete remove um registro do banco de dados.
func (r *productRepository) Delete(ctx context.Context, id int) error {
	return nil
}
