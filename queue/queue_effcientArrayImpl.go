package queue

type efficientArrayQueue struct {
	Len   int
	Cap   int
	Arr   []int
	Front int
	Rear  int
}

var QueueEfficientArray Queue = &efficientArrayQueue{}

func (q *efficientArrayQueue) Init(len, cap int) {
	q.Arr = make([]int, cap)
	q.Len = len
	q.Cap = cap
	q.Front = 0
	q.Rear = 0
}
func (q *efficientArrayQueue) Size() int {
	return q.Len
}

func (q *efficientArrayQueue) IsFull() bool {
	return q.Len == q.Cap
}

func (q *efficientArrayQueue) IsEmpty() bool {
	return q.Len == 0
}
func (q *efficientArrayQueue) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	return (q.Front + q.Len - 1) % q.Cap
}

func (q *efficientArrayQueue) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Front
}
func (q *efficientArrayQueue) EnQueue(x int) {
	if q.IsFull() {
		return
	}
	q.Rear = (q.GetRear() + 1) % q.Cap
	q.Arr[q.Rear] = x
	q.Len++
}
func (q *efficientArrayQueue) DeQueue() {
	if q.IsEmpty() {
		return
	}
	q.Front = (q.Front + 1) % q.Cap
	q.Len--
}
