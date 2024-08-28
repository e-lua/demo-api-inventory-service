package kardex_supply

import (
	"context"
	"time"
)

func (ksr *KardexSupplyRepository) DeleteOne(input_idkadex string) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := ksr.ConnMasterPostgres

	query := `DELETE FROM KardexSupply WHERE id=$1`
	_, err_query := db.Exec(ctx, query, input_idkadex)

	if err_query != nil {
		return err_query
	}

	return nil
}
