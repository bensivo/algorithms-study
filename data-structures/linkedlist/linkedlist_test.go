package linkedlist_test

import (
	"testing"

	"bensivo.com/leetcode/data-structures/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	assert.Equal(t, []int{}, ll.ToArray())

	ll.Append(0)
	assert.Equal(t, []int{0}, ll.ToArray())

	ll.Append(1)
	assert.Equal(t, []int{0, 1}, ll.ToArray())
}

func TestPrepend(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	assert.Equal(t, []int{}, ll.ToArray())

	ll.Prepend(0)
	assert.Equal(t, []int{0}, ll.ToArray())

	ll.Prepend(1)
	assert.Equal(t, []int{1, 0}, ll.ToArray())
}

func TestRemove_Empty(t *testing.T) {
	ll := &linkedlist.LinkedList{}

	err := ll.Remove(0)
	assert.EqualError(t, err, "value not found")
}

func TestRemove_RemoveHead(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append(0)

	ll.Remove(0)
	assert.Equal(t, []int{}, ll.ToArray())
}

func TestRemove_RemoveMiddle(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append(0)
	ll.Append(1)
	ll.Append(2)

	ll.Remove(1)
	assert.Equal(t, []int{0, 2}, ll.ToArray())
}

func TestRemove_RemoveEnd(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append(0)
	ll.Append(1)
	ll.Append(2)

	ll.Remove(2)
	assert.Equal(t, []int{0, 1}, ll.ToArray())
}

func TestRemove_NotFound(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append(0)
	ll.Append(1)
	ll.Append(2)

	err := ll.Remove(3)
	assert.EqualError(t, err, "value not found")
	assert.Equal(t, []int{0, 1, 2}, ll.ToArray())
}

func TestRemoveAt_Begining(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append(3)
	ll.Append(6)
	ll.Append(9)

	ll.RemoveAt(0)
	assert.Equal(t, []int{6, 9}, ll.ToArray())
}

func TestRemoveAt_Middle(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append(3)
	ll.Append(6)
	ll.Append(9)

	ll.RemoveAt(1)
	assert.Equal(t, []int{3, 9}, ll.ToArray())
}

func TestRemoveAt_End(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append(3)
	ll.Append(6)
	ll.Append(9)

	ll.RemoveAt(2)
	assert.Equal(t, []int{3, 6}, ll.ToArray())
}

func TestRemoveAt_OutOfBounds(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append(3)
	ll.Append(6)
	ll.Append(9)

	err := ll.RemoveAt(5)
	assert.EqualError(t, err, "out of bounds")
}

func TestReverse(t *testing.T) {
	ll := &linkedlist.LinkedList{}

	ll.Append(0)
	ll.Append(1)
	ll.Append(2)

	linkedlist.ReverseList(ll)

	assert.Equal(t, []int{2, 1, 0}, ll.ToArray())
}
