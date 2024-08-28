package provider

import (
	"errors"

	provider_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/provider"
)

func (ps *ProviderService) Update(input_idprovider string, input_provider *provider_model.Provider) (int, error) {

	//Validation of the Business Rules
	error_valid, is_valid := input_provider.IsValid()
	if !is_valid {
		return 4052, error_valid
	}

	//Update the Provider
	if input_provider.LegalData == nil {
		input_provider.LegalData = &provider_model.LegalData{}
	}
	if input_provider.AddressData == nil {
		input_provider.AddressData = &provider_model.AddressData{}
	}
	if input_provider.ContactData == nil {
		input_provider.ContactData = &provider_model.ContactData{}
	}

	updated_provider := provider_model.NewProvider(input_idprovider, input_provider.IdBusiness, input_provider.Name, input_provider.Description, input_provider.LegalData, input_provider.AddressData, input_provider.ContactData, &provider_model.DeletedData{})

	error_update_provider := ps.ProviderPostgresRepository.UpdateOne(updated_provider)
	if error_update_provider != nil {
		return 5057, errors.New("error update provider, details: " + error_update_provider.Error())
	}

	//OK
	return 0, nil
}
