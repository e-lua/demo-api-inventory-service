package kardex_supply

import (
	kardex_postgres_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/kardex_supply"
)

type KardexSupplyService struct {
	KSPostgresRepository *kardex_postgres_repository.KardexSupplyRepository
}

// NewKardexSupplyService will create an object that represent the kardex_supply.Service interface
func NewKardexSupplyService(kardex_postgres_repository *kardex_postgres_repository.KardexSupplyRepository) *KardexSupplyService {
	return &KardexSupplyService{
		KSPostgresRepository: kardex_postgres_repository,
	}
}
