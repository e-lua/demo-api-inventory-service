package kardex_supply

import (
	"fmt"
	"time"
)

// Model

type KardexSupply struct {
	Id         string    `json:"id"`
	IdBusiness string    `json:"idBusiness"`
	IdSupply   string    `json:"idSupply"`
	IdType     int       `json:"idType"`
	IdCategory int       `json:"idCategory"`
	Date       time.Time `json:"date"`
	Quantity   float32   `json:"quantity"`
	TotalCost  float32   `json:"totalCost"`
	UnitCost   float32   `json:"unitCost"`
	UpdatedBy  string    `json:"updatedBy"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

//Constructor

func NewKardexSupply(id string, id_business string, id_supply string, date time.Time, id_type int, id_category int, quantity float32, total_cost float32, updated_by string) *KardexSupply {
	return &KardexSupply{
		Id:         id,
		IdBusiness: id_business,
		IdSupply:   id_supply,
		Date:       date,
		IdType:     id_type,
		IdCategory: id_category,
		Quantity:   quantity,
		TotalCost:  total_cost,
		UpdatedBy:  updated_by,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

//Method

func (kardex_supply *KardexSupply) IsValid() (error, bool) {

	if kardex_supply.IdSupply == "" {
		return fmt.Errorf("idSupply must be sent"), false
	}

	if kardex_supply.IdBusiness == "" {
		return fmt.Errorf("idBusiness must be sent"), false
	}

	if kardex_supply.Date.IsZero() {
		return fmt.Errorf("date must be sent"), false
	}

	if kardex_supply.IdType == 0 {
		return fmt.Errorf("idType can not be 0"), false
	}

	if kardex_supply.IdCategory == 0 {
		return fmt.Errorf("idCategory can not be 0"), false
	}

	if kardex_supply.Quantity == 0 {
		return fmt.Errorf("quantity can not be 0"), false
	}

	if kardex_supply.Quantity > 999999 {
		return fmt.Errorf("quantity can not exceed 6 digits"), false
	}

	if kardex_supply.TotalCost > 999999 {
		return fmt.Errorf("totalCost can not exceed 6 digits"), false
	}

	return nil, true
}
