package supply

import (
	"errors"
	"time"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (ss *SupplyService) SendToDelete(input_idsupply string) (int, error) {

	//Search the supply
	supply_found, error_find_supply := ss.SupplyPostgresRepository.FindOne(input_idsupply, "", "")
	if error_find_supply != nil {
		return 5057, error_find_supply
	}
	if supply_found.Id == "" {
		return 4055, errors.New("this supply does not exists")
	}

	//Update the Supply
	if supply_found.ProviderData == nil {
		supply_found.ProviderData = &supply_model.GenericData{}
	}
	if supply_found.WarehouseData == nil {
		supply_found.WarehouseData = &supply_model.GenericData{}
	}
	if supply_found.MeasureData == nil {
		supply_found.MeasureData = &supply_model.GenericData{}
	}

	updated_supply := supply_model.NewSupply(supply_found.Id, supply_found.IdBusiness, supply_found.SKU, supply_found.Name, supply_found.Description, supply_found.MeasureData, supply_found.WarehouseData, supply_found.ProviderData, &supply_model.DeletedData{Is_deleted: false, Is_sended_to_delete: true, Sended_to_delete_at: time.Now(), Deleted_at: time.Now().Add(168 * time.Hour)})

	error_update_supply := ss.SupplyPostgresRepository.UpdateOne(updated_supply)
	if error_update_supply != nil {
		return 5057, errors.New("error update supply, details: " + error_update_supply.Error())
	}

	//OK
	return 0, nil
}
