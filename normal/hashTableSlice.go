package normal

import "fmt"

type HashTableSlice interface {
	AddValueToSlice(inp ...int)
	AddValueIndex(val int, index int)
	IsValueExist(val int) bool
	GetIndexOfValue(val int) int
	PrintAll()
}

type HashTable struct {
	mySlice []int
	myMap   map[int]int
}

func NewHashTable() HashTableSlice {
	return &HashTable{
		mySlice: make([]int, 0),
		myMap:   make(map[int]int),
	}
}

func (mt *HashTable) AddValueToSlice(inp ...int) {
	for _, val := range inp {
		temp := len(mt.mySlice)
		if mt.IsValueExist(val) {
			continue
		}
		mt.mySlice = append(mt.mySlice, val)
		if temp == 0 {
			temp = 0
		} else {
			temp -= 1
		}
		mt.AddValueIndex(val, temp)
	}
}

func (mt *HashTable) AddValueIndex(val int, index int) {
	mt.myMap[val] = index
}

func (mt *HashTable) IsValueExist(val int) bool {
	_, ok := mt.myMap[val]
	return ok
}

func (mt *HashTable) GetIndexOfValue(val int) int {
	idx, ok := mt.myMap[val]
	if !ok {
		return -1
	}
	return idx
}

func (mt *HashTable) PrintAll() {
	for _, val := range mt.mySlice {
		fmt.Println(val)
	}
}
