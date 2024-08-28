package kardex_supply_test

import (
	"testing"
	"time"

	kardex_supply_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/kardex_supply"
)

func TestKardexSupply_IsValid(t *testing.T) {

	// Case 1: Valid
	kardex_supply := kardex_supply_model.KardexSupply{
		Id:         "ac28ec48-bd4a-4fb8-b12a-2e7ddba5b116",
		IdBusiness: "c85d5db9-a676-4c83-bf04-767da6a5074b",
		IdSupply:   "dc647735-5870-4ca8-95d0-e142201852f3",
		IdType:     4,
		IdCategory: 9,
		Date:       time.Now(),
		Quantity:   11.50,
		TotalCost:  15.5,
		UnitCost:   35.5,
		UpdatedBy:  "Admin",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	new_kardex := kardex_supply_model.NewKardexSupply(kardex_supply.Id, kardex_supply.IdBusiness, kardex_supply.IdSupply, kardex_supply.Date, kardex_supply.IdType, kardex_supply.IdCategory, kardex_supply.Quantity, kardex_supply.TotalCost, kardex_supply.UpdatedBy)

	err, valid := new_kardex.IsValid()
	if err != nil || !valid {
		t.Errorf("IsValid failed with valid data: %v, %v", err, valid)
	}

	// Case 2: Empty IdSupply
	kardex_supply.IdSupply = ""
	err, valid = kardex_supply.IsValid()
	if err == nil || valid {
		t.Errorf("IsValid failed with empty IdSupply: %v, %v", err, valid)
	}

	// Case 3: Empty IdBusiness
	kardex_supply.IdSupply = "dc647735-5870-4ca8-95d0-e142201852f3"
	kardex_supply.IdBusiness = ""
	err, valid = kardex_supply.IsValid()
	if err == nil || valid {
		t.Errorf("IsValid failed with empty IdBusiness: %v, %v", err, valid)
	}

	// Case 4: Empty Date
	kardex_supply.IdBusiness = "c85d5db9-a676-4c83-bf04-767da6a5074b"
	kardex_supply.Date = time.Time{}
	err, valid = kardex_supply.IsValid()
	if err == nil || valid {
		t.Errorf("IsValid failed with empty Date: %v, %v", err, valid)
	}

	// Case 5: Empty IdType
	kardex_supply.Date = time.Now()
	kardex_supply.IdType = 0
	err, valid = kardex_supply.IsValid()
	if err == nil || valid {
		t.Errorf("IsValid failed with empty IdType: %v, %v", err, valid)
	}

	// Case 6: Empty IdCategory
	kardex_supply.IdType = 4
	kardex_supply.IdCategory = 0
	err, valid = kardex_supply.IsValid()
	if err == nil || valid {
		t.Errorf("IsValid failed with empty IdCategory: %v, %v", err, valid)
	}

	// Case 7: Empty Quantity
	kardex_supply.IdCategory = 9
	kardex_supply.Quantity = 0
	err, valid = kardex_supply.IsValid()
	if err == nil || valid {
		t.Errorf("IsValid failed with empty Quantity: %v, %v", err, valid)
	}

}
