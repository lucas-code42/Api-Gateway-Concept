package db

import (
	"fmt"

	"github.com/api-server/lcs42/config"
	"github.com/go-redis/redis"
)

// MountRds return's an instance of interface
func MountRds() Cache {
	// consult docker-compose
	var rds Cache = NewCacheDb(redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.REDIS_HOST, config.REDIS_PORT),
		Password: config.REDIS_PASSWORD,
		DB:       0, // default db
	}))
	return rds
}
