package warehouse

import (
	"errors"
	"strconv"
	"time"

	warehouse_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/warehouse"
)

func (ws *WarehouseService) SendToDelete(input_idwarehouse string) (int, error) {

	//Validation of the Business Rules
	if input_idwarehouse == "" {
		return 4052, errors.New("id must be sent")
	}

	//Check if there is supplies with this warehouse
	list_supplies, error_find_supply := ws.SupplyPostgresRepository.FindMany(input_idwarehouse, "", "", "", "false", "false", 10, 0)
	if error_find_supply != nil {
		return 5057, errors.New("error find supply, details: " + error_find_supply.Error())
	}
	if len(list_supplies) > 0 {
		return 40520, errors.New("this provider has " + strconv.Itoa(len(list_supplies)) + " supplies")
	}

	//Search the warehouse
	warehouse_found, error_find_warehouse := ws.WarehousePostgresRepository.FindOne(input_idwarehouse)
	if error_find_warehouse != nil {
		return 5057, errors.New("error find warehouse, details: " + error_find_warehouse.Error())
	}
	if warehouse_found.Id == "" {
		return 4055, errors.New("this warehouse does not exists")
	}

	//Update the Warehouse
	updated_warehouse := warehouse_model.NewWarehouse(warehouse_found.Id, warehouse_found.IdBusiness, warehouse_found.Name, warehouse_found.Description, &warehouse_model.DeletedData{Is_deleted: false, Is_sended_to_delete: true, Sended_to_delete_at: time.Now(), Deleted_at: time.Now().Add(168 * time.Hour)})

	error_update_warehouse := ws.WarehousePostgresRepository.UpdateOne(updated_warehouse)
	if error_update_warehouse != nil {
		return 5057, errors.New("error update warehouse, details: " + error_update_warehouse.Error())
	}

	//OK
	return 0, nil
}
