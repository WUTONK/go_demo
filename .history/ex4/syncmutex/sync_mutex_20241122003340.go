package syncmutex

import (
	"sync"
)

// 互斥锁
type MutexCounter struct {
	Lck   sync.Mutex
	Count int
}

func (c *MutexCounter) Add() {
	c.Lck.Lock()
	defer c.Lck.Unlock()
	c.Count += 1
}

func (c *MutexCounter) Get() int {
	c.Lck.Lock()
	defer c.Lck.Unlock()

	ref := c.Count
	return ref

}

// 读写锁
type RWMutesCounter struct {
	Lck   sync.RWMutex
	Count int
}

// 使用写锁（Lock）可以阻止其他的读操作和写操作
func (c *RWMutesCounter) Add(count int) {
	c.Lck.Lock()
	defer c.Lck.Unlock()
	c.Count = count
}

// 使用读锁（RLock）可以允许多个并发读操作，但阻止写操作。
func (c *RWMutesCounter) Get() int {
	c.Lck.RLock()
	defer c.Lck.RUnlock()

	ref := c.Count
	return ref
}
