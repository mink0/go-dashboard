/*
	Simple in-memory FIFO queue with upper size limit based on doubly linked lists from container/list.
	This will balance memory usage and garbage collection for a long running jobs with long queues.
*/

package queue

import (
	"container/list"
)

type Queue struct {
	Elements *list.List
	Size     int
}

func New(size int) *Queue {
	return &Queue{
		Size:     size,
		Elements: list.New(),
	}
}

func (q *Queue) Push(element interface{}) {
	if q.Elements.Len() == q.Size {
		q.Pop()
	}
	q.Elements.PushBack(element)
}

func (q *Queue) Pop() (element *list.Element) {
	if q.Elements.Len() == 0 {
		return
	}

	element = q.Elements.Front()
	q.Elements.Remove(element)
	return
}

func (q *Queue) Get(max_items int) (elements []interface{}) {
	for e, i := q.Elements.Back(), 0; e != nil && i < max_items; e, i = e.Prev(), i+1 {
		// elements = preppend(elements, e.Value)
		elements = append(elements, nil)
		copy(elements[1:], elements)
		elements[0] = e.Value
	}

	return
}
