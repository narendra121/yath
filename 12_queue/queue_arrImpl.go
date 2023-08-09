package queue

type QueueArrImpl struct {
	Len int
	Cap int
	Arr []interface{}
}

func (q *QueueArrImpl) Init(len, cap int) {
	q.Arr = make([]interface{}, cap)
	q.Len = 0
	q.Cap = cap
}

func (q *QueueArrImpl) Size() int {
	return q.Len
}

func (q *QueueArrImpl) IsFull() bool {
	return q.Len == q.Cap
}

func (q *QueueArrImpl) IsEmpty() bool {
	return q.Len == 0
}

func (q *QueueArrImpl) EnQueue(x interface{}) {
	if q.IsFull() {
		return
	}
	q.Arr[q.Len] = x
	q.Len++
}
func (q *QueueArrImpl) DeQueue() {
	if q.IsEmpty() {
		return
	}
	for i := 0; i < q.Len-1; i++ {
		q.Arr[i] = q.Arr[i+1]
	}
	for i := q.Len - 1; i < q.Cap; i++ {
		q.Arr[i] = 0
	}
	q.Len--
}

func (q *QueueArrImpl) GetFront() interface{} {
	if q.IsEmpty() {
		return -1
	}
	return 0
}

func (q *QueueArrImpl) GetRear() interface{} {
	if q.IsEmpty() {
		return -1
	}
	return q.Len
}
func (q *QueueArrImpl) GetQueue() interface{} {
	return q.Arr
}
