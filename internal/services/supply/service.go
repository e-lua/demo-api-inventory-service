package supply

import (
	supply_postgres_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/supply"
)

type SupplyService struct {
	SupplyPostgresRepository *supply_postgres_repository.SupplyRepository
}

// NewSupplyService will create an object that represent the supply.Service interface
func NewSupplyService(supply_postgres_repository *supply_postgres_repository.SupplyRepository) *SupplyService {
	return &SupplyService{
		SupplyPostgresRepository: supply_postgres_repository,
	}
}
