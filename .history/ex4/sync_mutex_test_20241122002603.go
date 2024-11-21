package main

import (
	"sync"
	"testing"

	syn "go_demo/ex4/syncmutex"
)

func TestMutexCounterAdd(t *testing.T) {

	mutexCounterTest := syn.MutexCounter{
		Count: 0,
	}

	// 协程池
	var wg sync.WaitGroup
	wg.Add(3)

	//并发 3 个 add
	go mutexCounterTest.Add()
	go mutexCounterTest.Add()
	go mutexCounterTest.Add()

	wg.Wait()
	// 读取结果
	println("count:", mutexCounterTest.Get())

}
