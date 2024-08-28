package supply

import (
	"context"
	"time"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (sr *SupplyRepository) UpdateOne(input_supply *supply_model.Supply) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := sr.ConnMasterPostgres

	query := `UPDATE Supply SET sku=$1,name=$2,description=$3,measure_data=$4,warehouse_data=$5,provider_data=$6,deleted_data=$7,updated_at=$8,updated_etl=$9,loaded_etl=$10 WHERE id=$11 AND id_business=$12`
	_, err_query := db.Exec(ctx, query, input_supply.SKU, input_supply.Name, input_supply.Description, input_supply.MeasureData, input_supply.WarehouseData, input_supply.ProviderData, input_supply.DeletedData, input_supply.UpdatedAt, true, false, input_supply.Id, input_supply.IdBusiness)

	if err_query != nil {
		return err_query
	}

	return nil
}
