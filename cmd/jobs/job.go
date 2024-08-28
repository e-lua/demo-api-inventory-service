package job

import (
	automated_service "github.com/e-lua/demo-api-inventory-clean-architecture/internal/services/automated"
)

type JobConfig struct {
	AutomatedService *automated_service.AutomatedService
}

func (job *JobConfig) Start() {
	go job.CleanDataFromTrash()
}
