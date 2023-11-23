package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l LinkedList) PrintList() {
	toPrint := l.Head
	for l.Length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		l.Length--
	}

	fmt.Printf("\n")
}

func (l *LinkedList) Prepend(node *Node) {
	second := l.Head
	l.Head = node
	l.Head.Next = second

	l.Length++
}

func (l *LinkedList) Delete(value int) {
	if l.Length == 0 {
		return
	}

	if l.Head.Data == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	toDelete := l.Head
	for toDelete.Next.Data != value {
		if toDelete.Next.Next == nil {
			return
		}

		toDelete = toDelete.Next
	}
	toDelete.Next = toDelete.Next.Next
	l.Length--
}

func main() {
	myLinkedList := LinkedList{}

	node1 := &Node{Data: 12}
	node2 := &Node{Data: 32}
	node3 := &Node{Data: 131}
	node4 := &Node{Data: 432}
	node5 := &Node{Data: 23}
	node6 := &Node{Data: 473}

	myLinkedList.Prepend(node1)
	myLinkedList.Prepend(node2)
	myLinkedList.Prepend(node3)
	myLinkedList.Prepend(node4)
	myLinkedList.Prepend(node5)
	myLinkedList.Prepend(node6)

	myLinkedList.PrintList()

	myLinkedList.Delete(432)
	myLinkedList.PrintList()

	myLinkedList.Delete(473)
	myLinkedList.PrintList()

	myLinkedList.Delete(100)
	myLinkedList.PrintList()
}
