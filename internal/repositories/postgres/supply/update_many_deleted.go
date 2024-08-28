package supply

import (
	"context"
	"time"
)

func (sr *SupplyRepository) UpdateManyDelete() error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := sr.ConnMasterPostgres

	query := `UPDATE Supply SET sku=concat(sku,'-',id),deleted_data = deleted_data || '{"isDeleted": "true"}'  WHERE deleted_data->'isSendedToDelete'='true' AND (deleted_data->>'sendedToDeleteAt')::timestamp<=now()`
	_, err_query := db.Exec(ctx, query)

	if err_query != nil {
		return err_query
	}

	return nil
}
