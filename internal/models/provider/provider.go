package provider

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

type LegalData struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Legal_identity string `json:"legalIdentity"`
}

type AddressData struct {
	Address   string `json:"address"`
	Reference string `json:"reference"`
}

type ContactData struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

type Provider struct {
	Id          string       `json:"id"`
	IdBusiness  string       `json:"idBusiness"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	LegalData   *LegalData   `json:"legalData"`
	AddressData *AddressData `json:"addressData"`
	ContactData *ContactData `json:"contactData"`
	DeletedData *DeletedData `json:"deletedData"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

//Constructor

func NewProvider(id string, id_business string, name string, description string, legal_data *LegalData, address_data *AddressData, contact_data *ContactData, deleted_data *DeletedData) *Provider {
	return &Provider{
		Id:          id,
		IdBusiness:  id_business,
		Name:        name,
		Description: description,
		LegalData:   legal_data,
		AddressData: address_data,
		ContactData: contact_data,
		DeletedData: deleted_data,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

//Method

func (provider *Provider) IsValid() (error, bool) {

	if provider.IdBusiness == "" {
		return fmt.Errorf("idBusiness must be sent"), false
	}

	if len(provider.Name) < 5 || len(provider.Name) > 100 {
		return fmt.Errorf("name can not be less than 5 length or exceed 100 length"), false
	}

	if len(provider.Description) > 250 {
		return fmt.Errorf("description can not exceed 250 length"), false
	}

	if regexp.MustCompile(`[<>]`).MatchString(provider.Description) || regexp.MustCompile(`[<>]`).MatchString(provider.Name) {
		return fmt.Errorf("name and description must be sent without '<' '>'"), false
	}

	return nil, true
}
