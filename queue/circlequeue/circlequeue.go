package circlequeue

type ArrayCrycleQueue struct {
	data   []interface{}
	head   int
	tail   int
	length int
}

func NewQueue(size int) *ArrayCrycleQueue {
	if size <= 0 {
		return nil
	}

	q := &ArrayCrycleQueue{
		data: make([]interface{}, size+1),
		head: 0,
		tail: 0,
	}
	return q
}

func (q *ArrayCrycleQueue) EnQueue(v interface{}) {
	if (q.tail+1)%len(q.data) == q.head {
		return
	}

	q.data[q.tail] = v
	q.tail = (q.tail + 1) % len(q.data)

	q.length++
}

func (q *ArrayCrycleQueue) DeQueue() interface{} {
	if q.tail == q.head {
		return nil
	}

	v := q.data[q.head]
	q.head = (q.head + 1 + len(q.data)) % len(q.data)

	q.length--
	return v
}

func (q *ArrayCrycleQueue) Length() int {
	return q.length
}
