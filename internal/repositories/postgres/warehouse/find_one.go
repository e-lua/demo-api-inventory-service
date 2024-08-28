package warehouse

import (
	"context"
	"fmt"
	"strings"
	"time"

	warehouse_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/warehouse"
)

func (wr *WarehouseRepository) FindOne(input_id string) (warehouse_model.Warehouse, error) {

	//Initialization
	var oWarehouse warehouse_model.Warehouse

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_id != "" {
		filters["id"] = input_id
		counter_filters += 1
	}

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := wr.ConnMasterPostgres

	//Define the query
	q := `SELECT id,id_business,name,description,deleted_data,created_at,updated_at FROM Warehouse `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}

	rows, error_find := db.Query(ctx, q)
	if error_find != nil {
		return oWarehouse, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		rows.Scan(&oWarehouse.Id, &oWarehouse.IdBusiness, &oWarehouse.Name, &oWarehouse.Description, &oWarehouse.DeletedData, &oWarehouse.CreatedAt, &oWarehouse.UpdatedAt)
	}

	//Return the warehouse
	return oWarehouse, nil
}
