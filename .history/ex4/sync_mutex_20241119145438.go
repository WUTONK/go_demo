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

	for i := 0; i < 10000; i++ {
		c.count += 1
	}
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

func main() {

	MutexCounterTest := MutexCounter{
		count: 0,
	}

	//并发 3 个 add
	MutexCounterTest.Add()
	MutexCounterTest.Add()
	MutexCounterTest.Add()
	//输出结果
	println(MutexCounterTest.Get())

}
