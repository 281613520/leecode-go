package main

import (
	"fmt"
)

func main() {
	fmt.Print(111111)
	obj := Constructor(3)
	obj.EnQueue(1)
	obj.EnQueue(2)
	obj.EnQueue(3)
	obj.EnQueue(4)
	obj.Rear()
	obj.IsFull()
	obj.DeQueue()
	obj.EnQueue(4)
	obj.Rear()

}

type MyCircularQueue struct {
	queue       []int
	front, rear int
	full        bool
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{queue: make([]int, k), front: 0, rear: 0, full: false}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.rear] = value

	this.rear = (this.rear + 1) % len(this.queue)

	if this.rear == this.front {
		this.full = true
	}

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.front = (this.front + 1) % len(this.queue)
	this.full = false

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[(this.rear-1+len(this.queue))%len(this.queue)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.rear == this.front && !this.full
}

func (this *MyCircularQueue) IsFull() bool {
	return this.full
}
