package main

import (
	"fmt"
)

// LRU is a structure of LRU cache
type LRU struct {
	keyStore    map[int]*Node
	valuesStore map[int][]int
	cacheList   List
	capacity    int
}

// Node is structure of single element of doubly linkedlist
type Node struct {
	value int
	next  *Node
	prev  *Node
}

// List is a structure of List
type List struct {
	len  int
	head *Node
	tail *Node
}

// initList initialise the doubly linked List
func initList() *List {
	return &List{}
}

// Front returns the first node of the list
func (l *List) Front() *Node {
	return l.head
}

// Next return the next node of the node
func (N *Node) Next() *Node {
	return N.next
}

// Back return node at the end of the list
func (l *List) Back() *Node {
	return l.tail
}

// PushFront push the new node with given data at the front of the list
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

// PushPack push the new node at the end of the list
// func (l *List) PushBack(data int) {
// 	if l.head == nil {
// 		newNode := &Node{
// 			value: data,
// 		}
// 		l.head = newNode
// 		l.tail = newNode
// 	} else {
// 		newNode := &Node{
// 			value: data,
// 		}
// 		l.tail.next = newNode
// 		newNode.prev = l.tail
// 		l.head.prev = newNode
// 		newNode.next = l.head
// 		l.tail = newNode
// 	}
// 	l.len++
// }

// MoveFront shift ptr node in the list to the front of the doubly linkedlist
func (l *List) MoveToFront(ptr *Node) {
	if ptr == l.head {
		return
	} else if ptr == l.tail {
		l.tail = ptr.prev
		// removing ptr pointer
		ptr.prev.next = ptr.next
		ptr.next.prev = ptr.prev

		// adding ptr pointer in front
		ptr.next = l.head
		l.head.prev = ptr
		l.head = ptr
		l.tail.next = l.head
		l.head.prev = l.tail
	} else {
		// removing ptr pointer
		ptr.prev.next = ptr.next
		ptr.next.prev = ptr.prev

		// adding ptr pointer in front
		ptr.next = l.head
		l.head.prev = ptr
		l.head = ptr
		l.tail.next = l.head
		l.head.prev = l.tail
	}
}

// PopBack remove the last node of the list
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
	lru.Set(7, 8)

	fmt.Println(lru.Get(3))
	lru.PrintLRU()
}

func (c *LRU) Get(key int) []int {
	elementPointers, ok := c.keyStore[key]
	if !ok {
		return []int{-1}
	} else {
		c.cacheList.MoveToFront(elementPointers)
		c.keyStore[key] = c.cacheList.head
		return c.valuesStore[key]
	}
}

func (c *LRU) PrintLRU() {
	size := c.cacheList.len
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
