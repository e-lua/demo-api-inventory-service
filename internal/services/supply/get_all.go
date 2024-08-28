package supply

import (
	"errors"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (ss *SupplyService) GetAll(input_idbusiness string, input_sku string, input_name string, input_idwarehouse string, input_idprovider string, input_limit int, input_offset int) (int, []*supply_model.Supply, error) {

	//Get the all supplies
	list_supplies, error_find_supply := ss.SupplyPostgresRepository.FindMany(input_idwarehouse, input_idprovider, input_idbusiness, input_sku+input_name, "false", "false", input_limit, input_offset)
	if error_find_supply != nil {
		return 5057, []*supply_model.Supply{}, errors.New("error find supply, details: " + error_find_supply.Error())
	}

	//OK
	return 0, list_supplies, nil
}
