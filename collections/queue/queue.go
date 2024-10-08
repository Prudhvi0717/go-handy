package queue

import "errors"

// errors
var (
	ErrQueueEmpty = errors.New("queue is empty")
)

// Queue defines the queue collection
type Queue[T any] struct {
	collection []T
}

// Deque removes and returns the top element of the queue
// or an error if empty
func (q *Queue[T]) Deque() (item T, err error) {
	if q.Size() > 0 {
		item = q.collection[0]
		q.collection = q.collection[1:]
		return item, nil
	}
	return item, ErrQueueEmpty
}

// Enqueue adds an element at the end of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.collection = append(q.collection, item)
}

// Peek returns the top element of the queue or an error if empty
func (q *Queue[T]) Peek() (item T, err error) {
	if q.Size() > 0 {
		return q.collection[0], nil
	}
	return item, ErrQueueEmpty
}

// Size returns the size of the queue
func (q *Queue[T]) Size() int {
	return len(q.collection)
}

// New constructs a new queue
func New[T any]() *Queue[T] {
	return &Queue[T]{
		collection: []T{},
	}
}

// NewFromArray contrcuts a new queue from given collection of elements
func NewFromArray[T any](collection []T) *Queue[T] {
	return &Queue[T]{
		collection: collection,
	}
}
