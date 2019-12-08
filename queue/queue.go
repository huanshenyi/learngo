package queue

// An FIFO queue.
type Queue []int

// Pushes the element into the queue
// e.g.  q.Push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// return
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
