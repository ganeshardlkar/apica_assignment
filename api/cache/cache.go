package cache

import (
	"container/list"
	"database/sql"
	"log"
	"sync"
	"time"
)

type Cache struct {
	capacity   int
	expiration time.Duration
	items      map[string]*list.Element
	eviction   *list.List
	db         *sql.DB
	mu         sync.Mutex
}

type CacheItem struct {
	key        string
	value      string
	expiration time.Time
}

func NewCache(capacity int, expiration time.Duration, db *sql.DB) *Cache {
	return &Cache{
		capacity:   capacity,
		expiration: expiration,
		items:      make(map[string]*list.Element),
		eviction:   list.New(),
		db:         db,
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		if time.Now().After(item.Value.(*CacheItem).expiration) {
			c.eviction.Remove(item)
			delete(c.items, key)
			return "", false
		}
		c.eviction.MoveToFront(item)
		return item.Value.(*CacheItem).value, true
	}

	return "", false
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(c.expiration)
	item := &CacheItem{
		key:        key,
		value:      value,
		expiration: expiration,
	}

	if existing, ok := c.items[key]; ok {
		c.eviction.MoveToFront(existing)
		existing.Value = item
	} else {
		if len(c.items) >= c.capacity {
			toEvict := c.eviction.Back()
			if toEvict != nil {
				evicted := c.eviction.Remove(toEvict)
				delete(c.items, evicted.(*CacheItem).key)
			}
		}
		c.items[key] = c.eviction.PushFront(item)
	}

	_, err := c.db.Exec("INSERT INTO lru_cache (key, value, expiration) VALUES ($1, $2, $3)", key, value, expiration)
	if err != nil {
		log.Printf("Error storing data in the database: %v", err)
	}
}

func (c *Cache) StartCleanup() {
	ticker := time.NewTicker(c.expiration / 2)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()
		for key, item := range c.items {
			if time.Now().After(item.Value.(*CacheItem).expiration) {
				c.eviction.Remove(item)
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
