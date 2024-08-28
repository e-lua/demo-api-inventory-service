package provider

import (
	provider_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/provider"
)

type ProviderHandler struct {
	ProviderService *provider_service.ProviderService
}

// NewProviderHandler will create an object that represent the supply.Handler interface
func NewProviderHandler(provider_services *provider_service.ProviderService) *ProviderHandler {
	return &ProviderHandler{
		ProviderService: provider_services,
	}
}
