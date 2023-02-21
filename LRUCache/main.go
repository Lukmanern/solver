package main

import (
	"container/list"
	"fmt"
)

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // should return 1

	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // should return -1

	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // should return -1
	fmt.Println(cache.Get(3)) // should return 3
	fmt.Println(cache.Get(4)) // should return 4
}

// LRUCache is a struct that represents a LRU cache.
type LRUCache struct {
	cache    map[int]*list.Element
	linkList *list.List
	capacity int
	// cache 	 : a map to store the cache values and their corresponding list element
	// linkList  : a doubly linked list to store the order of access for cache values
	// capacity  : the maximum capacity of the cache
}

// Constructor is a function that creates 
// a new LRUCache instance with a given capacity.
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*list.Element, capacity),
		linkList: list.New(),
		capacity: capacity,
	}
}

// Get is a function that retrieves a value 
// from the cache given a key. It returns 
// the value or -1 if the key is not found.
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	elem := this.cache[key]
	this.linkList.MoveToFront(elem)
	return elem.Value.([]int)[1]
}

// Put is a function that adds a new key-value pair 
// to the cache. If the key already exists in the cache, 
// the value is updated and the key is moved to the front 
// of the list. If the cache is at capacity, the least 
// recently used key-value pair is evicted from the cache.
func (this *LRUCache) Put(key int, value int) {
	// if key already exists, update 
	// value and move key to front of list
	if elem, ok := this.cache[key]; ok {
		this.linkList.Remove(elem)
		newelem := this.linkList.PushFront([]int{key, value})
		this.cache[key] = newelem
		return
	}
	// if cache is at capacity, evict least 
	// recently used key-value pair
	if len(this.cache) == this.capacity {
		elem := this.linkList.Back()
		v := this.linkList.Remove(elem)
		delete(this.cache, v.([]int)[0])
	}
	// add new key-value pair to front of list and cache
	newelem := this.linkList.PushFront([]int{key, value})
	this.cache[key] = newelem
}
