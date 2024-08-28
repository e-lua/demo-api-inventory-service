package kardex_supply

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type KardexSupplyRepository struct {
	ConnMasterPostgres *pgxpool.Pool
}

// NewKardexSupplyRepository will create an object that represent the kardex_supply.Repository interface
func NewKardexSupplyRepository(conn_master_postgres *pgxpool.Pool) *KardexSupplyRepository {
	return &KardexSupplyRepository{
		ConnMasterPostgres: conn_master_postgres,
	}
}
