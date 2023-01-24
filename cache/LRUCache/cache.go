package lrucache

import (
	"container/list"
	"log"
)

type Cache interface {
	Get(key string) (string, bool)
	Set(key string, value string)
}

type cachedElement struct {
	el    *list.Element
	value string
}

type LRUCache struct {
	m   map[string]*cachedElement
	cap int
	l   list.List
}

func NewLRU(cap int) LRUCache {
	return LRUCache{
		m:   make(map[string]*cachedElement),
		cap: cap,
		l:   list.List{},
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	cacheStored, ok := c.m[key]
	if !ok {
		return "", false
	}

	c.l.MoveToFront(cacheStored.el)
	return cacheStored.value, true
}

func (c *LRUCache) Set(key string, value string) {
	if cacheStored, ok := c.m[key]; !ok {
		c.m[key] = &cachedElement{
			el:    c.l.PushFront(key),
			value: value,
		}

		if c.l.Len() > c.cap {
			lastEl := c.l.Back()
			key, ok := lastEl.Value.(string)
			if !ok {
				log.Fatal("Error casting to string")
			}
			c.l.Remove(lastEl)
			delete(c.m, key)
		}
	} else {
		cacheStored.value = value
		c.l.MoveToFront(cacheStored.el)
	}
}

// CREDITS
// I use this [guide](https://www.youtube.com/watch?v=DvmMYD8oaZw) as a reference
