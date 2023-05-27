package db

import "github.com/go-redis/redis"

type CacheDb struct {
	db *redis.Client
}

func InitCacheDb(db *redis.Client) *CacheDb {
	return &CacheDb{db}
}

func (c *CacheDb) Get() {

}

func (c *CacheDb) GetAll() {

}

func (c *CacheDb) Delete() {

}

func (c *CacheDb) Update() {

}