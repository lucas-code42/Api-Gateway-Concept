package db

import (
	"time"

	"github.com/api-server/lcs42/models"
	"github.com/api-server/lcs42/utils"
	"github.com/go-redis/redis"
)

type Cache interface {
	Get() error
	GetAll() error
	Create(user models.User) error
	Update() error
	Delete() error
	CloseRds()
}

type CacheDb struct {
	Db *redis.Client
}

func NewCacheDb(redisClient *redis.Client) *CacheDb {
	return &CacheDb{Db: redisClient}
}

func (c *CacheDb) Create(user models.User) error {
	userBuffer, err := utils.ParseToBytes(
		map[string]string{
			"userId":   user.Id,
			"userName": user.Name,
			"userEmal": user.Email,
		},
	)
	if err != nil {
		return err
	}

	rdsResponse := c.Db.Set("user", userBuffer, time.Duration(10*time.Minute))

	if rdsResponse.Err() != nil {
		return rdsResponse.Err()
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

func (c *CacheDb) CloseRds() {
	c.Db.Close()
}
