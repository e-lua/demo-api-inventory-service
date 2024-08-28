package warehouse

import (
	"errors"

	warehouse_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/warehouse"
)

func (ws *WarehouseService) GetTrash(input_idbusiness string, input_name string, input_limit int, input_offset int) (int, []*warehouse_model.Warehouse, error) {

	//Get the all warehouses
	list_warehouses, error_find_warehouses := ws.WarehousePostgresRepository.FindMany(input_idbusiness, "", "false", "true", input_limit, input_offset)
	if error_find_warehouses != nil {
		return 5057, []*warehouse_model.Warehouse{}, errors.New("error find warehouses, details: " + error_find_warehouses.Error())
	}

	//OK
	return 0, list_warehouses, nil
}
