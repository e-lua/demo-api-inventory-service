package supply

import (
	"context"
	"fmt"
	"strings"
	"time"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (sr *SupplyRepository) FindMany(input_warehouse string, input_provider string, input_idbusiness string, input_search_text string, input_isdelete string, input_issendedtodelete string, input_limit int, input_offset int) ([]*supply_model.Supply, error) {

	//Initialization
	var oListSupply []*supply_model.Supply

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_idbusiness != "" {
		filters["su.id_business"] = input_idbusiness
		counter_filters += 1
	}
	if input_search_text != "" {
		filters["su.idx_sku_name"] = input_search_text
		counter_filters += 1
	}
	if input_warehouse != "" {
		filters["su.warehouse_data->>'id'"] = input_warehouse
		counter_filters += 1
	}
	if input_provider != "" {
		filters["su.provider_data->>'id'"] = input_provider
		counter_filters += 1
	}
	if input_isdelete != "" {
		filters["su.deleted_data->>'isDeleted'"] = input_isdelete
		counter_filters += 1
	}
	if input_issendedtodelete != "" {
		filters["su.deleted_data->>'isSendedToDelete'"] = input_issendedtodelete
		counter_filters += 1
	}

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := sr.ConnMasterPostgres

	//Define the query
	q := `SELECT su.id,su.id_business,su.sku,su.name,COALESCE(sum(CASE WHEN ks.id_type = 1 THEN ks.quantity WHEN ks.id_type = 2 THEN -ks.quantity ELSE 0 END),0) AS stock,su.description,su.measure_data,su.warehouse_data,su.provider_data,su.deleted_data,su.created_at,su.updated_at FROM Supply AS su LEFT JOIN KardexSupply AS ks ON su.id=ks.id_supply`
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			if key == "su.idx_sku_name" {
				clausulas = append(clausulas, fmt.Sprintf("%s @@ to_tsquery('spanish', '%s:*')", key, value))
			} else {
				clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
			}
		}
		q += strings.Join(clausulas, " AND ")

	}

	rows, error_find := db.Query(ctx, q+" GROUP BY su.id,su.id_business,su.sku,su.name,su.description,su.measure_data,su.warehouse_data,su.provider_data,su.deleted_data,su.created_at,su.updated_at ORDER BY stock ASC LIMIT $1 OFFSET $2", input_limit, input_offset)
	if error_find != nil {
		return oListSupply, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		oSupply := &supply_model.Supply{}
		rows.Scan(&oSupply.Id, &oSupply.IdBusiness, &oSupply.SKU, &oSupply.Name, &oSupply.Stock, &oSupply.Description, &oSupply.MeasureData, &oSupply.WarehouseData, &oSupply.ProviderData, &oSupply.DeletedData, &oSupply.CreatedAt, &oSupply.UpdatedAt)
		oListSupply = append(oListSupply, oSupply)
	}

	//Return the list of supply
	return oListSupply, nil
}
