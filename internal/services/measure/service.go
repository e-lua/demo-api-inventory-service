package measure

import (
	measure_postgres_repository "github.com/e-lua/demo-api-inventory-clean-architecture/internal/repositories/postgres/measure"
)

type MeasureService struct {
	MeasurePostgresRepository *measure_postgres_repository.MeasureRepository
}

// NewMeasureService will create an object that represent the measure.Service interface
func NewMeasureService(measure_postgres_repository *measure_postgres_repository.MeasureRepository) *MeasureService {
	return &MeasureService{
		MeasurePostgresRepository: measure_postgres_repository,
	}
}
