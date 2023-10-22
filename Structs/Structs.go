package Structs

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}
type DoubleLinkedList struct {
	Head *Node
}

func (l *DoubleLinkedList) Len() int {
	var len = 1
	for CurrentNode := l.Head; CurrentNode.next != nil; CurrentNode = CurrentNode.next {
		len += 1
	}
	return len
}
func (l *DoubleLinkedList) Add(value int) {
	LastNode := l.Head
	for CurrentNode := l.Head; CurrentNode != nil; CurrentNode = CurrentNode.next {
		LastNode = CurrentNode
	}
	LastNode.next = &Node{value: value, next: nil, prev: LastNode}
}
func New(count int) *DoubleLinkedList {
	Head := Node{value: 0, prev: nil, next: nil}
	CurrentNode := &Head
	for i := 0; i < count; i++ {
		CurrentNode.next = &Node{value: 0, prev: CurrentNode, next: nil}
		CurrentNode = CurrentNode.next
	}
	return &DoubleLinkedList{Head: &Head}
}
func NewFormList(arr []int) *DoubleLinkedList {
	var NewLinkedList = DoubleLinkedList{Head: &Node{value: arr[0], prev: nil, next: nil}}
	CurrentNode := NewLinkedList.Head
	for i := 1; i < len(arr); i++ {
		CurrentNode.next = &Node{value: arr[i], prev: CurrentNode, next: nil}
		CurrentNode = CurrentNode.next
	}
	return &NewLinkedList
}
func (l *DoubleLinkedList) Print() {
	for CurrentNode := l.Head; CurrentNode != nil; CurrentNode = CurrentNode.next {
		if CurrentNode.next != nil {
			fmt.Printf("%d <-> ", CurrentNode.value)
		} else {
			fmt.Printf("%d\n", CurrentNode.value)
		}
	}
}
func (l *DoubleLinkedList) Pop() {
	LastNode := l.Head
	for CurrentNode := l.Head; CurrentNode.next != nil; CurrentNode = CurrentNode.next {
		LastNode = CurrentNode.next
	}
	LastNode.prev.next = nil
}
func (l *DoubleLinkedList) At(index int) int {
	CurrentNode := l.Head
	for i := 0; i < index; i++ {
		CurrentNode = CurrentNode.next
	}
	return CurrentNode.value
}
func (l *DoubleLinkedList) DeleteFrom(index int) {
	CurrentNode := l.Head
	for i := 0; i < index; i++ {
		CurrentNode = CurrentNode.next
	}
	if CurrentNode.prev != nil {
		CurrentNode.prev.next = CurrentNode.next
	} else {
		l.Head = CurrentNode.next
		l.Head.prev = nil
	}
}
func (l *DoubleLinkedList) UpdateAt(index int, value int) {
	CurrentNode := l.Head
	for i := 0; i < index; i++ {
		CurrentNode = CurrentNode.next
	}
	CurrentNode.value = value
}
func (l *DoubleLinkedList) Insert(index int, value int) {
	CurrentNode := l.Head
	for i := 0; i < index; i++ {
		CurrentNode = CurrentNode.next
	}
	NewNode := &Node{value: value, prev: CurrentNode.prev, next: CurrentNode}
	if CurrentNode.prev != nil {
		CurrentNode.prev.next = NewNode
		CurrentNode.prev = NewNode
	} else {
		l.Head = NewNode
	}
}
func Join(l *DoubleLinkedList, m *DoubleLinkedList) *DoubleLinkedList{
	NewDoubleLinkedList := New(l.Len()+m.Len())
	CurrentNode1 := NewDoubleLinkedList.Head
	for CurrentNode2 := l.Head; CurrentNode2 != nil; CurrentNode2 = CurrentNode2.next{
		CurrentNode1.value = CurrentNode2.value
		CurrentNode1 = CurrentNode1.next
	}
	for CurrentNode2 := m.Head; CurrentNode2 != nil; CurrentNode2 = CurrentNode2.next{
		CurrentNode1.value = CurrentNode2.value
		CurrentNode1 = CurrentNode1.next
	}
	return NewDoubleLinkedList
}
