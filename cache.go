package main

import (
	"time"
)

type Item struct {
	data           int64
	expirationTime int64
	index int
}

type Cache struct {
	Map map[string]*Item
}

type Parameters struct {
	data int64
	key string
	TTL int64
}

func (c *Cache) set(parameter Parameters) (item *Item, exists bool){
	if parameter.TTL == 0{
		parameter.TTL = 60
	} else if parameter.TTL < 0 {
		panic("ERR: TTL cannot be negative.")
	}
	_, exists = c.Map[parameter.key]
	c.Map[parameter.key] = &Item{data: parameter.data, expirationTime: time.Now().Unix() + parameter.TTL}
	item = c.Map[parameter.key]
	return
}

func (c *Cache) get(key string) (value int64) {
	var state bool
	var item *Item
	item, state = c.Map[key]
	value = item.data
	if !state {
		panic("ERR: Key not found.")
	}
	return
}

func (c *Cache) delete(key string) (item *Item) {
	var state bool
	item, state = c.Map[key]
	if !state {
		panic("ERR: Key not found.")
	}
	delete(c.Map, key)
	return
}

func newCache() (cache Cache) {
	cache = Cache{Map: make(map[string]*Item)}
	return
}