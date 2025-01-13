package cache

import (
	"fmt"
	"os"
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

		<-ticker.C

		var timesToPrint []float64

		c.mx.Lock()
		if len(c.Data) == 0 {
			c.mx.Unlock()
			fmt.Fprintf(os.Stderr, "\nNothing to clean, cache is empty!\n")
			continue
		}

		for k, v := range c.Data {
			timePassed := time.Since(v.createdAt)
			if timePassed > c.interval {

				delete(c.Data, k)
				timesToPrint = append(timesToPrint, timePassed.Seconds())

			}

		}
		c.mx.Unlock()
		for _, sec := range timesToPrint {
			fmt.Fprintf(os.Stderr, "\nRemoved from cache!\nit was %2.f seconds old!\n", sec)
		}
		if len(timesToPrint) > 0 {
			fmt.Fprintf(os.Stderr, "\nCache cleaned! Total: %v\n", len(timesToPrint))
		}
		//fmt.Printf("\nCache cleaned!\ntotal:%v\n", delCounter)

	}
}
