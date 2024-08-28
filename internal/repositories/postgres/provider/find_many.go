package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	provider_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/provider"
)

func (pr *ProviderRepository) FindMany(input_idbusiness string, input_search_text string, input_isdelete string, input_issendedtodelete string, input_limit int, input_offset int) ([]*provider_model.Provider, error) {

	//Initialization
	var oListProvider []*provider_model.Provider

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_idbusiness != "" {
		filters["id_business"] = input_idbusiness
		counter_filters += 1
	}
	if input_search_text != "" {
		filters["idx_name_descripcion"] = input_search_text
		counter_filters += 1
	}
	if input_isdelete != "" {
		filters["deleted_data->>'isDeleted'"] = input_isdelete
		counter_filters += 1
	}
	if input_issendedtodelete != "" {
		filters["deleted_data->>'isSendedToDelete'"] = input_issendedtodelete
		counter_filters += 1
	}

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := pr.ConnMasterPostgres

	//Define the query
	q := `SELECT id,id_business,name,description,legal_data,address_data,contact_data,deleted_data,created_at,updated_at FROM Provider `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			if key == "idx_name_descripcion" {
				clausulas = append(clausulas, fmt.Sprintf("%s @@ to_tsquery('spanish', '%s:*')", key, value))
			} else {
				clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
			}
		}
		q += strings.Join(clausulas, " AND ")

	}

	rows, error_find := db.Query(ctx, q+" ORDER BY name ASC LIMIT $1 OFFSET $2", input_limit, input_offset)
	if error_find != nil {
		return oListProvider, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		oProvider := &provider_model.Provider{}
		rows.Scan(&oProvider.Id, &oProvider.IdBusiness, &oProvider.Name, &oProvider.Description, &oProvider.LegalData, &oProvider.AddressData, &oProvider.ContactData, &oProvider.DeletedData, &oProvider.CreatedAt, &oProvider.UpdatedAt)
		oListProvider = append(oListProvider, oProvider)
	}

	//Return the list of provider
	return oListProvider, nil
}
