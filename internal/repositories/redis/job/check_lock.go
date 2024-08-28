package job

import (
	"context"
	"time"
)

func (jr *JobRedisRepository) CheckLock(input_key string, input_value string) (bool, error) {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	rdb := jr.ConnMasterRedis

	reply, err := rdb.SetNX(ctx, input_value, input_value, 60).Result()
	if err != nil {
		return false, err
	}

	return reply, nil
}
