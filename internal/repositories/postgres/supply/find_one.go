package supply

import (
	"context"
	"fmt"
	"strings"
	"time"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (sr *SupplyRepository) FindOne(input_id string, input_idbusiness string, input_sku string) (*supply_model.Supply, error) {

	//Initialization
	oSupply := &supply_model.Supply{}

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_id != "" {
		filters["id"] = input_id
		counter_filters += 1
	}
	if input_idbusiness != "" {
		filters["id_business"] = input_idbusiness
		counter_filters += 1
	}
	if input_sku != "" {
		filters["sku"] = input_sku
		counter_filters += 1
	}

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := sr.ConnMasterPostgres

	//Define the query
	q := `SELECT id,id_business,sku,name,description,measure_data,warehouse_data,provider_data,deleted_data,created_at,updated_at FROM Supply`
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
		return oSupply, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		rows.Scan(&oSupply.Id, &oSupply.IdBusiness, &oSupply.SKU, &oSupply.Name, &oSupply.Description, &oSupply.MeasureData, &oSupply.WarehouseData, &oSupply.ProviderData, &oSupply.DeletedData, &oSupply.CreatedAt, &oSupply.UpdatedAt)
	}

	//Return the supply
	return oSupply, nil
}
