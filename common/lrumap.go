package common

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	lruList  *list.List
}

type Entry struct {
	key   string
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		lruList:  list.New(),
	}
}

func (lru *LRUCache) Get(key string) (int, bool) {
	if elem, found := lru.cache[key]; found {
		lru.lruList.MoveToFront(elem)
		return elem.Value.(*Entry).value, true
	}
	return 0, false
}

func (lru *LRUCache) Put(key string, value int) {
	if elem, found := lru.cache[key]; found {
		// Update existing entry
		elem.Value.(*Entry).value = value
		lru.lruList.MoveToFront(elem)
	} else {
		// Add new entry
		if lru.lruList.Len() >= lru.capacity {
			// Remove the least recently used element
			oldest := lru.lruList.Back()
			if oldest != nil {
				delete(lru.cache, oldest.Value.(*Entry).key)
				lru.lruList.Remove(oldest)
			}
		}
		entry := &Entry{key: key, value: value}
		newElem := lru.lruList.PushFront(entry)
		lru.cache[key] = newElem
	}
}
