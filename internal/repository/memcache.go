package repository

import (
	"complaint_service/internal/config"
	"fmt"
	"log"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type SessionCache struct {
	cache *memcache.Client
}

// NewSessionCache является конструктором структуры SessionCache. Принимает на вход переменную типа sqlx.DB и возвращает SessionCache
func NewSessionCache() *SessionCache {
	configs, err := config.LoadEnv()
	if err != nil {
		log.Println(err)
	}

	connStr := fmt.Sprintf("%v:%v", configs.CacheHost, configs.CachePort)
	log.Println("Строка подключения к мемкешу:", connStr)

	cacheServer := memcache.New(connStr)
	cacheServer.Timeout = 5 * time.Second
	return &SessionCache{
		cache: cacheServer,
	}
}

// Set сохраняет данные в кеш. На вход принимает ключ типа string, данные типа []byte и время жизни expiration типа int32. Возвращает ошибку.
func (c *SessionCache) Set(key string, value []byte, expiration int32) error {
	item := &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}

	return c.cache.Set(item)
}

// Get возвращает данные из кеша по ключу. На вход принимает ключ типа string и возвращает данные типа []byte из кеша
func (c *SessionCache) Get(key string) ([]byte, error) {
	item, err := c.cache.Get(key)

	if err != nil {
		return nil, err
	}

	return item.Value, err

}
