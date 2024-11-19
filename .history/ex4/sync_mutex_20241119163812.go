package main

import (
	"fmt"
	"sync"
)

// 互斥锁
type MutexCounter struct {
	lck   sync.Mutex
	count int
}
type MutexCounterResult struct {
	count int
}

func (c *MutexCounter) Add(wg *sync.WaitGroup) {
	// Your code here
	c.lck.Lock()
	defer c.lck.Unlock()

	for i := 0; i < 10000; i++ {
		c.count += 1
	}

	wg.Done()
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

type RWMutesCounterResult struct {
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

	// 协程池
	var wg sync.WaitGroup
	wg.Add(3)
	//并发 3 个 add
	go MutexCounterTest.Add(&wg)
	go MutexCounterTest.Add(&wg)
	go MutexCounterTest.Add(&wg)
	wg.Wait()
	//输出结果
	fmt.Println("MutexCounterTest", MutexCounterTest.Get())

}
