package provider

import (
	"errors"
	"strconv"
	"time"

	provider_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/provider"
)

func (ps *ProviderService) SendToDelete(intput_idprovider string) (int, error) {

	//Check if there are supplies with this provider
	list_supplies, error_find_supplies := ps.SupplyPostgresRepository.FindMany("", intput_idprovider, "", "", "false", "false", 10, 0)
	if error_find_supplies != nil {
		return 5057, errors.New("error find supplies, details: " + error_find_supplies.Error())
	}
	if len(list_supplies) > 0 {
		return 4020, errors.New("this provider has " + strconv.Itoa(len(list_supplies)) + " supplies")
	}

	//Search the provider
	provider_found, error_find_provider := ps.ProviderPostgresRepository.FindOne(intput_idprovider)
	if error_find_provider != nil {
		return 5057, errors.New("error find provider, details: " + error_find_provider.Error())
	}
	if provider_found.Id == "" {
		return 4055, errors.New("this provider does not exists")
	}

	//Update the Provider
	if provider_found.LegalData == nil {
		provider_found.LegalData = &provider_model.LegalData{}
	}
	if provider_found.AddressData == nil {
		provider_found.AddressData = &provider_model.AddressData{}
	}
	if provider_found.ContactData == nil {
		provider_found.ContactData = &provider_model.ContactData{}
	}

	updated_provider := provider_model.NewProvider(provider_found.Id, provider_found.IdBusiness, provider_found.Name, provider_found.Description, provider_found.LegalData, provider_found.AddressData, provider_found.ContactData, &provider_model.DeletedData{Is_deleted: false, Is_sended_to_delete: true, Sended_to_delete_at: time.Now(), Deleted_at: time.Now().Add(168 * time.Hour)})

	error_update_provider := ps.ProviderPostgresRepository.UpdateOne(updated_provider)
	if error_update_provider != nil {
		return 5057, error_update_provider
	}

	//OK
	return 0, nil
}
