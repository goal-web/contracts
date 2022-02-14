package contracts

import "time"

// CacheStoreProvider 缓存存储提供者
// cache storage provider.
type CacheStoreProvider func(cacheConfig Fields) CacheStore

type CacheFactory interface {

	// Store 按名称获取缓存实例
	// Get a cache store instance by name.
	Store(name ...string) CacheStore

	// Extend 扩展缓存实例
	// Extended cache instance
	Extend(drive string, cacheStoreProvider CacheStoreProvider)
}

type CacheStore interface {

	// Get 按键从缓存中检索项目
	// Retrieve an item from the cache by key.
	Get(key string) interface{}

	// Many 按键从缓存中检索多个项目
	// Retrieve multiple items from the cache by key.
	Many(keys []string) []interface{}

	// Put 在缓存中存储一个项目
	// Store an item in the cache.
	Put(key string, value interface{}, seconds time.Duration) error

	// Add 如果键不存在，则将项目存储在缓存中
	// Store an item in the cache if the key does not exist.
	Add(key string, value interface{}, ttl ...time.Duration) bool

	// Pull 从缓存中检索项目并将其删除
	// Retrieve an item from the cache and delete it.
	Pull(key string, defaultValue ...interface{}) interface{}

	// PutMany 将多个项目存储在缓存中
	// Store multiple items in the cache.
	PutMany(values map[string]interface{}, seconds time.Duration) error

	// Increment 增加缓存中项目的值
	// increment the value of an item in the cache.
	Increment(key string, value ...int64) (int64, error)

	// Decrement 减少缓存中项目的值
	// decrement the value of an item in the cache.
	Decrement(key string, value ...int64) (int64, error)

	// Forever 将项目无限期地存储在缓存中
	// Store an item in the cache indefinitely.
	Forever(key string, value interface{}) error

	// Forget 从缓存中删除项目
	// Remove an item from the cache.
	Forget(key string) error

	// Flush 从缓存中删除所有项目
	// Remove all items from the cache.
	Flush() error

	// GetPrefix 获取缓存键前缀
	// get the cache key prefix.
	GetPrefix() string

	// Remember 从缓存中获取一个项目，或者执行给定的闭包并存储结果。
	// get an item from the cache, or execute the given Closure and store the result.
	Remember(key string, ttl time.Duration, provider InstanceProvider) interface{}

	// RememberForever 从缓存中获取一个项目，或者执行给定的闭包并永久存储结果
	// get an item from the cache, or execute the given Closure and store the result forever.
	RememberForever(key string, provider InstanceProvider) interface{}
}
