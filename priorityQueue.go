package main

import (
	"container/heap"
	"fmt"
)

type QueueItem struct {
	key string
	item *Item
}

func createQueueItem(key string, item *Item)(queueItem *QueueItem) {
	queueItem = &QueueItem{key: key, item: item}
	return
}

type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].item.expirationTime < pq[j].item.expirationTime
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].item.index = i
	pq[j].item.index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QueueItem)
	item.item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  
	item.item.index = -1 
	*pq = old[0:n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item) {
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) checkExpiry(cache *Cache){
	for range cache.timer.C {
		cache.m.Lock()
		queueItem := heap.Pop(pq).(*QueueItem)
		cache.delete(queueItem.key, pq)
		cache.m.Unlock()
	}
}

func (pq *PriorityQueue) printList() {
	for i := 0; i < pq.Len(); i++ {
		fmt.Println((*pq)[i].item.index, (*pq)[i].item.expirationTime, (*pq)[i].key)
	}
}