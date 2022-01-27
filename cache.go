package main

import (
	"fmt"
	"math"
	"time"
	"container/heap"
	"sync"
)

// Item Structure to store the cache values
type Item struct {
	data           int64 // data to be stored in cache
	expirationTime int64 // Expiration time for the item
	TTL int // Time To Live for the Item
	index int // index of the item in the Priority Queue
}

type Cache struct {
	Map map[string]*Item // Key Value pair for the cache
	timer *time.Timer //Timer to check expired item
	m sync.Mutex
}



// Input argument for set function
type Parameters struct {
	data int64 // data to be stored in cache
	key string // key assoicated with the value
	TTL int64 // Time To Live for the item in cache
}

// Set Function for Cache
func (c *Cache) set(parameter Parameters, pq *PriorityQueue) (item *Item, exists bool){
	if parameter.TTL == 0{
		parameter.TTL = 60 // if no TTL is provided assign default value as 60s
	} else if parameter.TTL < 0 {
		panic("ERR: TTL cannot be negative.") // raise error if TTL provided is negative
	}
	old, exists := c.Map[parameter.key] // check whether key already exists in cache 
	expirationTime := time.Now().Unix() + parameter.TTL // update expiration Time
	c.Map[parameter.key] = &Item{data: parameter.data, expirationTime: expirationTime, TTL: int(parameter.TTL), index: pq.Len() + 1} //asign value to cache

	if exists {
		c.Map[parameter.key].index = old.index
		item = c.Map[parameter.key]
		(*pq)[old.index].item = item
		(*pq).update(item)
	} else {
		item = c.Map[parameter.key]
		queueItem := createQueueItem(parameter.key,item)
		heap.Push(pq, queueItem)
	}
	timeLeft := (*pq)[0].item.expirationTime - time.Now().Unix() 
	c.timer.Reset(time.Duration(timeLeft * int64(math.Pow(10,9))))
	return
}

// Get function for cache
func (c *Cache) get(key string) (value int64, item *Item) {
	var state bool
	item, state = c.Map[key] //get value from the Map
	if !state {
		panic("ERR: Key not found.") // raise error if key does not exist in cache
	}
	value = item.data
	return
}

func (c *Cache) delete(key string, pq *PriorityQueue) (item *Item) {
	var state bool
	item, state = c.Map[key]
	if !state {
		fmt.Println(c.Map)
		panic("ERR: Key not found.")
	}
	if pq.Len() == 0 {
		c.timer.Stop()
	} else {
		timeLeft := (*pq)[0].item.expirationTime - time.Now().Unix()
		c.timer.Reset(time.Duration(timeLeft * int64(math.Pow(10,9))))
	}
	delete(c.Map, key)
	return
}

func (c *Cache) printMap() {
	for key, value := range c.Map {
		fmt.Printf("%s: %d %d %d\n",key, value.data, value.expirationTime, value.TTL)
	}
}

func newCache() (cache Cache) {
	cache = Cache{Map: make(map[string]*Item)}
	cache.timer = time.NewTimer(time.Hour)
	cache.timer.Stop()
	return
}