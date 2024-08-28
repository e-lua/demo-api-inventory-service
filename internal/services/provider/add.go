package provider

import (
	"errors"

	"github.com/google/uuid"

	provider_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/provider"
)

func (ps *ProviderService) Add(input_provider *provider_model.Provider) (int, error) {

	//Validation of the Business Rules
	error_valid, is_valid := input_provider.IsValid()
	if !is_valid {
		return 4052, error_valid
	}

	//Storage the New Provider
	if input_provider.LegalData == nil {
		input_provider.LegalData = &provider_model.LegalData{}
	}
	if input_provider.AddressData == nil {
		input_provider.AddressData = &provider_model.AddressData{}
	}
	if input_provider.ContactData == nil {
		input_provider.ContactData = &provider_model.ContactData{}
	}

	new_provider := provider_model.NewProvider(uuid.New().String(), input_provider.IdBusiness, input_provider.Name, input_provider.Description, input_provider.LegalData, input_provider.AddressData, input_provider.ContactData, &provider_model.DeletedData{})
	error_create_provider := ps.ProviderPostgresRepository.InsertOne(new_provider)
	if error_create_provider != nil {
		return 5057, errors.New("error create provider, details: " + error_create_provider.Error())
	}

	//OK
	return 0, nil
}
