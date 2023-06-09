package stack_test

import (
	"testing"

	"bensivo.com/leetcode/data-structures/stack"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {

	s := stack.New()

	s.Push("A")

	assert.Equal(t, []string{"A"}, s.ToArray())

	s.Push("B")
	assert.Equal(t, []string{"B", "A"}, s.ToArray())
}

func TestPop(t *testing.T) {
	s := stack.New()

	s.Push("B")
	s.Push("C")
	s.Push("D")
	v, _ := s.Pop()

	assert.Equal(t, v, "D")
	assert.Equal(t, []string{"C", "B"}, s.ToArray())
}

func TestPop_Empty(t *testing.T) {
	s := stack.New()

	_, err := s.Pop()
	assert.EqualError(t, err, "stack is empty")
}
