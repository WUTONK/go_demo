package main

import (
	"slices"
	"sync/atomic"
	"time"
)

type NS map[string]*PNode

type PNode struct {
	LastUseTime int64        // 最后使用时间
	Loaded      bool         // 是否加载
	LoadedTime  int64        // 载入时间
	FreqCount   atomic.Int64 // 计数器
}

func (p *PNode) Load() {
	p.Loaded = true
	p.LastUseTime = time.Now().Unix()
	p.LoadedTime = time.Now().Unix()
}

func (p *PNode) Unload() {
	p.Loaded = false
}

func LRU(ns NS, k int) {
	var times []int64 //时间列表

	// 遍历所有节点
	for _, node := range ns {

		// 如果节点未加载，则跳过
		if !node.Loaded {
			continue
		}

		// 将节点最后使用时间添加到时间列表中
		times = append(times, node.LastUseTime)
	}
	// 如果时间列表长度小于等于k，则返回（这意味着队列未满，无需卸载）
	if len(times) <= k {
		return
	}

	// 对时间列表进行排序
	slices.Sort(times)

	// 获取时间列表的最后一个元素
	seg := times[len(times)-k]

	// 遍历所有节点
	for _, node := range ns {

		// 如果节点最后使用时间小于等于seg，则跳过（这意味着节点未到被携带）
		if node.LastUseTime >= seg {
			continue
		}
		// 卸载节点
		node.Unload()
	}

}

func LFU(ns NS, k int) {

	var FreqCount_list []int64

	for _, node := range ns {
		//获取计数器
		FreqCount_list = append(FreqCount_list, node.FreqCount.Load())
	}

	slices.Sort(FreqCount_list)

	min_freq := FreqCount_list[k]

	for _, node := range ns {

		// 如果节点计数器大于等于min_freq，则跳过（这意味着节点未到被卸载点）
		if node.FreqCount.Load() >= min_freq {
			continue
		}
		// 卸载节点
		node.Unload()
	}

}

func FIFO(ns NS, k int) {

	var count []int64 //时间列表

	// 遍历所有节点
	for _, node := range ns {

		// 如果节点未加载，则跳过
		if !node.Loaded {
			continue
		}

		// 将节点载入时间添加到列表中
		count = append(count, node.LoadedTime)
	}

	// 如果载入列表长度小于等于k，则返回（这意味着队列未满，无需卸载）
	if len(count) <= k {
		return
	}

	//对载入顺序进行排序
	slices.Sort(count)

	// 获取时间列表的最后一个元素
	seg := count[len(count)-k]

	// 遍历所有节点
	for _, node := range ns {

		// 如果节点最后使用时间小于等于seg，则跳过（这意味着节点未到卸载点）
		if node.LastUseTime <= seg {
			continue
		}
		// 卸载节点
		node.Unload()
	}

}
