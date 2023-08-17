package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func find_max_elements(array []int32) int32 {
	n := array[0]
	target := array[1]

	dp := make([]int32, target+1)
	for i := int32(2); i <= n+1; i++ {
		num := array[i]
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+1)
		}
	}

	return dp[target]
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrayCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var array []int32

	// Read the array count and target
	arrayItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 32)
	checkError(err)
	arrayItem := int32(arrayItemTemp)
	array = append(array, arrayItem)

	for i := 0; i < int(arrayCount); i++ {
		arrayItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 32)
		checkError(err)
		arrayItem := int32(arrayItemTemp)
		array = append(array, arrayItem)
	}

	result := find_max_elements(array)

	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
