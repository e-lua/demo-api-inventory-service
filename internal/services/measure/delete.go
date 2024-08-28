package measure

import "errors"

func (ms *MeasureService) Delete(input_idmeasure string) (int, error) {

	//Delete the Measure
	error_delete_measure := ms.MeasurePostgresRepository.DeleteOne(input_idmeasure)
	if error_delete_measure != nil {
		return 5057, errors.New("error delete measure, details: " + error_delete_measure.Error())
	}

	//OK
	return 0, nil
}
