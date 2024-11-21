package main

import "sync"

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
	// 读取结果
	println("count:", mutexCounterTest.Get())

}
