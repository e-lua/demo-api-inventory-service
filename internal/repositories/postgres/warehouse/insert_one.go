package warehouse

import (
	"context"
	"time"

	warehouse_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/warehouse"
)

func (wr *WarehouseRepository) InsertOne(input_warehouse *warehouse_model.Warehouse) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := wr.ConnMasterPostgres

	query := `INSERT INTO Warehouse (id,id_business,name,description,deleted_data,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err_query := db.Exec(ctx, query, input_warehouse.Id, input_warehouse.IdBusiness, input_warehouse.Name, input_warehouse.Description, input_warehouse.DeletedData, input_warehouse.CreatedAt, input_warehouse.UpdatedAt)

	if err_query != nil {
		return err_query
	}

	return nil
}
