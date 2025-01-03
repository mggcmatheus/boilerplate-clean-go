package gatec

import (
	"context"
	"database/sql"
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
	db *sql.DB
}

// NewCustomerRepository cria uma instância de repositório de clientes.
// Recebe uma configuração de conexão com Oracle e um logger como parâmetros.
// Retorna uma instância de CustomerRepository ou um erro, caso ocorra.
func NewCustomerRepository(db *sql.DB) (CustomerRepository, error) {

	// Retorna o repositório inicializada.
	return &customerRepository{
		db: db,
	}, nil
}

// Insert insere um novo cliente no banco de dados.
func (r *customerRepository) Insert(ctx context.Context, data interface{}) (interface{}, error) {
	return nil, nil
}

// Get busca um cliente pelo seu identifier.
func (r *customerRepository) Get(ctx context.Context, identifier, system string) (interface{}, error) {
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
