package measure

import (
	"context"
	"time"

	measure_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/measure"
)

func (mr *MeasureRepository) InsertOne(input_measure *measure_model.Measure) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	//Start the connection
	db := mr.ConnMasterPostgres

	query := `INSERT INTO Measure (id,id_business,name) VALUES ($1,$2,$3)`
	_, err_query := db.Exec(ctx, query, input_measure.Id, input_measure.IdBusiness, input_measure.Name)

	if err_query != nil {
		return err_query
	}

	return nil
}
