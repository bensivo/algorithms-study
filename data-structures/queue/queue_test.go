package queue_test

import (
	"testing"

	"bensivo.com/leetcode/data-structures/queue"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {

	s := queue.New()

	s.Push("A")

	assert.Equal(t, []string{"A"}, s.ToArray())

	s.Push("B")
	assert.Equal(t, []string{"A", "B"}, s.ToArray())
}

func TestPop(t *testing.T) {
	s := queue.New()

	s.Push("A")
	s.Push("B")
	s.Push("C")
	v, _ := s.Pop()

	assert.Equal(t, v, "A")
	assert.Equal(t, []string{"B", "C"}, s.ToArray())
}

func TestPop_Empty(t *testing.T) {
	s := queue.New()

	_, err := s.Pop()
	assert.EqualError(t, err, "queue is empty")
}
