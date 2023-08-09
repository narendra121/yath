package queue

type QueueEfficientArray struct {
	Len   int
	Cap   int
	Arr   []interface{}
	Front int
	Rear  int
}

func (q *QueueEfficientArray) Init(len, cap int) {
	q.Arr = make([]interface{}, cap)
	q.Len = len
	q.Cap = cap
	q.Front = 0
	q.Rear = 0
}
func (q *QueueEfficientArray) Size() int {
	return q.Len
}

func (q *QueueEfficientArray) IsFull() bool {
	return q.Len == q.Cap
}

func (q *QueueEfficientArray) IsEmpty() bool {
	return q.Len == 0
}
func (q *QueueEfficientArray) GetRear() interface{} {
	if q.IsEmpty() {
		return -1
	}
	return (q.Front + q.Len - 1) % q.Cap
}

func (q *QueueEfficientArray) GetFront() interface{} {
	if q.IsEmpty() {
		return -1
	}
	return q.Front
}
func (q *QueueEfficientArray) EnQueue(x interface{}) {
	if q.IsFull() {
		return
	}
	q.Rear = (q.GetRear().(int) + 1) % q.Cap
	q.Arr[q.Rear] = x
	q.Len++
}
func (q *QueueEfficientArray) DeQueue() {
	if q.IsEmpty() {
		return
	}
	q.Front = (q.Front + 1) % q.Cap
	q.Len--
}
func (q *QueueEfficientArray) GetQueue() interface{} {
	return q.Arr
}
