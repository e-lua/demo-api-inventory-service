package kardex_supply

import (
	"context"
	"time"

	kardex_supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/kardex_supply"
)

func (ksr *KardexSupplyRepository) UpdateOne(input_kardex *kardex_supply_model.KardexSupply) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := ksr.ConnMasterPostgres

	query := `UPDATE KardexSupply SET date=$1,id_type=$2,id_category=$3,quantity=$4,total_cost=$5,updated_by=$6,updated_at=$7,updated_etl=$8,loaded_etl=$9 WHERE id=$10 AND id_business=$11`
	_, err_query := db.Exec(ctx, query, input_kardex.Date, input_kardex.IdType, input_kardex.IdCategory, input_kardex.Quantity, input_kardex.TotalCost, input_kardex.UpdatedBy, input_kardex.UpdatedAt, true, false, input_kardex.Id, input_kardex.IdBusiness)

	if err_query != nil {
		return err_query
	}

	return nil
}
