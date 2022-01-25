package main

import (
	"container/heap"
	"fmt"
	"math"
	"time"
)

type Item struct {
	data           int64
	expirationTime int64
	TTL int 
	index int
}

type Cache struct {
	Map map[string]*Item
	timer *time.Timer
}

type Parameters struct {
	data int64
	key string
	TTL int64
}

func (c *Cache) set(parameter Parameters, pq *PriorityQueue) (item *Item, exists bool){
	if parameter.TTL == 0{
		parameter.TTL = 60
	} else if parameter.TTL < 0 {
		panic("ERR: TTL cannot be negative.")
	}
	_, exists = c.Map[parameter.key]
	expirationTime := time.Now().Unix() + parameter.TTL
	c.Map[parameter.key] = &Item{data: parameter.data, expirationTime: expirationTime, TTL: int(parameter.TTL)}
	item = c.Map[parameter.key]
	if (*pq).Len() == 0 {
		c.timer.Reset(time.Duration(parameter.TTL * int64(math.Pow(10,9))))
	} else if (*pq)[(*pq).Len() - 1].item.expirationTime > expirationTime{
		c.timer.Reset(time.Duration(parameter.TTL * int64(math.Pow(10,9))))
	}
	return
}

func (c *Cache) get(key string, pq *PriorityQueue) (value int64, item *Item) {
	var state bool
	item, state = c.Map[key]
	if !state {
		panic("ERR: Key not found.")
	}
	value = item.data
	item.expirationTime = time.Now().Unix() + int64(item.TTL)
	old := (*pq)[pq.Len() - 1]
	heap.Fix(pq, item.index)
	if old != (*pq)[pq.Len() - 1] {
		c.timer.Reset(time.Duration(((*pq)[pq.Len() - 1].item.expirationTime - time.Now().Unix()) * int64(math.Pow(10,9))))
	}
	return
}

func (c *Cache) delete(key string, pq *PriorityQueue) (item *Item) {
	var state bool
	item, state = c.Map[key]
	if !state {
		panic("ERR: Key not found.")
	}
	n := pq.Len()
	if pq.Len() == 0 {
		c.timer.Stop()
	} else if item.index == n - 1 || item.index == -1{
		c.timer.Reset(time.Duration(((*pq)[pq.Len() - 1].item.expirationTime - time.Now().Unix()) * int64(math.Pow(10,9))))
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