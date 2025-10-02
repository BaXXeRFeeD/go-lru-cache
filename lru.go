package lrucache

import "container/list"

func New(cap int) Cache {
	return &LruCache{
		cache: make(map[int]*list.Element, cap),
		cap:   cap,
		order: list.New(),
	}
}

type LruCache struct {
	cap   int
	cache map[int]*list.Element
	order *list.List
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
