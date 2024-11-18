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
	cache    map[interface{}]*Node // 哈希表，用于O(1)时间复杂度的查找
	head     *Node                 // 双向链表的头节点（最近使用的）
	tail     *Node                 // 双向链表的尾节点（最久未使用的）
}

// NewLRUCache 创建一个新的LRU缓存
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[interface{}]*Node),
		head:     nil,
		tail:     nil,
	}
}

type IPNode interface {
	Load()
	Unload()
}
