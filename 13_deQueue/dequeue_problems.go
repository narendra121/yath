package dequeue

import "fmt"

/*
inp:- []int{20, 40, 30, 10, 60}, 3

out:
40
40
60
*/
func MaximumsInallSubArraysOfSizeK(arr []int, k int) {
	dq := DeQueueArrImpl{}
	dq.Init(k)

	for i := 0; i < k; i++ {
		f := dq.GetRear().(int)
		for !dq.IsEmpty() && arr[i] >= arr[f] {
			dq.DeleteRear()
			f = dq.GetRear().(int)
		}
		dq.InsertRear(i)
	}

	for i := k; i < len(arr); i++ {
		fmt.Println(arr[dq.GetFront().(int)])

		for !dq.IsEmpty() && dq.GetFront().(int) <= i-k {
			dq.DeleteFront()
		}

		f := dq.GetRear().(int)
		for !dq.IsEmpty() && arr[i] >= arr[f] {
			dq.DeleteRear()
			f = dq.GetRear().(int)
		}

		dq.InsertRear(i)
	}
	fmt.Println(arr[dq.GetFront().(int)])
}

/*
inp:=[]int{4, 8, 7, 4}, []int{6, 5, 3, 5})

out:= 2

*/
func FirstCircularTourNaiveSolution(pertrol, distance []int) int {
	for start := 0; start < len(distance); start++ {
		currPetrol := 0
		end := start
		for {
			currPetrol += pertrol[end] - distance[end]
			if currPetrol < 0 {
				break
			}
			end = (end + 1) % len(distance)
			if start == end {
				return start + 1
			}
		}

	}
	return -1
}
func FirstCircularTourNaive(pertrol, distance []int) int {
	start, curr_petrol, prev_petrol := 0, 0, 0
	for i := 0; i < len(pertrol); i++ {
		curr_petrol += (pertrol[i] - distance[i])
		if curr_petrol < 0 {
			start = i + 1
			prev_petrol += curr_petrol
			curr_petrol = 0
		}
	}
	if (curr_petrol + prev_petrol) >= 0 {
		return start + 1
	}
	return -1
}
