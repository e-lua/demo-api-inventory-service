package provider

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func (pr *ProviderRepository) FindQuantity(input_idbusiness string) (int, error) {

	//Initialization
	var quantity int

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_idbusiness != "" {
		filters["id_business"] = input_idbusiness
		counter_filters += 1
	}
	filters["deleted_data->>'isDeleted'"] = "false"
	counter_filters += 1

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := pr.ConnMasterPostgres

	//Define the query
	q := `SELECT COUNT(*) FROM Provider `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}

	rows, error_find := db.Query(ctx, q)
	if error_find != nil {
		return quantity, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		rows.Scan(&quantity)
	}

	//Return the quantity
	return quantity, nil
}
