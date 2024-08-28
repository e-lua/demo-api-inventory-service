package job

import (
	"log"
	"time"
)

func (job *JobConfig) CleanDataFromTrash() {

	var task_frequency = 1 * time.Hour

	for {

		time.Sleep(task_frequency)
		err := job.AutomatedService.CleanTrash()
		if err != nil {
			log.Println(err.Error())
		}

	}

}
