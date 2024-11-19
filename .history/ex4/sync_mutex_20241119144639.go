package main

import "sync"

// 互斥锁
type MutexCounter struct {
	lck   sync.Mutex
	count int
}

func (c *MutexCounter) Add() {
	// Your code here
	c.lck.Lock()
	defer c.lck.Unlock()
	c.count++
}

func (c *MutexCounter) Get() int {
	c.lck.Lock()
	defer c.lck.Unlock()
	ref := c.count

	return ref
}

// 读写锁
type RWMutesCounter struct {
	lck   sync.RWMutex
	count int
}

func (c *RWMutesCounter) Add() {
	// Your code here
}

func (c *RWMutesCounter) Get() int {
	// Your code here
	return 0
}
