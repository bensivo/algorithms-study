package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node

	// NOTE: if we added a pointer to the Tail, we could make Append O(1) instead of O(n). Addds a little complexity
}

func (l *LinkedList) Append(value int) {
	if l.Head == nil {
		l.Head = &Node{
			Value: value,
			Next:  nil,
		}
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = &Node{
		Value: value,
		Next:  nil,
	}
}

func (l *LinkedList) Prepend(value int) {
	if l.Head == nil {
		l.Head = &Node{
			Value: value,
			Next:  nil,
		}
		return
	}

	newHead := &Node{
		Value: value,
		Next:  l.Head,
	}

	l.Head = newHead
}

func (l *LinkedList) RemoveAt(index int) error {
	if l.Head == nil {
		return errors.New("list is empty")
	}

	if index == 0 {
		l.Head = l.Head.Next
		return nil
	}

	i := 0
	curr := l.Head
	for curr.Next != nil {
		if i == index-1 {
			if curr.Next != nil {
				curr.Next = curr.Next.Next
			} else {
				curr.Next = nil
			}
			return nil
		}

		curr = curr.Next
		i++
	}

	return errors.New("out of bounds")
}

func (l *LinkedList) Remove(value int) error {
	if l.Head == nil {
		return errors.New("value not found")
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return nil
	}

	curr := l.Head
	for curr.Next != nil {
		if curr.Next.Value == value {
			curr.Next = curr.Next.Next
			return nil
		}
		curr = curr.Next
	}
	return errors.New("value not found")
}

func (ll *LinkedList) ToArray() []int {
	arr := []int{}

	if ll.Head == nil {
		return arr
	}

	curr := ll.Head
	for curr.Next != nil {
		arr = append(arr, curr.Value)
		curr = curr.Next
	}
	arr = append(arr, curr.Value)

	return arr
}

func ReverseList(ll *LinkedList) {
	var prev *Node
	var curr *Node
	var next *Node

	curr = ll.Head
	for curr != nil {
		fmt.Println(curr.Value)
		next = curr.Next

		if prev != nil {
			fmt.Printf("Setting %d.Next = %d", curr.Value, prev.Value)
		}
		curr.Next = prev
		prev = curr
		curr = next
	}

	ll.Head = prev
}
