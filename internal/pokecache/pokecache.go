package cache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte // raw caching data
}

type Cache struct {
	mx       sync.Mutex
	Data     map[string]cacheEntry
	interval time.Duration
}

func NewCache(lifetime time.Duration) *Cache {
	c := &Cache{
		Data:     make(map[string]cacheEntry),
		interval: lifetime,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {

	c.mx.Lock()

	c.Data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mx.Unlock()
	fmt.Print("\nCache updated!\n\n")
}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mx.Lock()
	defer c.mx.Unlock()
	if val, ok := c.Data[key]; !ok {

		fmt.Print("\nNo data in cache!\n")

		return nil, ok

	} else {
		fmt.Print("\nFound in cache!\n\n")
		return val.val, ok
	}
}

func (c *Cache) reapLoop() {

	ticker := time.NewTicker(c.interval)

	for {
		delCounter := 0

		<-ticker.C
		c.mx.Lock()
		entries := c.Data
		for k, v := range entries {
			timePassed := time.Since(v.createdAt)
			if timePassed > c.interval {

				delete(c.Data, k)
				delCounter++
				fmt.Printf("Removed from cache!\nit was %2.f seconds old!\n", timePassed.Seconds())

			}

		}
		c.mx.Unlock()

		//fmt.Printf("\nCache cleaned!\ntotal:%v\n", delCounter)

	}
}
