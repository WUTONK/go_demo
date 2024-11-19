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
	text  string
}

// 使用写锁（Lock）可以阻止其他的读操作和写操作
func (c *RWMutesCounter) Add(wg *sync.WaitGroup) {
	c.lck.Lock()
	defer c.lck.Unlock()

}

// 使用读锁（RLock）可以允许多个并发读操作，但阻止写操作。
func (c *RWMutesCounter) Get() int {
	c.lck.RLock()

	return 0
}

func main() {

	mutexCounterTest := MutexCounter{
		count: 0,
	}

	// 协程池
	var wg sync.WaitGroup
	wg.Add(3)
	//并发 3 个 add
	go mutexCounterTest.Add(&wg)
	go mutexCounterTest.Add(&wg)
	go mutexCounterTest.Add(&wg)
	wg.Wait()
	//输出结果
	fmt.Println("mutexCounterTest:", mutexCounterTest.Get())

	//并发 3 个 读
	wg.Add(3)

}
