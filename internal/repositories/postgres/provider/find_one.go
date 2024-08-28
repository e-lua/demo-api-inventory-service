package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	provider_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/provider"
)

func (pr *ProviderRepository) FindOne(input_id string) (*provider_model.Provider, error) {

	//Initialization
	oProvider := &provider_model.Provider{}

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_id != "" {
		filters["id"] = input_id
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
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}

	rows, error_find := db.Query(ctx, q)
	if error_find != nil {
		return oProvider, error_find
	}
	defer rows.Close()

	//Scan the row
	for rows.Next() {
		rows.Scan(&oProvider.Id, &oProvider.IdBusiness, &oProvider.Name, &oProvider.Description, &oProvider.LegalData, &oProvider.AddressData, &oProvider.ContactData, &oProvider.DeletedData, &oProvider.CreatedAt, &oProvider.UpdatedAt)
	}

	//Return the provider
	return oProvider, nil
}
