package provider

import (
	provider_postgres_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/provider"
	supply_postgres_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/supply"
)

type ProviderService struct {
	ProviderPostgresRepository *provider_postgres_repository.ProviderRepository
	SupplyPostgresRepository   *supply_postgres_repository.SupplyRepository
}

// NewProviderService will create an object that represent the provider.Service interface
func NewProviderService(provider_postgres_repository *provider_postgres_repository.ProviderRepository, supply_postgres_repository *supply_postgres_repository.SupplyRepository) *ProviderService {
	return &ProviderService{
		ProviderPostgresRepository: provider_postgres_repository,
		SupplyPostgresRepository:   supply_postgres_repository,
	}
}
