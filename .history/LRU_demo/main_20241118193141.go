package main

import (
	"./LRU_demo" // 导入你的 LRU 包
)

func main() {
	// 调用示例函数
	LRU_demo.Example()

	lru := LRU_demo.NewLRUCache(100)

	// 测试 MyNode
	myNode := &LRU_demo.MyNode{Data: "测试数据"}
	LRU_demo.UseNode(myNode)
}
