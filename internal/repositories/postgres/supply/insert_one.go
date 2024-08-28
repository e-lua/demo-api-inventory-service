package supply

import (
	"context"
	"time"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (sr *SupplyRepository) InsertOne(input_supply *supply_model.Supply) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := sr.ConnMasterPostgres

	query := `INSERT INTO Supply (id,id_business,sku,name,description,measure_data,warehouse_data,provider_data,deleted_data,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	_, err_query := db.Exec(ctx, query, input_supply.Id, input_supply.IdBusiness, input_supply.SKU, input_supply.Name, input_supply.Description, input_supply.MeasureData, input_supply.WarehouseData, input_supply.ProviderData, input_supply.DeletedData, input_supply.CreatedAt, input_supply.UpdatedAt)

	if err_query != nil {
		return err_query
	}

	return nil
}
