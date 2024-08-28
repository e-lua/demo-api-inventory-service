package supply

import (
	"context"
	"time"

	supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/supply"
)

func (sr *SupplyRepository) FindOutOfStock() ([]*supply_model.SupplyOutOfStock, error) {

	//Initialization
	var oListSupplyOutOfStock []*supply_model.SupplyOutOfStock

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := sr.ConnMasterPostgres

	//Define the query
	q := `WITH supply AS 
(SELECT su.id_business AS id_business,su.id AS id,COALESCE(sum(CASE WHEN ks.id_type = 1 THEN ks.quantity WHEN ks.id_type = 2 THEN -ks.quantity ELSE 0 END),0) AS stock FROM Supply AS su LEFT JOIN KardexSupply AS ks ON su.id=ks.id_supply
WHERE deleted_data->>'isSendedToDelete'='false'
GROUP BY su.id_business,su.id
HAVING COALESCE(sum(CASE WHEN ks.id_type = 1 THEN ks.quantity WHEN ks.id_type = 2 THEN -ks.quantity ELSE 0 END),0) <= 0)
SELECT supply.id_business,COUNT(supply.id) FROM supply GROUP BY supply.id_business`

	rows, error_find := db.Query(ctx, q)
	if error_find != nil {
		return oListSupplyOutOfStock, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		oSupplyOutOfStock := &supply_model.SupplyOutOfStock{}
		rows.Scan(&oSupplyOutOfStock.IdBusiness, &oSupplyOutOfStock.Quantity)
		oListSupplyOutOfStock = append(oListSupplyOutOfStock, oSupplyOutOfStock)
	}

	//Return the list of supply out of stock
	return oListSupplyOutOfStock, nil
}
