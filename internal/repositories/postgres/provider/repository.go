package provider

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProviderRepository struct {
	ConnMasterPostgres *pgxpool.Pool
}

// NewMeasureRepository will create an object that represent the provider.Repository interface
func NewMeasureRepository(conn_master_postgres *pgxpool.Pool) *ProviderRepository {
	return &ProviderRepository{
		ConnMasterPostgres: conn_master_postgres,
	}
}
