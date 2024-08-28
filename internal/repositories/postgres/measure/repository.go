package measure

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type MeasureRepository struct {
	ConnMasterPostgres *pgxpool.Pool
}

// NewMeasureRepository will create an object that represent the measure.Repository interface
func NewMeasureRepository(conn_master_postgres *pgxpool.Pool) *MeasureRepository {
	return &MeasureRepository{
		ConnMasterPostgres: conn_master_postgres,
	}
}
