package main

import (
	"sync"
	"testing"

	syn "go_demo/ex4/syncmutex"
)

func TestMutexCounterAdd(t *testing.T) {

	var mutexCounterTest = syn.MutexCounter{
		Lck:   sync.Mutex{},
		Count: 0,
	}

	// 协程池

	//并发 3 个 add
	go mutexCounterTest.Add()
	go mutexCounterTest.Add()
	go mutexCounterTest.Add()

	// 读取结果
	println("count:", mutexCounterTest.Get())

}