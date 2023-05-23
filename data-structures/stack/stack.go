package stack

import (
	"errors"

	"bensivo.com/leetcode/data-structures/linkedlist"
)

type Stack struct {
	// Many implementations of stacks use arrays, but we're using a linkedlist because insertion and removal (from head) is always constant time
	ll *linkedlist.LinkedList
}

func New() *Stack {
	return &Stack{
		ll: &linkedlist.LinkedList{},
	}
}

func (s *Stack) Push(v int) {
	s.ll.Prepend(v)
}

func (s *Stack) Pop() (int, error) {
	if s.ll.Head == nil {
		return 0, errors.New("stack is empty")
	}

	value := s.ll.Head.Value
	err := s.ll.RemoveAt(0)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Stack) ToArray() []int {
	return s.ll.ToArray()
}
