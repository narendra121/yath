package stacks

/*
Init  -- Initialize
Push  -- Add element or store
Pop  -- Delete the Data
Peek -- See the top element
*/
type Stack interface {
	Init()
	Push(i interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
	IsEmpty() bool
}
