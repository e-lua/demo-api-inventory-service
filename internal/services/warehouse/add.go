package warehouse

import (
	"errors"

	"github.com/google/uuid"

	warehouse_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/warehouse"
)

func (ws *WarehouseService) Add(input_warehouse *warehouse_model.Warehouse) (int, error) {

	//Validation of the Business Rules
	error_valid, is_valid := input_warehouse.IsValid()
	if !is_valid {
		return 4052, error_valid
	}

	//Storage the New Warehouse
	new_warehouse := warehouse_model.NewWarehouse(uuid.New().String(), input_warehouse.IdBusiness, input_warehouse.Name, input_warehouse.Description, &warehouse_model.DeletedData{})
	error_create_warehouse := ws.WarehousePostgresRepository.InsertOne(new_warehouse)
	if error_create_warehouse != nil {
		return 5057, errors.New("error create warehouse, details:" + error_create_warehouse.Error())
	}

	//OK
	return 0, nil
}
