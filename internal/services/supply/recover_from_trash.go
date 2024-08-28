package supply

import (
	"errors"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (ss *SupplyService) RecoverFromTrash(input_idsupply string) (int, error) {

	//Search the supply
	supply_found, error_find_supply := ss.SupplyPostgresRepository.FindOne(input_idsupply, "", "")
	if error_find_supply != nil {
		return 5057, errors.New("error find supply, details: " + error_find_supply.Error())
	}
	if supply_found.Id == "" {
		return 4055, errors.New("this suspply does not exits")
	}

	//Update the Supply
	supply_found.DeletedData.Is_sended_to_delete = false

	udpated_supply := supply_model.NewSupply(supply_found.Id, supply_found.IdBusiness, supply_found.SKU, supply_found.Name, supply_found.Description, supply_found.MeasureData, supply_found.WarehouseData, supply_found.ProviderData, &supply_model.DeletedData{})

	error_update_supply := ss.SupplyPostgresRepository.UpdateOne(udpated_supply)
	if error_update_supply != nil {
		return 5057, errors.New("error update supply, details: " + error_update_supply.Error())
	}

	//OK
	return 0, nil
}
