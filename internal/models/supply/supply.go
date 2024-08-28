package supply

import (
	"fmt"
	"regexp"
	"time"
)

// Model

type DeletedData struct {
	Is_deleted          bool      `json:"isDeleted"`
	Is_sended_to_delete bool      `json:"isSendedToDelete"`
	Sended_to_delete_at time.Time `json:"sendedToDeleteAt"`
	Deleted_at          time.Time `json:"deletedAt"`
}

type GenericData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Supply struct {
	Id            string       `json:"id"`
	IdBusiness    string       `json:"idBusiness"`
	SKU           string       `json:"sku"`
	Name          string       `json:"name"`
	Stock         float32      `json:"stock"`
	Description   string       `json:"description"`
	MeasureData   *GenericData `json:"measureData"`
	WarehouseData *GenericData `json:"warehouseData"`
	ProviderData  *GenericData `json:"providerData"`
	DeletedData   *DeletedData `json:"deletedData"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
}

type SupplyOutOfStock struct {
	IdBusiness string `json:"idBusiness"`
	Quantity   int    `json:"quantity"`
}

//Constructor

func NewSupply(id string, id_business string, sku string, name string, description string, measure_data *GenericData, warehouse_data *GenericData, provider_data *GenericData, deleted_data *DeletedData) *Supply {
	return &Supply{
		Id:            id,
		IdBusiness:    id_business,
		SKU:           sku,
		Name:          name,
		Description:   description,
		MeasureData:   measure_data,
		WarehouseData: warehouse_data,
		ProviderData:  provider_data,
		DeletedData:   deleted_data,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

//Method

func (supply *Supply) IsValid() (error, bool) {

	if supply.IdBusiness == "" {
		return fmt.Errorf("idBusiness must be sent"), false
	}

	if len(supply.SKU) < 1 || len(supply.SKU) > 50 {
		return fmt.Errorf("sku can not be less than 5 length or exceed 100 length"), false
	}

	if len(supply.Name) < 5 || len(supply.Name) > 100 {
		return fmt.Errorf("name can not be less than 5 length or exceed 100 length"), false
	}

	if len(supply.Description) > 250 {
		return fmt.Errorf("description can not exceed 250 length"), false
	}

	if regexp.MustCompile(`[<>]`).MatchString(supply.Description) || regexp.MustCompile(`[<>]`).MatchString(supply.Name) || regexp.MustCompile(`[<>]`).MatchString(supply.SKU) {
		return fmt.Errorf("sku, name and description must be sent without '<' '>'"), false
	}

	return nil, true
}
