package measure

import (
	"context"
	"time"
)

func (mr *MeasureRepository) DeleteOne(input_idmeasure string) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := mr.ConnMasterPostgres

	query := `DELETE FROM Measure WHERE id=$1`
	_, err_query := db.Exec(ctx, query, input_idmeasure)

	if err_query != nil {
		return err_query
	}

	return nil
}
