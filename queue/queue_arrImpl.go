package queue

type arrayQueue struct {
	Len int
	Cap int
	Arr []int
}

var QueueArrImpl Queue = &arrayQueue{}

func (q *arrayQueue) Init(len, cap int) {
	q.Arr = make([]int, cap)
	q.Len = 0
	q.Cap = cap
}

func (q *arrayQueue) Size() int {
	return q.Len
}

func (q *arrayQueue) IsFull() bool {
	return q.Len == q.Cap
}

func (q *arrayQueue) IsEmpty() bool {
	return q.Len == 0
}

func (q *arrayQueue) EnQueue(x int) {
	if q.IsFull() {
		return
	}
	q.Arr[q.Len] = x
	q.Len++
}
func (q *arrayQueue) DeQueue() {
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

func (q *arrayQueue) GetFront() interface{} {
	if q.IsEmpty() {
		return -1
	}
	return 0
}

func (q *arrayQueue) GetRear() interface{} {
	if q.IsEmpty() {
		return -1
	}
	return q.Len
}
func (q *arrayQueue) GetQueue() interface{} {
	return q.Arr
}
