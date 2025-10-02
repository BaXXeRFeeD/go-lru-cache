//go:build !change

package lrucache

import "container/list"

type Cache interface {
	// Get returns value associated with the key.
	//
	// The second value is a bool that is true if the key exists in the cache,
	// and false if not.
	Get(key int) (int, bool)
	// Set updates value associated with the key.
	//
	// If there is no key in the cache new (key, value) pair is created.
	Set(key, value int)
	// Range calls function f on all elements of the cache
	// in increasing access time order.
	//
	// Stops earlier if f returns false.
	Range(f func(key, value int) bool)
	// Clear removes all keys and values from the cache.
	Clear()
}

type entry struct {
	key   int
	value int
}

func (c *LruCache) Get(key int) (int, bool) {
	if e, ok := c.cache[key]; ok {
		c.order.MoveToFront(e)
		return e.Value.(*entry).value, true
	}
	return 0, false
}

func (c *LruCache) Set(key, value int) {
	if e, ok := c.cache[key]; ok {
		e.Value.(*entry).value = value
		c.order.MoveToFront(e)
		return
	}
	if c.cap > 0 && c.order.Len() >= c.cap {
		e := c.order.Back()
		if e != nil {
			c.order.Remove(e)
			delete(c.cache, e.Value.(*entry).key)
		}
	}
	if c.cap > 0 {
		e := c.order.PushFront(&entry{key, value})
		c.cache[key] = e
	}
}

func (c *LruCache) Range(f func(key, value int) bool) {
	for e := c.order.Back(); e != nil; e = e.Prev() {
		ent := e.Value.(*entry)
		if !f(ent.key, ent.value) {
			break
		}
	}
}

func (c *LruCache) Clear() {
	cache := make(map[int]*list.Element, c.cap)
	c.cache = cache
	c.order = list.New()
}
