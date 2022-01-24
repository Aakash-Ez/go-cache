package main

import (
	"fmt"
	"time"
	"container/heap"
)


func main() {
	var cache1 = newCache()
	var mode, value, TTL int
	var key string
	var exists bool
	var item *Item
	var queueItem *QueueItem
	pq := make(PriorityQueue, len(cache1.Map))
	heap.Init(&pq)
	go pq.checkExpiry(&cache1)
	for {
		fmt.Println(time.Now().Unix())
		if len(pq) > 0 {
			for i := 0; i < pq.Len(); i++{
				fmt.Println(i)
				fmt.Println(pq[i])
			}
		}
		fmt.Println(pq)
		fmt.Println("Enter mode:\n 1. Set Value with default TTL\n 2. Set Value with TTL \n 3. Get Value \n 4. Delete Value \n 5. Print Map \n 6. Quit")
		fmt.Scanln(&mode)
		if mode == 1 {
			fmt.Print("Enter key, value: ")
			fmt.Scan(&key, &value)
			item, exists = cache1.set(Parameters{data: int64(value), key: key})
			if exists {
				fmt.Println(item.data, item.expirationTime)
				pq[item.index].item = item
				pq.update(item)
			} else {
				queueItem = createQueueItem(key,item)
				heap.Push(&pq, queueItem)
			}
		} else if mode == 2 {
			fmt.Print("Enter key, value, TTL: ")
			fmt.Scan(&key, &value, &TTL)
			item, exists = cache1.set(Parameters{data: int64(value), key: key, TTL: int64(TTL)})
			if exists {
				fmt.Println(item.data, item.expirationTime)
				pq[item.index].item = item
				pq.update(item)
			} else {
				queueItem = createQueueItem(key,item)
				heap.Push(&pq, queueItem)
			}
		} else if mode == 3 {
			fmt.Print("Enter key: ")
			fmt.Scan(&key)
			fmt.Println(cache1.get(key))
		} else if mode == 4 {
			fmt.Print("Enter key: ")
			fmt.Scan(&key)
			item = cache1.delete(key)
			heap.Remove(&pq, item.index)
		} else if mode == 5 {
			cache1.printMap()
		} else if mode == 6 {
			break
		}
	}
}
