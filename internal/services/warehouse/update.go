package warehouse

import (
	"errors"

	warehouse_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/warehouse"
)

func (ws *WarehouseService) Update(input_idwarehouse string, input_warehouse *warehouse_model.Warehouse) (int, error) {

	//Validation of the Business Rules
	if input_idwarehouse == "" {
		return 4052, errors.New("id must be sent")
	}

	//Validation of the Business Rules
	error_valid, is_valid := input_warehouse.IsValid()
	if !is_valid {
		return 4052, error_valid
	}

	//Update the Warehouse
	updated_warehouse := warehouse_model.NewWarehouse(input_idwarehouse, input_warehouse.IdBusiness, input_warehouse.Name, input_warehouse.Description, &warehouse_model.DeletedData{})

	error_update_warehouse := ws.WarehousePostgresRepository.UpdateOne(updated_warehouse)
	if error_update_warehouse != nil {
		return 5057, errors.New("error update warehouse, details: " + error_update_warehouse.Error())
	}

	//OK
	return 0, nil
}
