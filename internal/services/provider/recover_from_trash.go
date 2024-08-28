package provider

import (
	"errors"

	provider_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/provider"
)

func (ps *ProviderService) RecoverFromTrash(inout_idprovider string) (int, error) {

	//Search the provider
	provider_found, error_find_provider := ps.ProviderPostgresRepository.FindOne(inout_idprovider)
	if error_find_provider != nil {
		return 5057, error_find_provider
	}
	if provider_found.Id == "" {
		return 4055, errors.New("this provider does not exists")
	}

	//Update the Provider
	updated_supply := provider_model.NewProvider(provider_found.Id, provider_found.IdBusiness, provider_found.Name, provider_found.Description, provider_found.LegalData, provider_found.AddressData, provider_found.ContactData, &provider_model.DeletedData{})

	error_update_provider := ps.ProviderPostgresRepository.UpdateOne(updated_supply)
	if error_update_provider != nil {
		return 5057, errors.New("error update provider, details: " + error_update_provider.Error())
	}

	//OK
	return 0, nil
}
