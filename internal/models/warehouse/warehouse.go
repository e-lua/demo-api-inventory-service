package warehouse

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

type Warehouse struct {
	Id          string       `json:"id"`
	IdBusiness  string       `json:"idBusiness"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	DeletedData *DeletedData `json:"deletedData"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

//Constructor

func NewWarehouse(id string, id_business string, name string, description string, deleted_data *DeletedData) *Warehouse {
	return &Warehouse{
		Id:          id,
		IdBusiness:  id_business,
		Name:        name,
		Description: description,
		DeletedData: deleted_data,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

//Method

func (warehouse *Warehouse) IsValid() (error, bool) {

	if warehouse.IdBusiness == "" {
		return fmt.Errorf("idBusiness must be sent"), false
	}

	if len(warehouse.Name) < 5 || len(warehouse.Name) > 50 {
		return fmt.Errorf("name can not be less than 5 length or exceed 50 length"), false
	}

	if len(warehouse.Description) > 250 {
		return fmt.Errorf("description can not exceed 250 length"), false
	}

	if regexp.MustCompile(`[<>]`).MatchString(warehouse.Description) || regexp.MustCompile(`[<>]`).MatchString(warehouse.Name) {
		return fmt.Errorf("name and description must be sent without '<' '>'"), false
	}

	return nil, true
}
