package dequeue

/*
InsertFront:- front=(front-1+cap)%cap
DeleteFront:- front=(front+1)%cap
InsertRear:- rear=(rear+1)%cap
DeleteRear:- rear=(rear-1+cap)%cap
*/

type DeQueueArrImpl struct {
	arr   []interface{}
	front int
	cap   int
	size  int
}

func (d *DeQueueArrImpl) Init(cap int) {
	d.arr = make([]interface{}, cap)
	d.cap = cap
	d.front = 0
	d.size = 0
}

func (d *DeQueueArrImpl) Size() int {
	return d.size
}

func (d *DeQueueArrImpl) IsFull() bool {
	return d.size == d.cap
}

func (d *DeQueueArrImpl) IsEmpty() bool {
	return d.size == 0
}

func (d *DeQueueArrImpl) GetFront() interface{} {
	if d.IsEmpty() {
		return -1
	}
	return d.arr[d.front]
}

func (d *DeQueueArrImpl) GetRear() interface{} {
	if d.IsEmpty() {
		return -1
	}
	rear := (d.front + d.size - 1) % d.cap
	return d.arr[rear]
}

func (d *DeQueueArrImpl) InsertFront(x interface{}) {
	if d.IsFull() {
		return
	}
	d.front = (d.front + d.cap - 1) % d.cap
	d.arr[d.front] = x
	d.size++
}

func (d *DeQueueArrImpl) InsertRear(x interface{}) {
	if d.IsFull() {
		return
	}
	rear := (d.front + d.size) % d.cap
	d.arr[rear] = x
	d.size++
}

func (d *DeQueueArrImpl) DeleteFront() {
	if d.IsEmpty() {
		return
	}
	d.arr[d.front] = nil
	d.front = (d.front + 1) % d.cap
	d.size--
}

func (d *DeQueueArrImpl) DeleteRear() {
	if d.IsEmpty() {
		return
	}
	rear := (d.front + d.size - 1) % d.cap
	d.arr[rear] = nil
	d.size--
}

func (d *DeQueueArrImpl) GetArray() interface{} {
	return d.arr
}
