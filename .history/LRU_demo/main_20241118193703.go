package LRU_demo

import "fmt"

func main() {
	Example()
}

// Node 定义双向链表
type Node struct {
	Key   interface{} // 键
	Value interface{} // 值
	Prev  *Node       // 前驱节点
	Next  *Node       // 后继节点
}

// LRUCache 定义LRU缓存结构
type LRUCache struct {
	capacity int                   // 缓存容量
	cache    map[interface{}]*Node // 哈希表
	head     *Node                 // 双向链表的头节点（最近使用的）
	tail     *Node                 // 双向链表的尾节点（最久未使用的）
}

// NewLRUCache 创建一个新的LRU缓存
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,                    // 缓存容量
		cache:    make(map[interface{}]*Node), // 哈希表
		head:     nil,                         // 双向链表的头节点（最近使用的）
		tail:     nil,                         // 双向链表的尾节点（最久未使用的）
	}
}

type IPNode interface {
	Load()
	Unload()
}

type MyNode struct {
	Data string
}

func (node *MyNode) Load() {
	fmt.Println("load", node.Data)
}

func (node *MyNode) Unload() {
	fmt.Println("unload", node.Data)
}

func useNode(node IPNode) {
	node.Load()
	node.Unload()
}

func Example() {
	fmt.Println("开始运行Example函数")

	lru := NewLRUCache(100)
	fmt.Printf("创建的LRU缓存容量: %d\n", lru.capacity)

	node := &Node{Key: "hello", Value: "world"}
	lru.cache["hello"] = node
	foundNode := lru.cache["hello"]
	fmt.Printf("找到的节点: %+v\n", foundNode)
}
