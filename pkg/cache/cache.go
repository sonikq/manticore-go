package cache

import (
	"errors"
	"gitlab.geogracom.com/skdf/skdf-manticore-go/app/consts"
	"sync"
	"time"
)

type Item struct {
	Object     interface{}
	Expiration int64
}

func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

type Cache struct {
	*cache
}

type cache struct {
	items map[interface{}]Item
	mu    sync.RWMutex
}

func New() *Cache {
	return &Cache{
		newCache(make(map[interface{}]Item)),
	}
}

func newCache(m map[interface{}]Item) *cache {
	c := &cache{
		items: m,
	}
	return c
}

func (c *Cache) Set(key, value interface{}, ttlSeconds time.Duration) {
	exp := time.Now().Add(ttlSeconds).UnixNano()

	c.mu.Lock()

	c.items[key] = Item{
		Object:     value,
		Expiration: exp,
	}

	c.mu.Unlock()
}

func (c *Cache) Get(key interface{}) (interface{}, error) {
	c.mu.RLock()

	item, found := c.items[key]
	if !found {
		c.mu.RUnlock()
		return nil, errors.New(consts.AccessDenied)
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return nil, errors.New(consts.CacheTimeExpired)
		}
	}
	c.mu.RUnlock()
	return item.Object, nil
}

func (c *Cache) FlushCache() {
	c.mu.Lock()
	c.items = map[interface{}]Item{}
	c.mu.Unlock()
}
