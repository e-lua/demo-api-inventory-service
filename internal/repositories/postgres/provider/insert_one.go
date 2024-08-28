package provider

import (
	"context"
	"time"

	provider_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/provider"
)

func (pr *ProviderRepository) InsertOne(input_provider *provider_model.Provider) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := pr.ConnMasterPostgres

	query := `INSERT INTO Provider (id,id_business,name,description,legal_data,address_data,contact_data,deleted_data,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	_, err_query := db.Exec(ctx, query, input_provider.Id, input_provider.IdBusiness, input_provider.Name, input_provider.Description, input_provider.LegalData, input_provider.AddressData, input_provider.ContactData, input_provider.DeletedData, input_provider.CreatedAt, input_provider.UpdatedAt)

	if err_query != nil {
		return err_query
	}

	return nil
}
