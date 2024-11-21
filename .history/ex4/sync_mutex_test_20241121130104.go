package main

import (
	syn "go_demo/ex4/syncmutex"
	"sync"
	"testing"
)

func TestMutexCounterAdd(t *testing.T) {

	mutexCounterTest := syn.MutexCounter{
		Count: 0,
	}

	// 协程池
	var wg sync.WaitGroup
	wg.Add(3)

	//并发 3 个 add
	go mutexCounterTest.Add(&wg)
	go mutexCounterTest.Add(&wg)
	go mutexCounterTest.Add(&wg)

	wg.Wait()
	// 读取结果
	println("count:", mutexCounterTest.Get())

}