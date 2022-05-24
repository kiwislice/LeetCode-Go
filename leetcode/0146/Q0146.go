package leetcode

type LRUCache struct {
	limit    int
	cache    map[int]int
	keyOrder Queue
}

func Constructor(capacity int) LRUCache {
	rtn := LRUCache{limit: capacity, cache: make(map[int]int, capacity+1)}
	rtn.keyOrder = NewQueue()
	return rtn
}

func (lru *LRUCache) Get(key int) int {
	// fmt.Println("Get key=", key, ", m=", lru.cache)
	v, ok := lru.cache[key]
	if ok {
		lru.keyOrder.AddFirst(key)
		return v
	} else {
		return -1
	}
}

func (lru *LRUCache) Put(key int, value int) {
	lru.cache[key] = value
	lru.keyOrder.AddFirst(key)
	if len(lru.cache) > lru.limit {
		removeKey := lru.keyOrder.RemoveLast()
		delete(lru.cache, removeKey)
	}
	// fmt.Println("Put key=", key, ", m=", lru.cache)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type Queue struct {
	list        [10000 + 1]int
	recentIndex int
}

func NewQueue() Queue {
	q := Queue{}
	q.recentIndex = -1
	for i := range q.list {
		q.list[i] = -1
	}
	return q
}

func (queue *Queue) AddFirst(k int) {
	if queue.recentIndex == k {
		return
	}
	if queue.recentIndex == -1 {
		queue.recentIndex = k
		return
	}

	beforeK := queue.recentIndex
	for queue.list[beforeK] != k && queue.list[beforeK] != -1 {
		beforeK = queue.list[beforeK]
	}
	queue.list[beforeK] = queue.list[k]

	queue.list[k] = queue.recentIndex
	queue.recentIndex = k
}

func (queue *Queue) RemoveLast() int {
	if queue.recentIndex == -1 {
		return -1
	}
	if queue.list[queue.recentIndex] == -1 {
		rtn := queue.recentIndex
		queue.recentIndex = -1
		return rtn
	}

	beforeBeforeLast := queue.recentIndex
	beforeLast := queue.list[beforeBeforeLast]
	for queue.list[beforeLast] != -1 {
		beforeBeforeLast = beforeLast
		beforeLast = queue.list[beforeLast]
	}
	queue.list[beforeBeforeLast] = -1
	return beforeLast
}
