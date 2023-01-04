package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

	var c Duck = &Cat{Name: "draven"}
	c.Quack()
}

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//go:noinline
func (c *Cat) Quack() {
	println(c.Name + " meow")
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
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
