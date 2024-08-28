package provider_test

import (
	"testing"

	provider_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/provider"
)

func TestProvider_IsValid_IdBusiness(t *testing.T) {

	provider := provider_model.Provider{
		Id:          "ac28ec48-bd4a-4fb8-b12a-2e7ddba5b116",
		IdBusiness:  "",
		Name:        "Milkyway South",
		Description: "Provider of apples",
	}

	new_measure := provider_model.NewProvider(provider.Id, provider.IdBusiness, provider.Name, provider.Description, provider.LegalData, provider.AddressData, provider.ContactData, provider.DeletedData)

	err, is_valid := new_measure.IsValid()

	if is_valid != false {
		t.Error("Expected isValid to be false")
	}

	if err == nil {
		t.Error("Expected err to be not nil")
	}

	if err.Error() != "idBusiness must be sent" {
		t.Errorf("Expected err.Error() to be 'idBusiness must be sent', got '%s'", err.Error())
	}
}

func TestProvider_IsValid_Name(t *testing.T) {

	provider := provider_model.Provider{
		Id:          "ac28ec48-bd4a-4fb8-b12a-2e7ddba5b116",
		IdBusiness:  "c85d5db9-a676-4c83-bf04-767da6a5074b",
		Name:        "",
		Description: "Provider of apples",
	}

	new_measure := provider_model.NewProvider(provider.Id, provider.IdBusiness, provider.Name, provider.Description, provider.LegalData, provider.AddressData, provider.ContactData, provider.DeletedData)

	err, is_valid := new_measure.IsValid()

	if is_valid != false {
		t.Error("Expected isValid to be false")
	}

	if err == nil {
		t.Error("Expected err to be not nil")
	}

	if err.Error() != "name can not be less than 5 length or exceed 100 length" {
		t.Errorf("Expected err.Error() to be 'name can not be less than 5 length or exceed 100 length', got '%s'", err.Error())
	}
}

func TestProvider_IsValid_Description(t *testing.T) {

	provider := provider_model.Provider{
		Id:          "ac28ec48-bd4a-4fb8-b12a-2e7ddba5b116",
		IdBusiness:  "c85d5db9-a676-4c83-bf04-767da6a5074b",
		Name:        "Milkyway South",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam accumsan eros non mi molestie, eget bibendum augue iaculis. Sed luctus nisl eget nibh faucibus, eget facilisis nunc pulvinar. Mauris et eros eget diam convallis maximus. Donec tristique, odio vel varius posuere, nulla magna rhoncus erat, sit amet fermentum dui mauris sed tellus. Suspendisse id accumsan risus, eget mollis magna. Suspendisse fringilla nisi a velit suscipit, at fringilla quam maximus. Ut vitae fermentum mauris. Duis semper ut nulla sed semper. Nullam id pretium magna, vitae fringilla nisi. Sed et nisl maximus, auctor libero eu, tempor lectus.",
	}

	new_measure := provider_model.NewProvider(provider.Id, provider.IdBusiness, provider.Name, provider.Description, provider.LegalData, provider.AddressData, provider.ContactData, provider.DeletedData)

	err, is_valid := new_measure.IsValid()

	if is_valid != false {
		t.Error("Expected isValid to be false")
	}

	if err == nil {
		t.Error("Expected err to be not nil")
	}

	if err.Error() != "description can not exceed 250 length" {
		t.Errorf("Expected err.Error() to be 'description can not exceed 250 length', got '%s'", err.Error())
	}
}

func TestProvider_IsValid(t *testing.T) {

	provider := provider_model.Provider{
		Id:          "ac28ec48-bd4a-4fb8-b12a-2e7ddba5b116",
		IdBusiness:  "c85d5db9-a676-4c83-bf04-767da6a5074b",
		Name:        "Milkyway South",
		Description: "Provider of apples",
	}

	new_measure := provider_model.NewProvider(provider.Id, provider.IdBusiness, provider.Name, provider.Description, provider.LegalData, provider.AddressData, provider.ContactData, provider.DeletedData)

	err, is_valid := new_measure.IsValid()

	if is_valid != true {
		t.Error("Expected isValid to be true")
	}

	if err != nil {
		t.Error("Expected err to be nil")
	}
}
