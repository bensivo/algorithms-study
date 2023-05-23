package stack_test

import (
	"testing"

	"bensivo.com/leetcode/data-structures/stack"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {

	s := stack.New()

	s.Push(0)

	assert.Equal(t, []int{0}, s.ToArray())

	s.Push(1)
	assert.Equal(t, []int{1, 0}, s.ToArray())
}

func TestPop(t *testing.T) {
	s := stack.New()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	v, _ := s.Pop()

	assert.Equal(t, v, 3)
	assert.Equal(t, []int{2, 1}, s.ToArray())
}

func TestPop_Empty(t *testing.T) {
	s := stack.New()

	_, err := s.Pop()
	assert.EqualError(t, err, "stack is empty")
}
