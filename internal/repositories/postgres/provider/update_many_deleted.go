package provider

import (
	"context"
	"time"
)

func (pr *ProviderRepository) UpdateManyDelete() error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := pr.ConnMasterPostgres

	query := `UPDATE Provider SET deleted_data = deleted_data || '{"isDeleted": "true"}'  WHERE deleted_data->'isSendedToDelete'='true' AND (deleted_data->>'sendedToDeleteAt')::timestamp<=now()`
	_, err_query := db.Exec(ctx, query)

	if err_query != nil {
		return err_query
	}

	return nil
}
