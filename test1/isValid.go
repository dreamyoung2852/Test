package main

import (
	"fmt"
)

func main() {
	waitValid := "([])}"
	fmt.Println(isValid(waitValid))
}

// 1.校验是否成对出现,并且包含括号
// 2.校验括号位置,也就是不能 ")("
// 3.校验包裹正确性
func isValid(s string) bool {
	var result bool
	checkMap := make(map[string]string)
	checkMap1 := make(map[string]bool)
	checkMap["("] = ")"
	checkMap["["] = "]"
	checkMap["{"] = "}"
	checkMap[")"] = "("
	checkMap["]"] = "["
	checkMap["}"] = "{"
	checkMap1["("] = true
	checkMap1["["] = true
	checkMap1["{"] = true

	charSlice := []rune(s)
	fmt.Println(charSlice)
	charMap := make(map[string]int)
	for i, c := range charSlice {
		charMap[string(c)] = i
		fmt.Println("charMap", charMap)
	}

	for _, c := range charSlice {
		v1, exsit1 := checkMap[string(c)]
		//1
		if exsit1 == true {
			v2, exsit2 := charMap[v1]
			fmt.Println("exsit2", v2)
			if exsit2 == true {
				rightStr := ""
				leftStr := ""
				if checkMap1[string(c)] {
					rightStr = string(c)
					leftStr = v1
				} else {
					rightStr = v1
					leftStr = string(c)
				}
				fmt.Println("charMap[rightStr]", charMap[rightStr], charMap[leftStr])
				//2
				if charMap[rightStr] < charMap[leftStr] {
					//3
					num := charMap[leftStr] - charMap[rightStr]
					if num%2 != 0 {

					} else {
						result = false
						break
					}

				} else {
					result = false
					break
				}

				fmt.Println("rightStr", rightStr)

				result = true
			} else {
				result = false
				break
			}
		} else {

			result = false
			break
		}
		fmt.Println(string(c))
	}
	return result

}
