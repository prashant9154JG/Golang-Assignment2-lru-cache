package main

import (
	"fmt"
)

type LRU struct {
	keyStore    map[int]*Node
	valuesStore map[int][]int
	cacheList   List
	capacity    int
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type List struct {
	len  int
	head *Node
	tail *Node
}

func initList() *List {
	return &List{}
}

func (l *List) Front() *Node {
	return l.head
}

func (N *Node) Next() *Node {
	return N.next
}

func (l *List) Back() *Node {
	return l.tail
}

func (l *List) PushFront(data int) {
	if l.head == nil {
		newNode := &Node{
			value: data,
		}
		l.head = newNode
		l.tail = newNode
	} else {
		newNode := &Node{
			value: data,
		}
		l.tail.next = newNode
		newNode.prev = l.tail
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
	}
	l.len++
}

func (l *List) PushBack(data int) {
	if l.head == nil {
		newNode := &Node{
			value: data,
		}
		l.head = newNode
		l.tail = newNode
	} else {
		newNode := &Node{
			value: data,
		}
		l.tail.next = newNode
		newNode.prev = l.tail
		l.head.prev = newNode
		newNode.next = l.head
		l.tail = newNode
	}
	l.len++
}

func (l *List) MoveToFront(ptr *Node) {
	if ptr == l.head {
		return
	} else if ptr == l.tail {
		l.tail = ptr.prev
		ptr.prev.next = ptr.next
		ptr.next.prev = ptr.prev
		ptr.next = nil
		ptr.prev = nil
		l.PushFront(ptr.value)
	} else {
		ptr.prev.next = ptr.next
		ptr.next.prev = ptr.prev
		ptr.next = nil
		ptr.prev = nil
		l.PushFront(ptr.value)
	}
}

func (l *List) PopBack(ptr *Node) {
	l.tail = ptr.prev
	ptr.prev.next = ptr.next
	ptr.next.prev = ptr.prev
	ptr.next = nil
	ptr.prev = nil
	l.len--
}

func main() {
	var lru LRU
	lru.capacity = 5
	lru.cacheList = *initList()
	lru.keyStore = make(map[int]*Node)
	lru.valuesStore = make(map[int][]int)

	lru.Set(2, 3)
	lru.Set(3, 4)
	lru.Set(4, 5)
	lru.Set(5, 6)
	lru.Set(9, 10)
	lru.Set(8, 9)
	lru.Set(3, 7)

	fmt.Println(lru.Get(3))
	// fmt.Println(len(lru.keyStore))

	lru.PrintLRU()
}

func (c *LRU) Get(key int) []int {
	elementPointers, ok := c.keyStore[key]
	if !ok {
		return []int{-1}
	} else {
		// elementPointer := c.keyStore[key]
		// key := c.keyStore[key].value

		c.cacheList.MoveToFront(elementPointers)
		return c.valuesStore[key]
	}
}

func (c *LRU) PrintLRU() {
	size := c.cacheList.len - 1
	for e := c.cacheList.Front(); e != nil && size > 0; e = e.Next() {
		fmt.Println(e.value, " -> ", c.valuesStore[e.value])
		size--
	}
}

func (c *LRU) Set(key, value int) {
	elementPointer, ok := c.keyStore[key]
	if !ok {
		if len(c.keyStore) < c.capacity {
			c.cacheList.PushFront(key)
			c.keyStore[key] = c.cacheList.Front()
			c.valuesStore[key] = append(c.valuesStore[key], value)
		} else {
			lastPointer := c.cacheList.Back()
			delete(c.keyStore, lastPointer.value)
			delete(c.valuesStore, lastPointer.value)
			c.cacheList.PopBack(lastPointer)

			c.cacheList.PushFront(key)
			c.keyStore[key] = c.cacheList.Front()
			c.valuesStore[key] = append(c.valuesStore[key], value)
		}
	} else {
		// elementPointer := c.keyStore[key]
		c.cacheList.MoveToFront(elementPointer)
		c.keyStore[key] = c.cacheList.head
		c.valuesStore[key] = append(c.valuesStore[key], value)
	}
}
