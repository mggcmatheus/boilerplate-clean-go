package hub

import (
	"context"
	"database/sql"
)

// CustomerRepository define os métodos para interagir com os clientes no HUB.
type CustomerRepository interface {
	// InsertCustomer insere um novo cliente no banco de dados.
	// Recebe um contexto e os dados do cliente como parâmetros.
	// Retorna o ID do cliente inserido ou um erro, caso ocorra.
	InsertCustomer(ctx context.Context, customer interface{}) (interface{}, error)

	// GetCustomer busca um cliente pelo seu identificador.
	// Recebe um contexto e o identificador do cliente como parâmetros.
	// Retorna os dados do cliente encontrado ou um erro, caso ocorra.
	GetCustomer(ctx context.Context, identifier, system string) (interface{}, error)

	// GetAllCustomers retorna todos os clientes na coleção.
	// Recebe um contexto como parâmetro.
	// Retorna uma lista de clientes ou um erro, caso ocorra.
	GetAllCustomers(ctx context.Context) ([]interface{}, error)

	// GetCustomerBySystem busca clientes pelo campo "system".
	// Recebe um contexto e o valor de "system" como parâmetros.
	// Retorna uma lista de clientes ou um erro, caso ocorra.
	GetCustomerBySystem(ctx context.Context, system string) ([]interface{}, error)

	// UpdateCustomer atualiza os dados de um cliente existente.
	// Recebe um contexto, o identifier do cliente e os dados a serem atualizados como parâmetros.
	// Retorna um erro, caso ocorra.
	UpdateCustomer(ctx context.Context, identifier string, updateData interface{}) error

	// DeleteCustomer remove um cliente do banco de dados.
	// Recebe um contexto e o identifier do cliente como parâmetros.
	// Retorna um erro, caso ocorra.
	DeleteCustomer(ctx context.Context, identifier string) error
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) (CustomerRepository, error) {

	return &customerRepository{
		db: db,
	}, nil
}

// InsertCustomer insere um novo cliente no banco de dados.
func (r *customerRepository) InsertCustomer(ctx context.Context, customer interface{}) (interface{}, error) {
	return nil, nil
}

// GetCustomer busca um cliente pelo seu identifier.
func (r *customerRepository) GetCustomer(ctx context.Context, identifier, system string) (interface{}, error) {
	return nil, nil
}

// GetAllCustomers retorna todos os clientes na coleção.
func (r *customerRepository) GetAllCustomers(ctx context.Context) ([]interface{}, error) {
	return nil, nil
}

// GetCustomerBySystem busca clientes pelo campo "system".
func (r *customerRepository) GetCustomerBySystem(ctx context.Context, system string) ([]interface{}, error) {
	return nil, nil
}

// UpdateCustomer atualiza os dados de um cliente existente.
func (r *customerRepository) UpdateCustomer(ctx context.Context, identifier string, updateData interface{}) error {
	return nil
}

// DeleteCustomer remove um cliente do banco de dados.
func (r *customerRepository) DeleteCustomer(ctx context.Context, identifier string) error {
	return nil
}
