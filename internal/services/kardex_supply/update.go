package kardex_supply

import (
	"errors"

	kardex_supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/kardex_supply"
)

func (kss *KardexSupplyService) Update(input_idkardex string, input_fullname string, input_kardex *kardex_supply_model.KardexSupply) (int, error) {

	//Validation of the Business Rules
	error_valid, is_valid := input_kardex.IsValid()
	if !is_valid {
		return 4052, error_valid
	}

	//Update the Kardex
	updated_kardex := kardex_supply_model.NewKardexSupply(input_idkardex, input_kardex.IdBusiness, input_kardex.IdSupply, input_kardex.Date, input_kardex.IdType, input_kardex.IdCategory, input_kardex.Quantity, input_kardex.TotalCost, input_fullname)

	error_update_kardex := kss.KSPostgresRepository.UpdateOne(updated_kardex)
	if error_update_kardex != nil {
		return 5057, errors.New("error update kardex, details: " + error_update_kardex.Error())
	}

	//OK
	return 0, nil
}
