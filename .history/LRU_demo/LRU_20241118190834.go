package LRU_demo

// Node 定义双向链表的节点结构
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

func Example() {

	// 假设我们有一个LRU缓存实例
	lru := NewLRUCache(100)

	// cache可以这样使用
	node := &Node{Key: "hello", Value: "world"}
	lru.cache["hello"] = node       // 存储键值对
	foundNode := lru.cache["hello"] // 获取值
	print(foundNode)

}
