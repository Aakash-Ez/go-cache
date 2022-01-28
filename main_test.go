package main

import (
	"testing"
	"container/heap"
	"time"
)

func TestSameTimeExpiry(t *testing.T) {
	cache := newCache()
	pq := make(PriorityQueue, len(cache.Map))
	heap.Init(&pq)
	go pq.checkExpiry(&cache)

	cache.set(Parameters{key: "A", data: 100, TTL: 4}, &pq)

	time.Sleep(2 * time.Second)
	cache.set(Parameters{key: "B", data: 19, TTL: 4}, &pq)

	time.Sleep(2 * time.Second)
	time.Sleep(2100 * time.Millisecond)
	if len(cache.Map) != 0 {
		cache.printMap()
		t.Error("Map is not empty")
	}
}

func TestResetValue(t *testing.T) {
	cache := newCache()
	pq := make(PriorityQueue, len(cache.Map))
	heap.Init(&pq)
	go pq.checkExpiry(&cache)

	cache.set(Parameters{key: "A", data: 100, TTL: 4}, &pq)

	time.Sleep(2 * time.Second)

	cache.set(Parameters{key: "A", data: 19, TTL: 4}, &pq)

	time.Sleep(2 * time.Second)

	if len(cache.Map) == 0 {
		cache.printMap()
		t.Error("Map is empty before item expiration.")
	}

	time.Sleep(2100 * time.Millisecond)

	if len(cache.Map) != 0 {
		cache.printMap()
		t.Error("Map is not empty.")
	}
}

func TestReorderingExpiration(t *testing.T) {
	cache := newCache()
	pq := make(PriorityQueue, len(cache.Map))
	heap.Init(&pq)
	go pq.checkExpiry(&cache)

	cache.set(Parameters{key: "A", data: 100, TTL: 2}, &pq)

	cache.set(Parameters{key: "B", data: 19, TTL: 4}, &pq)

	cache.set(Parameters{key: "A", data: 100, TTL: 6}, &pq)
	time.Sleep(100 * time.Millisecond)
	time.Sleep(4 * time.Second)

	_, exists := cache.Map["B"]

	if exists {
		cache.printMap()
		t.Error("Element not deleted.")
	}

	time.Sleep(2000 * time.Millisecond)

	_, exists = cache.Map["A"]

	if exists {
		cache.printMap()
		t.Error("Element not deleted.")
	}
}

func TestDeleteOperation(t *testing.T) {
	cache := newCache()
	pq := make(PriorityQueue, len(cache.Map))
	heap.Init(&pq)
	go pq.checkExpiry(&cache)

	cache.set(Parameters{key: "A", data: 100, TTL: 4}, &pq)

	cache.set(Parameters{key: "B", data: 19, TTL: 6}, &pq)

	time.Sleep(100 * time.Millisecond)
	time.Sleep(2 * time.Second)
	cache.printMap()
	_, item, _ := cache.get("A")
	heap.Remove(&pq, item.index)
	cache.delete("A", &pq)

	if len(cache.Map) != 1 {
		cache.printMap()
		t.Errorf("Expected 1 element in map. Found %d", len(cache.Map))
	}

	time.Sleep(2 * time.Second)

	if len(cache.Map) != 1 {
		cache.printMap()
		t.Errorf("Expected 1 element in map. Found %d", len(cache.Map))
	}

	time.Sleep(2100 * time.Millisecond)

	if len(cache.Map) != 0 {
		cache.printMap()
		t.Errorf("Expected 0 elements in map. Found %d", len(cache.Map))
	}

}

func TestGetFunction(t *testing.T) {
	cache := newCache()
	pq := make(PriorityQueue, len(cache.Map))
	heap.Init(&pq)
	go pq.checkExpiry(&cache)

	cache.set(Parameters{key: "A", data: 100, TTL: 4}, &pq)

	cache.set(Parameters{key: "B", data: 19, TTL: 6}, &pq)

	time.Sleep(2 * time.Second)

	_, _, statusA := cache.get("A")
	_, _, statusB := cache.get("B")

	if !(statusA && statusB) {
		t.Errorf("Expected 2 elements in map. Found %d", len(cache.Map))
	}

	time.Sleep(2100 * time.Millisecond)

	_, _, statusA = cache.get("A")
	_, _, statusB = cache.get("B")

	if !(!statusA && statusB) {
		t.Error("Invalid entries in the map.")
	}

	time.Sleep(2000 * time.Millisecond)

	_, _, statusA = cache.get("A")
	_, _, statusB = cache.get("B")

	if statusA || statusB {
		t.Error("Invalid entries in the map.")
	}
}