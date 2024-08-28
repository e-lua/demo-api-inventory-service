package warehouse

import (
	"context"
	"fmt"
	"strings"
	"time"

	warehouse_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/warehouse"
)

func (wr *WarehouseRepository) FindMany(input_idbusiness string, input_search_text string, input_isdelete string, input_issendedtodelete string, input_limit int, input_offset int) ([]*warehouse_model.Warehouse, error) {

	//Initialization
	var oListWarehouse []*warehouse_model.Warehouse

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_idbusiness != "" {
		filters["id_business"] = input_idbusiness
		counter_filters += 1
	}
	if input_search_text != "" {
		filters["idx_name_description"] = input_search_text
		counter_filters += 1
	}
	if input_isdelete != "" {
		filters["deleted_data->>'isDeleted'"] = input_isdelete
		counter_filters += 1
	}
	if input_issendedtodelete != "" {
		filters["deleted_data->>'isSendedToDelete'"] = input_issendedtodelete
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
			if key == "ware.idx_name_description" {
				clausulas = append(clausulas, fmt.Sprintf("%s @@ to_tsquery('spanish', '%s:*')", key, value))
			} else {
				clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
			}
		}
		q += strings.Join(clausulas, " AND ")

	}

	rows, error_find := db.Query(ctx, q+" ORDER BY name ASC LIMIT $1 OFFSET $2", input_limit, input_offset)
	if error_find != nil {
		return oListWarehouse, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		oWarehouse := &warehouse_model.Warehouse{}
		rows.Scan(&oWarehouse.Id, &oWarehouse.IdBusiness, &oWarehouse.Name, &oWarehouse.Description, &oWarehouse.DeletedData, &oWarehouse.CreatedAt, &oWarehouse.UpdatedAt)
		oListWarehouse = append(oListWarehouse, oWarehouse)
	}

	//Return the list of warehouse
	return oListWarehouse, nil
}
