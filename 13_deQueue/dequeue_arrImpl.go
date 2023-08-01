package dequeue

/*
InsertFront:- front=(front-1+cap)%cap
DeleteFront:- front=(front+1)%cap
InsertRear:- rear=(rear+1)%cap
DeleteRear:- rear=(rear-1+cap)%cap
*/

type deQueue struct {
	arr   []int
	front int
	cap   int
	size  int
}

var DeQueueArrImpl DoublyEndedQueue = &deQueue{}

func (d *deQueue) Init(cap int) {
	d.arr = make([]int, cap)
	d.cap = cap
	d.front = 0
	d.size = 0
}

func (d *deQueue) Size() int {
	return d.size
}

func (d *deQueue) IsFull() bool {
	return d.size == d.cap
}

func (d *deQueue) IsEmpty() bool {
	return d.size == 0
}

func (d *deQueue) GetFront() interface{} {
	if d.IsEmpty() {
		return -1
	}
	return d.arr[d.front]
}

func (d *deQueue) GetRear() interface{} {
	if d.IsEmpty() {
		return -1
	}
	rear := (d.front + d.size - 1) % d.cap
	return d.arr[rear]
}

func (d *deQueue) InsertFront(x interface{}) {
	if d.IsFull() {
		return
	}
	d.front = (d.front + d.cap - 1) % d.cap
	d.arr[d.front] = x.(int)
	d.size++
}

func (d *deQueue) InsertRear(x interface{}) {
	if d.IsFull() {
		return
	}
	rear := (d.front + d.size) % d.cap
	d.arr[rear] = x.(int)
	d.size++
}

func (d *deQueue) DeleteFront() {
	if d.IsEmpty() {
		return
	}
	d.arr[d.front] = 0
	d.front = (d.front + 1) % d.cap
	d.size--
}

func (d *deQueue) DeleteRear() {
	if d.IsEmpty() {
		return
	}
	rear := (d.front + d.size - 1) % d.cap
	d.arr[rear] = 0
	d.size--
}

func (d *deQueue) GetArray() interface{} {
	return d.arr
}
