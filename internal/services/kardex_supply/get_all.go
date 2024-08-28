package kardex_supply

import (
	"errors"

	kardex_supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/kardex_supply"
)

func (kss *KardexSupplyService) GetAll(input_idsupply string, input_idbusiness string, input_type_mov int, input_category_mov int, input_limit int, input_offset int) (int, []*kardex_supply_model.KardexSupply, error) {

	//Get the all kardex
	list_kardex, error_find_kardex := kss.KSPostgresRepository.FindMany(input_idsupply, input_idbusiness, input_type_mov, input_category_mov, input_limit, input_offset)
	if error_find_kardex != nil {
		return 5057, []*kardex_supply_model.KardexSupply{}, errors.New("error fin kardex, details: " + error_find_kardex.Error())
	}

	//OK
	return 0, list_kardex, nil
}
