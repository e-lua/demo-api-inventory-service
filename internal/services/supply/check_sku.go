package supply

import "errors"

func (ss *SupplyService) CheckSku(input_sku string) (int, error) {

	//Get the one supply
	supply, error_find_supply := ss.SupplyPostgresRepository.FindOne("", "", input_sku)
	if error_find_supply != nil {
		return 5057, errors.New("error find supply, details: " + error_find_supply.Error())
	}
	if supply.Id != "" {
		return 4052, errors.New("sku already exists")
	}

	//OK
	return 0, nil
}
