package stack

import (
	"errors"

	"bensivo.com/leetcode/data-structures/linkedlist"
)

type Stack struct {
	items *linkedlist.LinkedList
}

func New() *Stack {
	return &Stack{
		items: &linkedlist.LinkedList{},
	}
}

func (s *Stack) Push(v int) {
	s.items.Prepend(v)
}

func (s *Stack) Pop() (int, error) {
	if s.items.Head == nil {
		return 0, errors.New("stack is empty")
	}

	value := s.items.Head.Value
	err := s.items.RemoveAt(0)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Stack) ToArray() []int {
	return s.items.ToArray()
}
