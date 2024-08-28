package kardex_supply

import (
	"context"
	"fmt"
	"strings"
	"time"

	kardex_supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/kardex_supply"
)

func (ksr *KardexSupplyRepository) FindMany(input_idsupply string, input_idbusiness string, input_type int, input_idcategory int, input_limit int, input_offset int) ([]*kardex_supply_model.KardexSupply, error) {

	//Initialization
	var oListKardexSupply []*kardex_supply_model.KardexSupply

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_idsupply != "" {
		filters["id_supply"] = input_idsupply
		counter_filters += 1
	}
	if input_idbusiness != "" {
		filters["id_business"] = input_idbusiness
		counter_filters += 1
	}
	if input_type != 0 {
		filters["id_type"] = input_type
		counter_filters += 1
	}
	if input_idcategory != 0 {
		filters["id_category"] = input_idcategory
		counter_filters += 1
	}

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := ksr.ConnMasterPostgres

	//Define the query
	q := `SELECT id,id_business,id_supply,date,id_type,id_category,quantity,total_cost,unit_cost,updated_by,created_at,updated_at FROM KardexSupply `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			if key == "id_type" || key == "id_category" {
				clausulas = append(clausulas, fmt.Sprintf("%s = %d", key, value))
			} else {
				clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
			}
		}
		q += strings.Join(clausulas, " AND ")

	}

	rows, error_find := db.Query(ctx, q+" ORDER BY date DESC LIMIT $1 OFFSET $2", input_limit, input_offset)
	if error_find != nil {
		return oListKardexSupply, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		oKardexSupply := &kardex_supply_model.KardexSupply{}
		rows.Scan(&oKardexSupply.Id, &oKardexSupply.IdBusiness, &oKardexSupply.IdSupply, &oKardexSupply.Date, &oKardexSupply.IdType, &oKardexSupply.IdCategory, &oKardexSupply.Quantity, &oKardexSupply.TotalCost, &oKardexSupply.UnitCost, &oKardexSupply.UpdatedBy, &oKardexSupply.CreatedAt, &oKardexSupply.UpdatedAt)
		oListKardexSupply = append(oListKardexSupply, oKardexSupply)
	}

	//Return the list of kardex_supply
	return oListKardexSupply, nil
}
