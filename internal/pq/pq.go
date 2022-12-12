package pq

import "container/heap"

type _item[T any] struct {
	value    T
	priority int
}

type _pq[T any] []_item[T]

func (pq _pq[T]) Len() int { return len(pq) }

func (pq _pq[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq _pq[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *_pq[T]) Push(x any) {
	*pq = append(*pq, x.(_item[T]))
}

func (pq *_pq[T]) Pop() any {
	old := *pq
	item := old[len(old)-1]
	*pq = old[:len(old)-1]
	return item
}

type PriorityQueue[T any] struct {
	_pq[T]
	getPriority func(T) int
}

func (pq *PriorityQueue[T]) Push(item T) {
	heap.Push(&pq._pq, _item[T]{item, pq.getPriority(item)})
}
func (pq *PriorityQueue[T]) Pop() T {
	return heap.Pop(&pq._pq).(_item[T]).value
}
func (pq PriorityQueue[T]) Empty() bool {
	return len(pq._pq) == 0
}

func NewPriorityQueue[T any](getPriority func(T) int, items ...T) PriorityQueue[T] {
	pq := PriorityQueue[T]{getPriority: getPriority}
	for _, item := range items {
		pq.Push(item)
	}
	return pq
}
