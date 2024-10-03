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

// Deque removes and returns an element from the start of the queue
// or an error if empty
func (q *Queue[T]) Deque() (*T, error) {
	if q.Size() > 0 {
		item := q.collection[0]
		q.collection = q.collection[1:]
		return &item, nil
	}
	return nil, ErrQueueEmpty
}

// Enqueue adds an element at the end of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.collection = append(q.collection, item)
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
