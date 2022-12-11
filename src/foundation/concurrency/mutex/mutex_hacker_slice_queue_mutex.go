package mutex

import "sync"

type SliceQueue struct {
	data []any
	mu   sync.Mutex
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]any, 0, n)}
}

func (q *SliceQueue) Enqueue(v any) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

func (q *SliceQueue) Dequeue() any {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}

	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}
