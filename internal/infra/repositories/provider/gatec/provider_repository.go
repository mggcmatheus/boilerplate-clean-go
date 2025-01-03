package gatec

import (
	domain "bifrost/internal/domain/entities"
	"bifrost/internal/infra/config"
	"context"
	"database/sql"
)

// ProviderRepository define os métodos para interagir com a coleção de fornecedores.
type ProviderRepository interface {
	// InsertProvider insere um novo fornecedor no banco de dados.
	// Recebe um contexto e os dados do fornecedor como parâmetros.
	// Retorna o identifier do fornecedor inserido ou um erro, caso ocorra.
	InsertProvider(ctx context.Context, provider interface{}) (interface{}, error)

	// GetProvider busca um fornecedor pelo seu identifier.
	// Recebe um contexto e o identifier do fornecedor como parâmetros.
	// Retorna os dados do fornecedor encontrado ou um erro, caso ocorra.
	GetProvider(ctx context.Context, identifier, system string) (interface{}, error)

	// GetAllProviders retorna todos os fornecedores na coleção.
	// Recebe um contexto como parâmetro.
	// Retorna uma lista de fornecedores ou um erro, caso ocorra.
	GetAllProviders(ctx context.Context) ([]interface{}, error)

	// GetProviderBySystem busca fornecedores pelo campo "system".
	// Recebe um contexto e o valor de "system" como parâmetros.
	// Retorna uma lista de fornecedores ou um erro, caso ocorra.
	GetProviderBySystem(ctx context.Context, system string) ([]interface{}, error)

	// UpdateProvider atualiza os dados de um fornecedor existente.
	// Recebe um contexto, o identifier do fornecedor e os dados a serem atualizados como parâmetros.
	// Retorna um erro, caso ocorra.
	UpdateProvider(ctx context.Context, identifier string, updateData interface{}) error

	// DeleteProvider remove um fornecedor do banco de dados.
	// Recebe um contexto e o identifier do fornecedor como parâmetros.
	// Retorna um erro, caso ocorra.
	DeleteProvider(ctx context.Context, identifier string) error
}

type providerRepository struct {
	db     *sql.DB
	logger domain.Logger
}

// NewProviderRepository cria uma instância de repositório de fornecedores.
// Recebe uma configuração de conexão com Oracle e um logger como parâmetros.
// Retorna uma instância de ProviderRepository ou um erro, caso ocorra.
func NewProviderRepository(cfg config.GatecConfig, logger domain.Logger) (ProviderRepository, error) {
	// Instancia a conexão com o Oracle.
	db, err := config.GetOracleConnection(cfg.User, cfg.Password, cfg.Address)
	if err != nil {
		logger.WithFields(domain.LogFormat{
			Error: err,
		}).Error("erro ao obter conexão oracle no repositório customer")
		return nil, err
	}
	// Retorna o repositório inicializada.
	return &providerRepository{
		db:     db,
		logger: logger,
	}, nil
}

// InsertProvider insere um novo fornecedor no banco de dados.
func (r *providerRepository) InsertProvider(ctx context.Context, provider interface{}) (interface{}, error) {
	return nil, nil
}

// GetProvider busca um provider pelo seu identificador.
func (r *providerRepository) GetProvider(ctx context.Context, identifier, system string) (interface{}, error) {
	return nil, nil
}

// GetAllProviders retorna todos os fornecedores na coleção.
func (r *providerRepository) GetAllProviders(ctx context.Context) ([]interface{}, error) {
	return nil, nil
}

// GetProviderBySystem busca fornecedor pelo campo "system".
func (r *providerRepository) GetProviderBySystem(ctx context.Context, system string) ([]interface{}, error) {
	return nil, nil
}

// UpdateProvider atualiza os dados de um fornecedor existente.
func (r *providerRepository) UpdateProvider(ctx context.Context, identifier string, updateData interface{}) error {
	return nil
}

// DeleteProvider remove um fornecedor do banco de dados.
func (r *providerRepository) DeleteProvider(ctx context.Context, identifier string) error {
	return nil
}
