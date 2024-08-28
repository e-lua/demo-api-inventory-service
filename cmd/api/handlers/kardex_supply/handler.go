package provider

import (
	kardex_supply_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/kardex_supply"
)

type KardexSupplyHandler struct {
	KardexSupplyService *kardex_supply_service.KardexSupplyService
}

// NewKardexSupplyHandler will create an object that represent the supply.Handler interface
func NewKardexSupplyHandler(kardex_supply_service *kardex_supply_service.KardexSupplyService) *KardexSupplyHandler {
	return &KardexSupplyHandler{
		KardexSupplyService: kardex_supply_service,
	}
}
