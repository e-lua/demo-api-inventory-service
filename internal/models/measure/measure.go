package measure

import (
	"fmt"
	"regexp"
)

// Model

type Measure struct {
	Id         string `json:"id"`
	IdBusiness string `json:"idBusiness"`
	Name       string `json:"name"`
}

//Constructor

func NewMeasure(id string, id_business string, name string) *Measure {
	return &Measure{
		Id:         id,
		IdBusiness: id_business,
		Name:       name,
	}
}

//Method

func (measure *Measure) IsValid() (error, bool) {

	if measure.IdBusiness == "" {
		return fmt.Errorf("idBusiness must be sent"), false
	}

	if len(measure.Name) < 1 || len(measure.Name) > 50 {
		return fmt.Errorf("name can not be less than 1 length or exceed 50 length"), false
	}

	if regexp.MustCompile(`[<>]`).MatchString(measure.Name) {
		return fmt.Errorf("name must be sent without '<' '>'"), false
	}

	return nil, true
}
