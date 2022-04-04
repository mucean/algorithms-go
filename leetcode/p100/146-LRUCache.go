package p100

type lruNode struct {
	key  int
	val  int
	pre  *lruNode
	next *lruNode
}

type LRUCache struct {
	cap      int
	len      int
	keyNode  map[int]*lruNode
	headNode *lruNode
	endNode  *lruNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cap: capacity, keyNode: make(map[int]*lruNode, capacity)}
}

func (c *LRUCache) Get(key int) int {
	if n, ok := c.keyNode[key]; ok {
		c.updateNode(n)
		return n.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	// not exist
	// new, put first, update keyNode
	n, ok := c.keyNode[key]
	if ok == false {
		n = &lruNode{key: key, val: value}
		c.keyNode[key] = n
		if c.headNode == nil {
			c.len++
			c.headNode = n
			c.endNode = n
			return
		}
		c.headNode.pre = n
		n.next = c.headNode
		c.headNode = n
		// if cap not empty
		if c.cap > c.len {
			c.len++
		} else {
			delete(c.keyNode, c.endNode.key)
			c.endNode.pre.next = nil
			c.endNode = c.endNode.pre
		}
		return
	}

	// exist
	// update value, put first, update keyNode
	n.val = value
	c.updateNode(n)
}

func (c *LRUCache) updateNode(n *lruNode) {
	if n == c.headNode {
		return
	}

	n.pre.next = n.next
	if n.next != nil {
		n.next.pre = n.pre
	} else {
		c.endNode = n.pre
	}
	n.next = c.headNode
	n.pre = nil
	c.headNode.pre = n
	c.headNode = n
}
