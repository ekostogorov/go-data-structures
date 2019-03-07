package queue

// Queue represents queue struct
type Queue struct {
	size    int
	storage []string
}

// New constructs queue
func New() *Queue {
	return &Queue{}
}

// Enqueue adds element to end of queue
func (q *Queue) Enqueue(element string) {
	q.storage = append(q.storage, element)
	q.size++
}

// Dequeue removes element from start of queue
func (q *Queue) Dequeue() string {
	element := q.storage[0]
	q.storage = q.storage[1:]

	return element
}
