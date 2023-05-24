package queue

import (
	"errors"

	"bensivo.com/leetcode/data-structures/linkedlist"
)

var ErrQueueEmpty error = errors.New("queue is empty")

type Queue struct {
	ll *linkedlist.LinkedList
}

func New() *Queue {
	return &Queue{
		ll: linkedlist.New(),
	}
}

func (q *Queue) Push(value string) {
	q.ll.Append(value)
}

func (q *Queue) Pop() (string, error) {
	if q.ll.Head == nil {
		return "", ErrQueueEmpty
	}

	value := q.ll.Head.Value
	q.ll.RemoveAt(0)

	return value, nil
}

func (q *Queue) ToArray() []string {
	return q.ll.ToArray()
}
