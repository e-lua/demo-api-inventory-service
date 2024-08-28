package automated

import (
	provider_postgres_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/provider"
	supply_postgres_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/supply"
	warehouse_postgres_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/warehouse"
	job_redis_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/redis/job"
)

type AutomatedService struct {
	WarehousePostgresRepository *warehouse_postgres_repository.WarehouseRepository
	SupplyPostgresRepository    *supply_postgres_repository.SupplyRepository
	ProviderPostgresRepository  *provider_postgres_repository.ProviderRepository
	JobRedisRepository          *job_redis_repository.JobRedisRepository
}

// NewAutomatedService will create an object that represent the automated.Service interface
func NewAutomatedService(warehouse_postgres_repository *warehouse_postgres_repository.WarehouseRepository, supply_postgres_repository *supply_postgres_repository.SupplyRepository, provider_postgres_repository *provider_postgres_repository.ProviderRepository, job_redis_repository *job_redis_repository.JobRedisRepository) *AutomatedService {
	return &AutomatedService{
		WarehousePostgresRepository: warehouse_postgres_repository,
		SupplyPostgresRepository:    supply_postgres_repository,
		ProviderPostgresRepository:  provider_postgres_repository,
		JobRedisRepository:          job_redis_repository,
	}
}
