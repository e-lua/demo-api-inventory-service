package measure

import (
	"context"
	"fmt"
	"strings"
	"time"

	measure_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/measure"
)

func (mr *MeasureRepository) FindMany(input_idbusiness string, input_search_text string, input_limit int, input_offset int) ([]*measure_model.Measure, error) {

	//Initialization
	var oListMeasure []*measure_model.Measure

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_idbusiness != "" {
		filters["id_business"] = input_idbusiness
		counter_filters += 1
	}
	if input_search_text != "" {
		filters["idx_name"] = input_search_text
		counter_filters += 1
	}

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := mr.ConnMasterPostgres

	//Define the query
	q := `SELECT id,id_business,name FROM Measure`
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			if key == "idx_name" {
				clausulas = append(clausulas, fmt.Sprintf("%s @@ to_tsquery('spanish', '%s:*')", key, value))
			} else {
				clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
			}
		}
		q += strings.Join(clausulas, " AND ")

	}

	rows, error_find := db.Query(ctx, q+" ORDER BY name ASC LIMIT $1 OFFSET $2", input_limit, input_offset)
	if error_find != nil {
		return oListMeasure, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		oMeasure := &measure_model.Measure{}
		rows.Scan(&oMeasure.Id, &oMeasure.IdBusiness, &oMeasure.Name)
		oListMeasure = append(oListMeasure, oMeasure)
	}

	//Return the list of measure
	return oListMeasure, nil
}
