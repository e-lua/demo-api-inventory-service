package supply

import (
	"errors"

	"github.com/google/uuid"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (ss *SupplyService) Add(input_supply *supply_model.Supply) (int, error) {

	//Validation of the Business Rules
	error_valid, is_valid := input_supply.IsValid()
	if !is_valid {
		return 4052, error_valid
	}

	//Check if the SKU already exists
	suppy_found, error_find_supply := ss.SupplyPostgresRepository.FindOne("", "", input_supply.SKU)
	if error_find_supply != nil {
		return 5057, errors.New("error find supply, details: " + error_find_supply.Error())
	}
	//if already exists any supply with this SKU
	if suppy_found.SKU != "" {
		return 4052, errors.New("this SKU already exists")
	}

	//Storage the New Supply
	if input_supply.ProviderData == nil {
		input_supply.ProviderData = &supply_model.GenericData{}
	}
	if input_supply.WarehouseData == nil {
		input_supply.WarehouseData = &supply_model.GenericData{}
	}
	if input_supply.MeasureData == nil {
		input_supply.MeasureData = &supply_model.GenericData{}
	}

	new_supply := supply_model.NewSupply(uuid.New().String(), input_supply.IdBusiness, input_supply.SKU, input_supply.Name, input_supply.Description, input_supply.MeasureData, input_supply.WarehouseData, input_supply.ProviderData, &supply_model.DeletedData{})
	error_create_supply := ss.SupplyPostgresRepository.InsertOne(new_supply)
	if error_create_supply != nil {
		return 5057, errors.New("error create supply, details: " + error_create_supply.Error())
	}

	//OK
	return 0, nil
}