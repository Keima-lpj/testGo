package algorithm

/**
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
进阶:
你是否可以在 O(1) 时间复杂度内完成这两种操作？
示例:
LRUCache cache = new LRUCache( 2 ) //缓存容量
cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4


分析上面的操作过程，要让 put 和 get 方法的时间复杂度为 O(1)O(1)，我们可以总结出 cache 这个数据结构必要的条件：查找快，插入快，删除快，有顺序之分。
因为显然 cache 必须有顺序之分，以区分最近使用的和久未使用的数据；而且我们要在 cache 中查找键是否已存在；如果容量满了要删除最后一个数据；每次访问还要把数据插入到队头。
那么，什么数据结构同时符合上述条件呢？哈希表查找快，但是数据无固定顺序；链表有顺序之分，插入删除快，但是查找慢。所以结合一下，形成一种新的数据结构：哈希链表。
LRU 缓存算法的核心数据结构就是哈希链表，双向链表和哈希表的结合体。
*/
//链表节点
type LRUNode struct {
	key  int
	val  int
	next *LRUNode
	prev *LRUNode
}

//插入节点数据
func (n *LRUNode) put(key, val int) {
	n.key = key
	n.val = val
}

//双向链表
type DoubleLRUNode struct {
	head *LRUNode //头部虚节点
	tail *LRUNode //尾部虚节点
	size int      //链表元素个数
}

//新建一个双向链表
func DoubleList() *DoubleLRUNode {
	head := &LRUNode{key: 0, val: 0}
	tail := &LRUNode{key: 0, val: 0}
	head.next = tail
	tail.prev = head
	return &DoubleLRUNode{
		head: head,
		tail: tail,
		size: 0,
	}
}

//在双向链表的头部添加节点
func (double *DoubleLRUNode) addFirst(node *LRUNode) {
	node.next = double.head.next
	node.prev = double.head
	double.head.next.prev = node
	double.head.next = node
	double.size++
}

//删除链表中的某个节点（这个节点一定存在）
func (double *DoubleLRUNode) remove(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	double.size--
}

//删除链表中的最后一个节点，并返回该节点
func (double *DoubleLRUNode) removeLast() *LRUNode {
	if double.tail.prev == double.head {
		return nil
	}
	//删除最后一个元素
	last := double.tail.prev
	double.remove(last)
	return last
}

//返回链表的长度
func (double *DoubleLRUNode) getSize() int {
	return double.size
}

//LRU缓存
type LRUCache struct {
	cap   int
	M     map[int]*LRUNode //m1存储键对应的节点
	cache *DoubleLRUNode   //cache存储双向链表的指针
}

//新建一个LRU缓存
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		M:     make(map[int]*LRUNode),
		cache: DoubleList(),
	}
}

func (this *LRUCache) Get(key int) int {
	//查询是否存在，如果存在，返回节点的值，否则返回-1
	if node, ok := this.M[key]; !ok {
		return -1
	} else {
		this.Put(node.key, node.val)
		return node.val
	}
}

func (this *LRUCache) Put(key int, value int) {
	node := &LRUNode{
		key: key,
		val: value,
	}
	if _, ok := this.M[key]; ok {
		//如果这个键原本就存在，那么删除旧的元素，将新元素插入到队列头部
		this.cache.remove(this.M[key])
		this.cache.addFirst(node)
		//更新map中对应的数据
		this.M[key] = node
	} else {
		//如果这个键本身不存在，那么则需要判断容量是否超了
		if this.cap == this.cache.size {
			//超了就删除队列的最后一个元素
			last := this.cache.removeLast()
			delete(this.M, last.key)
		}
		//将新的元素添加到头部，并更新map
		this.cache.addFirst(node)
		this.M[key] = node
	}
}
