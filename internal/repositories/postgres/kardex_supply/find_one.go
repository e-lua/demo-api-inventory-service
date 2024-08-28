package kardex_supply

import (
	"context"
	"fmt"
	"strings"
	"time"

	kardex_supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/kardex_supply"
)

func (ksr *KardexSupplyRepository) FindOne(input_id string, input_idbusiness string) (*kardex_supply_model.KardexSupply, error) {

	//Initialization
	oKardexSupply := &kardex_supply_model.KardexSupply{}

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

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := ksr.ConnMasterPostgres

	//Define the query
	q := `SELECT id,id_business,id_supply,date,id_type,id_category,quantity,total_cost,updated_by,created_at,updated_at FROM KardexSupply`
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
		return oKardexSupply, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		rows.Scan(&oKardexSupply.Id, &oKardexSupply.IdBusiness, &oKardexSupply.IdSupply, &oKardexSupply.Date, &oKardexSupply.IdType, &oKardexSupply.IdCategory, &oKardexSupply.Quantity, &oKardexSupply.TotalCost, &oKardexSupply.UpdatedBy, &oKardexSupply.CreatedAt, &oKardexSupply.UpdatedAt)
	}

	//Return the provider
	return oKardexSupply, nil
}
