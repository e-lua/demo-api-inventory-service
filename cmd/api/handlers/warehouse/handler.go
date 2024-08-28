package warehouse

import (
	warehouse_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/warehouse"
)

type WarehouseHandler struct {
	WarehouseService *warehouse_service.WarehouseService
}

// NewWarehouseHandler will create an object that represent the warehouse.Handler interface
func NewWarehouseHandler(warehouse_services *warehouse_service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{
		WarehouseService: warehouse_services,
	}
}
