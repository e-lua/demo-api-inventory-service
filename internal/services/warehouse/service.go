package warehouse

import (
	supply_postgres_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/supply"
	warehouse_postgres_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/warehouse"
)

type WarehouseService struct {
	WarehousePostgresRepository *warehouse_postgres_repository.WarehouseRepository
	SupplyPostgresRepository    *supply_postgres_repository.SupplyRepository
}

// NewWarehouseService will create an object that represent the warehouse.Service interface
func NewWarehouseService(warehouse_postgres_repository *warehouse_postgres_repository.WarehouseRepository, supply_postgres_repository *supply_postgres_repository.SupplyRepository) *WarehouseService {
	return &WarehouseService{
		WarehousePostgresRepository: warehouse_postgres_repository,
		SupplyPostgresRepository:    supply_postgres_repository,
	}
}
