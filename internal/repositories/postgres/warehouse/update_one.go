package warehouse

import (
	"context"
	"time"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
	warehouse_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/warehouse"
)

func (wr *WarehouseRepository) UpdateOne(input_warehouse *warehouse_model.Warehouse) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := wr.ConnMasterPostgres

	//BEGIN
	tx, error_tx := db.Begin(ctx)
	if error_tx != nil {
		tx.Rollback(ctx)
		return error_tx
	}

	//UPDATE WAREHOUSE
	query_warehouse := `UPDATE Warehouse SET name=$1,description=$2,deleted_data=$3,updated_at=$4,updated_etl=$5,loaded_etl=$6 WHERE id=$7 AND id_business=$8`
	if _, err := tx.Exec(ctx, query_warehouse, input_warehouse.Name, input_warehouse.Description, input_warehouse.DeletedData, input_warehouse.UpdatedAt, true, false, input_warehouse.Id, input_warehouse.IdBusiness); err != nil {
		tx.Rollback(ctx)
		return err
	}

	//UPDATE WAREHOUSE'NAME IN SUPPLY
	query_supply := `UPDATE Supply SET warehouse_data=$1,updated_etl=$2,loaded_etl=$3 WHERE warehouse_data->>'id'=$4`
	if _, err := tx.Exec(ctx, query_supply, &supply_model.GenericData{Id: input_warehouse.Id, Name: input_warehouse.Name}, true, false, input_warehouse.Id); err != nil {
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
