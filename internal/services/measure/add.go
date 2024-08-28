package measure

import (
	"errors"

	"github.com/google/uuid"

	measure_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/measure"
)

func (ms *MeasureService) Add(input_measure *measure_model.Measure) (int, error) {

	//Validation of the Business Rules
	error_valid, is_valid := input_measure.IsValid()
	if !is_valid {
		return 4052, error_valid
	}

	//Storage the New Measure
	new_measure := measure_model.NewMeasure(uuid.New().String(), input_measure.IdBusiness, input_measure.Name)
	error_create_measure := ms.MeasurePostgresRepository.InsertOne(new_measure)
	if error_create_measure != nil {
		return 5057, errors.New("error create measure, details: " + error_create_measure.Error())
	}

	//OK
	return 0, nil
}
