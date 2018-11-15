package main

import (
	"fmt"
)

type Node struct {
	//Value interface{}
	Value int
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type SinglyLinkedList struct {
	front  *Node
	length int
}

//init
func (s *SinglyLinkedList) Init() *SinglyLinkedList {
	s.length = 0
	return s
}

//new一个链表
func New() *SinglyLinkedList {
	return new(SinglyLinkedList).Init()
}

//返回头结点
func (s *SinglyLinkedList) Front() *Node {
	return s.front
}

//返回尾节点
func (s *SinglyLinkedList) Back() *Node {
	currentNode := s.front
	for currentNode != nil && currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

//添加尾节点
func (s *SinglyLinkedList) Append(n *Node) {
	if s.front == nil {
		s.front = n
	} else {
		currentNode := s.front

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = n
	}
	s.length++
}

//添加头节点
func (s *SinglyLinkedList) Prepend(n *Node) {
	if s.front == nil {
		s.front = n
	} else {
		n.next = s.front
		s.front = n
	}
	s.length++
}

//在指定节点前添加节点
func (s *SinglyLinkedList) InsertBefore(insert *Node, before *Node) {
	if s.front == before {
		insert.next = before
		s.front = insert
		s.length++
	} else {
		currentNode := s.front
		for currentNode != nil && currentNode.next != nil && currentNode != before {
			currentNode = currentNode.next
		}
		if currentNode.next == before {
			insert.next = before
			currentNode.next = insert
			s.length++
		}
	}
}

//在指定节点后添加节点
func (s *SinglyLinkedList) InsertAfter(insert *Node, after *Node) {
	currentNode := s.front
	for currentNode != nil && currentNode != after && currentNode.next != nil {
		currentNode = currentNode.next
	}

	if currentNode == after {
		insert.next = after.next
		currentNode.next = insert
		s.length++
	}
}

//删除指定节点
func (s *SinglyLinkedList) Remove(n *Node) {
	if s.front == n {
		s.front = n.next
		s.length--
	} else {
		currentNode := s.front

		//search for node n
		for currentNode != nil && currentNode.next != nil && currentNode != n {
			currentNode = currentNode.next
		}

		//see if current's next node is n
		//if it's not n, then node n wasn't found int list s
		if currentNode.next == n {
			currentNode.next = currentNode.next.next
			s.length--
		}
	}
}

func main() {
	//创建一个链表
	alist := New()
	fmt.Println("list length:", alist.length)

	node1 := new(Node)
	node1.Value = 1

	alist.Append(node1)
	fmt.Println("list length:", alist.length)

	node2 := new(Node)
	node2.Value = 2
	node1.next = node2

	alist.Append(node2)
	fmt.Println("list.length:", alist.length)
}
