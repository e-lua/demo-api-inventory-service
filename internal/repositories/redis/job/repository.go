package job

import (
	"github.com/go-redis/redis/v8"
)

type JobRedisRepository struct {
	ConnMasterRedis *redis.Client
}

// NewJobRedisRepository will create an object that represent the job.Repository interface
func NewJobRedisRepository(conn_master_redis *redis.Client) *JobRedisRepository {
	return &JobRedisRepository{
		ConnMasterRedis: conn_master_redis,
	}
}
