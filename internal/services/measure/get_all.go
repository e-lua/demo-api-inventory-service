package measure

import (
	"errors"

	measure_model "github.com/e-lua/demo-api-inventory-clean-architecture/internal/models/measure"
)

func (ms *MeasureService) GetAll(input_idbusiness string, input_name string, input_limit int, input_offset int) (int, []*measure_model.Measure, error) {

	//Get the all measures
	list_measures, error_find_measure := ms.MeasurePostgresRepository.FindMany(input_idbusiness, input_name, input_limit, input_offset)
	if error_find_measure != nil {
		return 5057, []*measure_model.Measure{}, errors.New("error find measure, details: " + error_find_measure.Error())
	}

	//OK
	return 0, list_measures, nil
}
