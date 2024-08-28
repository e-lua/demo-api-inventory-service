package automated

import (
	"errors"
)

func (as *AutomatedService) CleanTrash() error {

	//Check Job
	locked, err_check_lock := as.JobRedisRepository.CheckLock("clean_data_from_trash", "locked")
	if err_check_lock != nil {
		return errors.New("error get last_task_time from Redis:" + err_check_lock.Error())
	}
	if locked {
		return errors.New("job taked before")
	}

	//Clean Provider
	error_update_provider := as.ProviderPostgresRepository.UpdateManyDelete()
	if error_update_provider != nil {
		return errors.New("error clean Provider, details: " + error_update_provider.Error())
	}

	//Clean Supply
	error_update_supply := as.SupplyPostgresRepository.UpdateManyDelete()
	if error_update_supply != nil {
		return errors.New("error clean Supply, details: " + error_update_supply.Error())
	}

	//Clean Warehouse
	error_update_warehouse := as.WarehousePostgresRepository.UpdateManyDelete()
	if error_update_warehouse != nil {
		return errors.New("error clean Warehouse, details: " + error_update_warehouse.Error())
	}

	return nil
}
