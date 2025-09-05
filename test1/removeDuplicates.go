package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	point := &nums
	s := removeDuplicates(point)
	fmt.Println("结果", s, nums)
}

func removeDuplicates(nums *[]int) int {
	fmt.Println("nums", nums)
	var result int
	var newArr []int
	checkMap := make(map[int]bool)
	for i := range *nums {
		_, exsit := checkMap[(*nums)[i]]
		if exsit {
			//
		} else {
			newArr = append(newArr, (*nums)[i])
			checkMap[(*nums)[i]] = true

		}

	}
	for i := range newArr {
		(*nums)[i] = newArr[i]
	}
	*nums = (*nums)[0:len(newArr)]
	fmt.Println("缩减后nums", nums)
	result = len(*nums)
	return result
}
