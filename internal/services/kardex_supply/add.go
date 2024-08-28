package kardex_supply

import (
	"errors"

	"github.com/google/uuid"

	kardex_supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/kardex_supply"
)

func (kss *KardexSupplyService) Add(intput_fullname string, input_kardex *kardex_supply_model.KardexSupply) (int, error) {

	//Validation of the Business Rules
	error_valid, is_valid := input_kardex.IsValid()
	if !is_valid {
		return 4052, error_valid
	}

	//Storage the New Kardex
	new_kardex := kardex_supply_model.NewKardexSupply(uuid.New().String(), input_kardex.IdBusiness, input_kardex.IdSupply, input_kardex.Date, input_kardex.IdType, input_kardex.IdCategory, input_kardex.Quantity, input_kardex.TotalCost, intput_fullname)
	error_create_kardex := kss.KSPostgresRepository.InsertOne(new_kardex)
	if error_create_kardex != nil {
		return 5057, errors.New("error ad kardex, details " + error_create_kardex.Error())
	}

	//OK
	return 0, nil
}
