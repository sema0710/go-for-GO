package main

import "fmt"

type LData interface {}

type LQueue struct {
	front *LNode
	rear *LNode
	numOfData int
}

type LNode struct {
	data LData
	next *LNode
}

func NewLQueue() *LQueue {
	return &LQueue{
		front:     nil,
		rear:      nil,
		numOfData: 0,
	}
}

func NewLNode(data LData) *LNode {
	return &LNode{
		data: data,
		next: nil,
	}
}

func (pq *LQueue) QIsEmpty() bool {
	if pq.front == nil {
		return true
	}
	return false
}

func (pq *LQueue) Enqueue(data LData) {
	if data == nil {
		return
	}

	newNode := NewLNode(data)

	if pq.QIsEmpty() {
		pq.rear = newNode
		pq.front = newNode
	} else {
		pq.rear.next = newNode
		pq.rear = newNode
	}

	pq.numOfData++
}

func (pq *LQueue) Dequeue() LData {
	if pq.QIsEmpty() {
		return nil
	}

	rData := pq.front.data
	pq.front = pq.front.next
	pq.numOfData--
	return rData
}

func (pq *LQueue) QPrint() {
	fmt.Printf("현재 데이터의 갯수: %d\n", pq.numOfData)

	cur := pq.front
	for {
		fmt.Print(cur.data, " ")
		if cur.next == nil {
			break
		}
		cur = cur.next
	}
}
