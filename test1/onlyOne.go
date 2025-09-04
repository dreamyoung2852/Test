package main

import "fmt"

func main() {
	nums := []int{
		5, 3, 6, 99, 6, 3, 5,
	}
	fmt.Println(onlyOne(nums))
}

func onlyOne(nums []int) int {
	result := 0
	checkMap := make(map[int]int)
	for index := range nums {
		_, exist := checkMap[nums[index]]
		if exist == true {
			checkMap[nums[index]] += 1
		} else {
			checkMap[nums[index]] = 1
		}

	}
	for k, v := range checkMap {
		if v == 1 {
			result = k
		}
	}

	return result

}
