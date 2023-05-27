package db

import (
	"context"
	"time"

	"github.com/api-server/lcs42/models"
	"github.com/go-redis/redis"
)

type CacheDb struct {
	Db  *redis.Client
	Ctx *context.Context
}

type Cache interface {
	Get() error
	GetAll() error
	Create() error
	Update() error
	Delete() error
}

func NewCacheDb(redisClient *redis.Client, context *context.Context) *CacheDb {
	return &CacheDb{Db: redisClient, Ctx: context}
}

func (c *CacheDb) Create(user models.User) error {
	err := c.Db.Set(
		"user",
		map[string]string{
			"userId":   user.Id,
			"userName": user.Name,
			"userEmal": user.Email,
		},
		time.Duration(time.Minute*10),
	)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (c *CacheDb) Get() error {
	return nil
}

func (c *CacheDb) GetAll() error {
	return nil
}

func (c *CacheDb) Delete() error {
	return nil
}

func (c *CacheDb) Update() error {
	return nil
}
