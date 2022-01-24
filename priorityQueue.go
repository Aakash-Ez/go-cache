package main

import (
	"container/heap"
	"time"
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
	item := old[0]
	old[n-1] = nil  
	item.item.index = -1 
	if n > 1 {
		*pq = old[1:n]
	} else {
		*pq = make(PriorityQueue, 0)
	}
	return item
}

func (pq *PriorityQueue) update(item *Item) {
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) checkExpiry(cache *Cache){
	for {
		if (*pq).Len() > 0 && (*pq)[0].item.expirationTime < time.Now().Unix() {
			queueItem := heap.Pop(pq).(*QueueItem)
			cache.delete(queueItem.key)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
