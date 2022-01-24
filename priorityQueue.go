package main

import (
	"container/heap"
	"fmt"
)

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].expirationTime < pq[j].expirationTime
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  
	item.index = -1 
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int64, expirationTime int64) {
	item.data = value
	item.expirationTime = expirationTime
	fmt.Println(item.data, item.expirationTime, item.index)
	heap.Fix(pq, item.index)
}
