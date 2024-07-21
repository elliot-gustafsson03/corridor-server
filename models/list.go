package models

type Node struct {
	Value Image
	Next  *Node
}

type List struct {
	Head        *Node
	currentNode *Node
}

func (l List) IsEmpty() bool {
	return l.Head == nil
}

func (l *List) Insert(value Image) {
	newNode := &Node{Value: value}

	if l.IsEmpty() {
		l.Head = newNode
		newNode.Next = newNode
		l.currentNode = newNode
	} else {
		newNode.Next = l.currentNode.Next
		l.currentNode.Next = newNode
	}
}

func (l *List) NextValue() *Image {
	if l.IsEmpty() {
		return nil
	}

	l.currentNode = l.currentNode.Next
	return &l.currentNode.Value
}
