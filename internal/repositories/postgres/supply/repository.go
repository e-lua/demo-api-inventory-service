package supply

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type SupplyRepository struct {
	ConnMasterPostgres *pgxpool.Pool
}

// NewSupplyRepository will create an object that represent the supply.Repository interface
func NewSupplyRepository(conn_master_postgres *pgxpool.Pool) *SupplyRepository {
	return &SupplyRepository{
		ConnMasterPostgres: conn_master_postgres,
	}
}
