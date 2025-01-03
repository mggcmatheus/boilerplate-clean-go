package sapiens

import (
	"context"
	"database/sql"
)

// ProviderRepository define os métodos para interagir com a coleção de fornecedores.
type ProviderRepository interface {
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

type providerRepository struct {
	db *sql.DB
}

// NewProviderRepository cria uma instância de repositório de fornecedores.
// Recebe uma configuração de conexão com Oracle e um logger como parâmetros.
// Retorna uma instância de ProviderRepository ou um erro, caso ocorra.
func NewProviderRepository(db *sql.DB) (ProviderRepository, error) {

	// Retorna o repositório inicializada.
	return &providerRepository{
		db: db,
	}, nil
}

// Insert insere um novo fornecedor no banco de dados.
func (r *providerRepository) Insert(ctx context.Context, data interface{}) (interface{}, error) {
	return nil, nil
}

// Get busca um provider pelo seu identificador.
func (r *providerRepository) Get(ctx context.Context, identifier, system string) (interface{}, error) {
	return nil, nil
}

// GetAll retorna todos os registros na coleção.
func (r *providerRepository) GetAll(ctx context.Context) ([]interface{}, error) {
	return nil, nil
}

// GetBySystem busca registros pelo campo "system".
func (r *providerRepository) GetBySystem(ctx context.Context, system string) ([]interface{}, error) {
	return nil, nil
}

// Update atualiza os dados de um registro existente.
func (r *providerRepository) Update(ctx context.Context, identifier string, updateData interface{}) error {
	return nil
}

// Delete remove um registro do banco de dados.
func (r *providerRepository) Delete(ctx context.Context, identifier string) error {
	return nil
}
