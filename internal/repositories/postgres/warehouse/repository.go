package warehouse

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type WarehouseRepository struct {
	ConnMasterPostgres *pgxpool.Pool
}

// NewWarehouseRepository will create an object that represent the warehouse.Repository interface
func NewWarehouseRepository(conn_master_postgres *pgxpool.Pool) *WarehouseRepository {
	return &WarehouseRepository{
		ConnMasterPostgres: conn_master_postgres,
	}
}
