package provider

import (
	"errors"

	provider_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/provider"
)

func (ps *ProviderService) GetTrash(input_idbusiness string, input_limit int, input_offset int) (int, []*provider_model.Provider, error) {

	//Get the all providers
	list_providers, error_find_provider := ps.ProviderPostgresRepository.FindMany(input_idbusiness, "", "false", "true", input_limit, input_offset)
	if error_find_provider != nil {
		return 5057, []*provider_model.Provider{}, errors.New("error find provider, details: " + error_find_provider.Error())
	}

	//OK
	return 0, list_providers, nil
}
