package main

import "sync"

type MutexCounter struct {
	lck   sync.Mutex //互斥锁
	count int
}

func (c *MutexCounter) Add() {
	// Your code here
	c.lck.Lock()
	c.count++
	defer c.lck.Unlock()
}

func (c *MutexCounter) Get() int {
	c.lck.Lock()
	ref := c.count
	defer c.lck.Unlock()

	return ref
}

type RWMutesCounter struct {
	lck   sync.RWMutex //读写锁
	count int
}

func (c *RWMutesCounter) Add() {
	// Your code here
}

func (c *RWMutesCounter) Get() int {
	// Your code here
	return 0
}
