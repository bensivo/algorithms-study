package linkedlist_test

import (
	"testing"

	"bensivo.com/leetcode/data-structures/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	assert.Equal(t, []string{}, ll.ToArray())

	ll.Append("A")
	assert.Equal(t, []string{"A"}, ll.ToArray())

	ll.Append("B")
	assert.Equal(t, []string{"A", "B"}, ll.ToArray())
}

func TestPrepend(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	assert.Equal(t, []string{}, ll.ToArray())

	ll.Prepend("A")
	assert.Equal(t, []string{"A"}, ll.ToArray())

	ll.Prepend("B")
	assert.Equal(t, []string{"B", "A"}, ll.ToArray())
}

func TestRemove_Empty(t *testing.T) {
	ll := &linkedlist.LinkedList{}

	err := ll.Remove("A")
	assert.EqualError(t, err, "value not found")
}

func TestRemove_RemoveHead(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append("A")

	ll.Remove("A")
	assert.Equal(t, []string{}, ll.ToArray())
}

func TestRemove_RemoveMiddle(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	ll.Remove("B")
	assert.Equal(t, []string{"A", "C"}, ll.ToArray())
}

func TestRemove_RemoveEnd(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	ll.Remove("C")
	assert.Equal(t, []string{"A", "B"}, ll.ToArray())
}

func TestRemove_NotFound(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	err := ll.Remove("D")
	assert.EqualError(t, err, "value not found")
	assert.Equal(t, []string{"A", "B", "C"}, ll.ToArray())
}

func TestRemoveAt_Begining(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	ll.RemoveAt(0)
	assert.Equal(t, []string{"B", "C"}, ll.ToArray())
}

func TestRemoveAt_Middle(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	ll.RemoveAt(1)
	assert.Equal(t, []string{"A", "C"}, ll.ToArray())
}

func TestRemoveAt_End(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	ll.RemoveAt(2)
	assert.Equal(t, []string{"A", "B"}, ll.ToArray())
}

func TestRemoveAt_OutOfBounds(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	err := ll.RemoveAt(5)
	assert.EqualError(t, err, "out of bounds")
}

func TestReverse(t *testing.T) {
	ll := &linkedlist.LinkedList{}

	ll.Append("A")
	ll.Append("B")
	ll.Append("C")

	linkedlist.ReverseList(ll)

	assert.Equal(t, []string{"C", "B", "A"}, ll.ToArray())
}
