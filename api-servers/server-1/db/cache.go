package db

import (
	"time"

	"github.com/api-server/lcs42/models"
	"github.com/api-server/lcs42/utils"
	"github.com/go-redis/redis"
)

// Cache interface to execute all methods
type Cache interface {
	Get(userId string) (models.User, error)
	Create(user models.User) error
	Update(newData models.User) (models.User, error)
	Delete() error
	CloseRds()
}

// CacheDB store a Db for redisClient
type CacheDb struct {
	Db *redis.Client
}

// NewCacheDB constructor to CacheDB
func NewCacheDb(redisClient *redis.Client) *CacheDb {
	return &CacheDb{Db: redisClient}
}

// Create method to create a user in rds server
func (c *CacheDb) Create(user models.User) error {
	userBuffer, err := utils.ParseToBytes(
		map[string]string{
			"id":    user.Id,
			"name":  user.Name,
			"email": user.Email,
		},
	)
	if err != nil {
		return err
	}

	rdsResponse := c.Db.Set(user.Id, userBuffer, time.Duration(30*time.Minute))

	if rdsResponse.Err() != nil {
		return rdsResponse.Err()
	}
	return nil
}

// Get method to geta an user in rds server
func (c *CacheDb) Get(userId string) (models.User, error) {
	var user models.User

	result := c.Db.Get(userId)
	if result.Err() != nil {
		return user, result.Err()
	}

	userBytes, err := result.Bytes()
	if err != nil {
		return user, err
	}

	user, err = utils.ParseToModels(userBytes)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Delete method to delete an user in rds server
func (c *CacheDb) Delete() error {
	return nil
}

// Update method to update user data in rds server
func (c *CacheDb) Update(newData models.User) (models.User, error) {
	if err := c.Create(newData); err != nil {
		return models.User{}, err
	}

	return newData, nil
}

// CloseRds ensure that the connection with rds ends
func (c *CacheDb) CloseRds() {
	c.Db.Close()
}
