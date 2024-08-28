package supply

import (
	supply_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/supply"
)

type SupplyHandler struct {
	SupplyService *supply_service.SupplyService
}

// NewSupplyHandler will create an object that represent the supply.Handler interface
func NewSupplyHandler(supply_services *supply_service.SupplyService) *SupplyHandler {
	return &SupplyHandler{
		SupplyService: supply_services,
	}
}
