package kardex_supply

import (
	"context"
	"time"

	kardex_supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/kardex_supply"
)

func (ksr *KardexSupplyRepository) InsertOne(input_kadex *kardex_supply_model.KardexSupply) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := ksr.ConnMasterPostgres

	query := `INSERT INTO KardexSupply (id,id_business,id_supply,date,id_type,id_category,quantity,total_cost,updated_by,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	_, err_query := db.Exec(ctx, query, input_kadex.Id, input_kadex.IdBusiness, input_kadex.IdSupply, input_kadex.Date, input_kadex.IdType, input_kadex.IdCategory, input_kadex.Quantity, input_kadex.TotalCost, input_kadex.UpdatedBy, input_kadex.CreatedAt, input_kadex.UpdatedAt)

	if err_query != nil {
		return err_query
	}

	return nil
}
