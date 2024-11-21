package syncmutex

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁
type MutexCounter struct {
	Ick   sync.Mutex
	Count int
}

func (c *MutexCounter) Add(wg *sync.WaitGroup) {
	// Your code here
	c.Ick.Lock()
	defer c.Ick.Unlock()

	for i := 0; i < 10000; i++ {
		c.Count += 1
	}
	println("addtime:", time.Now().UnixNano())

	wg.Done()
}

func (c *MutexCounter) Get() int {
	c.Ick.Lock()
	defer c.Ick.Unlock()

	ref := c.Count
	fmt.Println("mutexCounter", ref, "readtime:", time.Now().UnixNano())
	return ref

}

// 读写锁
type RWMutesCounter struct {
	lck   sync.RWMutex
	count int
	text  string
}

// 使用写锁（Lock）可以阻止其他的读操作和写操作
func (c *RWMutesCounter) Add(wg *sync.WaitGroup, text string) {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.text = text
	println("write_time:", time.Now().UnixNano())
	wg.Done()
}

// 使用读锁（RLock）可以允许多个并发读操作，但阻止写操作。
func (c *RWMutesCounter) Get(wg *sync.WaitGroup) int {
	c.lck.RLock()
	defer c.lck.RUnlock()
	wg.Done()

	ref := c.count
	println("struct_text:", c.text, "readtime:", time.Now().UnixNano())
	return ref
}