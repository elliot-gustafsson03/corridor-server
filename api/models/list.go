package models

type Node struct {
	Value Image
	Next  *Node
}

type List struct {
	Head        *Node
	currentNode *Node
	length      int
}

func (l List) IsEmpty() bool {
	return l.length == 0
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

	l.length++
}

func (l *List) NextValue() *Image {
	if l.IsEmpty() {
		return nil
	}

	l.currentNode = l.currentNode.Next
	return &l.currentNode.Value
}

func (l *List) Delete(id string) {
	if l.length == 1 {
		*l = List{}
	} else {
		currentNode := l.Head

		for i := 0; i < l.length; i++ {
			if currentNode.Next.Value.Name == id {
				if l.currentNode == currentNode.Next {
					l.currentNode = l.currentNode.Next
				}
				if currentNode.Next == l.Head {
					l.Head = l.Head.Next
				}

				currentNode.Next = currentNode.Next.Next
				break
			}

			currentNode = currentNode.Next
		}
	}

	l.length--
}
