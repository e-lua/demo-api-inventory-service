package provider

import (
	"context"
	"time"

	provider_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/provider"
	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (pr *ProviderRepository) UpdateOne(input_provider *provider_model.Provider) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := pr.ConnMasterPostgres

	//BEGIN
	tx, error_tx := db.Begin(ctx)
	if error_tx != nil {
		tx.Rollback(ctx)
		return error_tx
	}

	//UPDATE PROVIDER
	query_provider := `UPDATE Provider SET name=$1,description=$2,legal_data=$3,address_data=$4,contact_data=$5,deleted_data=$6,updated_at=$7,updated_etl=$8,loaded_etl=$9 WHERE id=$10 AND id_business=$11`
	if _, err := tx.Exec(ctx, query_provider, input_provider.Name, input_provider.Description, input_provider.LegalData, input_provider.AddressData, input_provider.ContactData, input_provider.DeletedData, input_provider.UpdatedAt, true, false, input_provider.Id, input_provider.IdBusiness); err != nil {
		tx.Rollback(ctx)
		return err
	}

	//UPDATE PROVIDER'S NAME IN SUPPLY
	query_supply := `UPDATE Supply SET provider_data=$1,updated_etl=$2,loaded_etl=$3 WHERE provider_data->>'id'=$4`
	if _, err := tx.Exec(ctx, query_supply, &supply_model.GenericData{Id: input_provider.Id, Name: input_provider.Name}, true, false, input_provider.Id); err != nil {
		tx.Rollback(ctx)
		return err
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return err_commit
	}

	return nil
}
