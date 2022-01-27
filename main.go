package main

import (
	"fmt"
	"time"
	"container/heap"
)


func main() {
	var cache1 = newCache() //local instance of cache structure
	var mode, value, TTL int
	var key string 
	var item *Item
	pq := make(PriorityQueue, len(cache1.Map)) // Priority Queue with lowest expiration time having highest priority
	heap.Init(&pq) // Initializing Priority queue using Heap data structure
	go pq.checkExpiry(&cache1) // go routine to check if item has expired and remove it
	for {
		fmt.Println(time.Now().Unix())
		pq.printList()
		fmt.Println("Enter mode:\n 1. Set Value with default TTL\n 2. Set Value with TTL \n 3. Get Value \n 4. Delete Value \n 5. Print Map \n 6. Quit")
		fmt.Scanln(&mode)
		if mode == 1 { // Set value with default TTL (60 seconds)
			fmt.Print("Enter key, value: ")
			fmt.Scan(&key, &value)
			cache1.set(Parameters{data: int64(value), key: key}, &pq)
		} else if mode == 2 { // Set value with user-defined TTL
			fmt.Print("Enter key, value, TTL: ")
			fmt.Scan(&key, &value, &TTL)
			cache1.set(Parameters{data: int64(value), key: key, TTL: int64(TTL)}, &pq)
		} else if mode == 3 { // get value from the cache
			fmt.Print("Enter key: ")
			fmt.Scan(&key)
			value, _ := cache1.get(key)
			fmt.Println(value)
		} else if mode == 4 { // Delete key-value pair from cache
			fmt.Print("Enter key: ")
			fmt.Scan(&key)
			heap.Remove(&pq, item.index)
			item = cache1.delete(key, &pq)
		} else if mode == 5 { // Print the cache
			cache1.printMap()
		} else if mode == 6 { // Quit the Program
			break
		}
	}
}
