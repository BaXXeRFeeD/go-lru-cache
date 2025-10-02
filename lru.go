//go:build !solution

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
