package main

type Node struct {
	Key   int
	Value int
	pre   *Node
	next  *Node
}

// 链表结构，限制长度、创建指针指向头节点与尾节点
type LRUcache struct {
	limit   int
	HashMap map[int]*Node
	head    *Node
	end     *Node
}

func Constructor(capacity int) *LRUcache {
	lrucache := &LRUcache{limit: capacity}
	lrucache.HashMap = make(map[int]*Node)
	return lrucache
}

func (l *LRUcache) Get(key int) int {
	if v, ok := l.HashMap[key]; ok {
		l.refreshNode(v)
		return v.Value
	} else {
		return -1
	}
}

func (l *LRUcache) Put(key int, value int) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.limit {
			oldKey := l.removeNode(l.head)
			delete(l.HashMap, oldKey)
		}
		node := &Node{Key: key, Value: value}
		l.addNode(node)
		l.HashMap[key] = node
	} else {
		v.Value = value
		l.refreshNode(v)
	}
}

func (l *LRUcache) refreshNode(node *Node) {
	if node == l.end {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

func (l *LRUcache) removeNode(node *Node) int {
	if node == l.end {
		l.end = l.end.pre
		l.end.next = nil
	} else if node == l.head {
		l.head = l.head.next
		l.head.pre = nil
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	return node.Key
}

func (l *LRUcache) addNode(node *Node) {
	if l.end != nil {
		l.end.next = node
		node.pre = l.end
		node.next = nil
	}
	l.end = node
	if l.head == nil {
		l.head = node
	}
}
