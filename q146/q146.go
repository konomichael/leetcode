package q146

type LRUCache struct {
	head     *Node
	tail     *Node
	keyNode  map[int]*Node
	capacity int
	length   int
}

type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		panic("capacity should be greater than 0")
	}

	lru := LRUCache{
		capacity: capacity,
		length:   0,
		keyNode:  make(map[int]*Node, capacity),
		head:     new(Node),
		tail:     new(Node),
	}
	lru.head.next, lru.tail.prev = lru.tail, lru.head
	return lru
}

func (lru *LRUCache) Get(key int) int {
	// check if key in keyNode
	node, found := lru.keyNode[key]
	if !found {
		return -1
	}

	lru.kick(node)
	lru.enqueue(node)

	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	// check if key in keyNode
	node, found := lru.keyNode[key]
	if found {
		lru.kick(node)
		node.value = value
		lru.enqueue(node)
		return
	}

	if lru.length == lru.capacity {
		node = lru.head.next
		lru.kick(node)
		delete(lru.keyNode, node.key)
	} else {
		lru.length++
		node = new(Node)
	}

	node.key, node.value = key, value
	lru.enqueue(node)
	lru.keyNode[key] = node
}

func (lru *LRUCache) kick(node *Node) {
	node.prev.next, node.next.prev = node.next, node.prev
}

func (lru *LRUCache) enqueue(node *Node) {
	lru.tail.prev.next, node.prev, node.next, lru.tail.prev = node, lru.tail.prev, lru.tail, node
}
