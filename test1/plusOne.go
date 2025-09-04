package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{
		9, 6, 5,
	}

	fmt.Println(plusOne(nums))
}

// 1.排序
// 2.进位及拆分
func plusOne(digits []int) []int {
	var result []int
	sort.Ints(digits)
	var intStr string
	for i := range digits {
		intStr += strconv.Itoa(digits[i])
	}
	resultNum, err := strconv.Atoi(intStr)
	if err != nil {
		panic(err)
	}
	resultNum += 1
	for resultNum > 0 {
		result = append(result, resultNum%10)
		resultNum = resultNum / 10
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	fmt.Println(digits, intStr, resultNum)
	return result
}
