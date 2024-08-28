package kardex_supply

import "errors"

func (kss *KardexSupplyService) Delete(input_idkardex string) (int, error) {

	//Delete the KardexSupply
	error_delete_kardex := kss.KSPostgresRepository.DeleteOne(input_idkardex)
	if error_delete_kardex != nil {
		return 5057, errors.New("error delete kardex, details: " + error_delete_kardex.Error())
	}

	//OK
	return 0, nil
}
