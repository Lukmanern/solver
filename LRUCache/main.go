package main

import "container/list"

type LRUCache struct {
	cache    map[int]*list.Element
	linkList *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*list.Element, capacity),
		linkList: list.New(),
		capacity: capacity,
	}
}

// 0 is key and 1 is value
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	elem := this.cache[key]
	this.linkList.MoveToFront(elem)
	return elem.Value.([]int)[1]
}

func (this *LRUCache) Put(key int, value int) {
	// if capacity reached
	if elem, ok := this.cache[key]; ok {
		this.linkList.Remove(elem)
		newelem := this.linkList.PushFront([]int{key, value})
		this.cache[key] = newelem
		return
	}
	// max
	if len(this.cache) == this.capacity {
		elem := this.linkList.Back()
		v := this.linkList.Remove(elem)
		delete(this.cache, v.([]int)[0])
	}
	newelem := this.linkList.PushFront([]int{key, value})
	this.cache[key] = newelem
}