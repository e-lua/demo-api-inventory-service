package supply

import (
	"errors"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (ss *SupplyService) GetTrash(intput_idbusiness string, input_limit int, input_offset int) (int, []*supply_model.Supply, error) {

	//Get the all supplies
	list_supplies, error_find_supply := ss.SupplyPostgresRepository.FindMany("", "", intput_idbusiness, "", "false", "true", input_limit, input_offset)
	if error_find_supply != nil {
		return 5057, []*supply_model.Supply{}, errors.New("error find supply, details: " + error_find_supply.Error())
	}

	//OK
	return 0, list_supplies, nil
}
