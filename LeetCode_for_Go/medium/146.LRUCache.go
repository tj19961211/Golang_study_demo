package medium

// 节点结构
type DlinkedList struct {
	key  int
	val  int
	prev *DlinkedList
	next *DlinkedList
}

// LRUCache结构
type LRUCache struct {
	capacity int // 长度
	head     *DlinkedList
	tail     *DlinkedList
	cache    map[int]*DlinkedList
}

// 生成一个 LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DlinkedList),
	}
}

// 获取 LRUCache 的值，并更新查询值在 LRUCache 的位置
func (c *LRUCache) Get(key int) int {
	l, ok := c.cache[key]
	if !ok {
		return -1
	}

	c.removeFromChain(l)
	c.addToChain(l)
	return l.val
}

// 添加值到 LRUCache 中
func (c *LRUCache) Put(key int, value int) {
	// 判断数据是否存在的情况
	l, ok := c.cache[key]
	if !ok {
		l = &DlinkedList{key: key, val: value}
		c.cache[key] = l
	} else {
		l.val = value
		c.removeFromChain(l)
	}
	c.addToChain(l)
	if len(c.cache) > c.capacity {
		l := c.tail
		c.removeFromChain(l)
		delete(c.cache, l.key)
	}
}

func (c *LRUCache) addToChain(l *DlinkedList) {
	l.next = nil
	if c.head != nil {
		c.head.next = l
		l.prev = c.head
	}
	c.head = l
	if c.tail == nil {
		c.tail = l
	}
}

// 因为使用的是头插法，所以 next 的数插入的时间比 prev 的早
func (c *LRUCache) removeFromChain(l *DlinkedList) {
	if l == c.head {
		c.head = l.prev
	}
	if l == c.tail {
		c.tail = l.next
	}
	// 删除链表中的 l
	if l.next != nil {
		l.next.prev = l.prev
	}
	if l.prev != nil {
		l.prev.next = l.next
	}
}
