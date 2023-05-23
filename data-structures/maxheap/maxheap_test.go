package maxheap_test

import (
	"testing"

	"bensivo.com/leetcode/data-structures/maxheap"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	mh := maxheap.New()
	mh.Insert(0)
	mh.Insert(1)
	mh.Insert(2)
	mh.Insert(3)
	mh.Insert(4)
	mh.Insert(5)

	assert.Equal(t, []int{5, 3, 4, 0, 2, 1}, mh.Items)
}

func TestPop(t *testing.T) {
	mh := maxheap.New()
	mh.Insert(0)
	mh.Insert(1)
	mh.Insert(2)
	mh.Insert(3)
	mh.Insert(4)
	mh.Insert(5)

	v, _ := mh.Pop()
	assert.Equal(t, 5, v)

	v, _ = mh.Pop()
	assert.Equal(t, 4, v)

	v, _ = mh.Pop()
	assert.Equal(t, 3, v)
}
