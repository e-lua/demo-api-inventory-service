package provider

import (
	measure_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/measure"
)

type MeasureHandler struct {
	MeasureService *measure_service.MeasureService
}

// NewMeasureHandler will create an object that represent the supply.Handler interface
func NewMeasureHandler(measure_service *measure_service.MeasureService) *MeasureHandler {
	return &MeasureHandler{
		MeasureService: measure_service,
	}
}
