package main

import "fmt"

func main() {
	nums := []int{
		5, 28, 16, 9, 37, 15,
	}
	target := 52
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	returnArr := []int{}
	m := make(map[int]int)
	for index := range nums {
		m[nums[index]] = index
	}

	for index := range nums {
		waitSelectNum := target - nums[index]
		isHaveNum, exist := m[waitSelectNum]
		if exist == true {
			fmt.Println(index, isHaveNum)
			returnArr = append(returnArr, index)
			returnArr = append(returnArr, isHaveNum)
			break
		}

	}
	return returnArr

}
