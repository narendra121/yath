package matrix

import (
	"fmt"
)

/*
inp   -- [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})
out -- 1       2       3       4
8       7       6       5
9       10      11      12
16      15      14      13
*/
func PrintSnakePattern(inp [][]int) {
	for i := 0; i < len(inp); i++ {
		if i%2 == 0 {
			for j := 0; j < len(inp[0]); j++ {
				fmt.Print(inp[i][j], "\t")
			}
			fmt.Println()
		} else {
			for j := len(inp[0]) - 1; j >= 0; j-- {
				fmt.Print(inp[i][j], "\t")
			}
			fmt.Println()
		}
	}
}

/*
in [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
out        1       2       3       4
8       12      16
15      14      13
5       9
*/
func PrintMatrixBoundary(arr [][]int) {
	if len(arr) == 1 {
		for i := 0; i < len(arr[0]); i++ {
			fmt.Print(arr[0][i], "\t")
		}
		fmt.Println()
	} else if len(arr[0]) == 1 {
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i][0])
		}
		fmt.Println()
	} else {
		for i := 0; i < len(arr[0]); i++ {
			fmt.Print(arr[0][i], "\t")
		}
		fmt.Println()
		for i := 1; i < len(arr); i++ {
			fmt.Print(arr[i][len(arr[0])-1], "\t")
		}
		fmt.Println()
		for i := len(arr[0]) - 2; i >= 0; i-- {
			fmt.Print(arr[len(arr)-1][i], "\t")
		}
		fmt.Println()
		for i := 1; i < len(arr)-1; i++ {
			fmt.Print(arr[i][0], "\t")
		}
		fmt.Println()
	}
}

/*
[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

[[1 5 9 13] [2 6 10 14] [3 7 11 15] [4 8 12 16]]

*/

func TransposeOfMatrix(arr [][]int) [][]int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
	return arr
}

/*
inp [[1 2 3 4] [5 6 7 8] [9 10 11 12] [13 14 15 16]]
out [[4 8 12 16] [3 7 11 15] [2 6 10 14] [1 5 9 13]]
*/
func RotateMatrixAntiClockBy90(arr [][]int) [][]int {
	TransposeOfMatrix(arr)
	for i := 0; i < len(arr); i++ {
		low, high := 0, len(arr)-1
		for low < high {
			arr[low][i], arr[high][i] = arr[high][i], arr[low][i]
			low++
			high--
		}
	}
	return arr
}

/*
in [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

out
1       2       3       4
8       12      16
15      14      13
9       5
6       7
11
10
*/
func PrintSpiral(arr [][]int) {
	top, right, bottom, left := 0, len(arr[0])-1, len(arr)-1, 0
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			fmt.Print(arr[top][i], "\t")
		}
		top++
		fmt.Println()
		for i := top; i <= bottom; i++ {
			fmt.Print(arr[i][right], "\t")

		}
		right--
		fmt.Println()

		if top <= bottom {
			for i := right; i >= left; i-- {
				fmt.Print(arr[bottom][i], "\t")

			}
			bottom--
			fmt.Println()

		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				fmt.Print(arr[i][left], "\t")

			}
			left++
			fmt.Println()

		}
	}
}

/*
in [][]int{{10, 20, 30, 40}, {15, 25, 35, 45}, {27, 29, 37, 48}, {32, 33, 39, 50}}, 39))
out 3 ,2
*/

func SearchInRowWiseAndCloumnWiseSortedMatrix(arr [][]int, searchItem int) string {
	i, j := 0, len(arr)-1
	for i < len(arr) && j >= 0 {
		if arr[i][j] == searchItem {
			return fmt.Sprint(i, j)
		} else if arr[i][j] > searchItem {
			j--
		} else {
			i++
		}
	}
	return "Not Found"
}
