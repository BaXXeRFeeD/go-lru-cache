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
