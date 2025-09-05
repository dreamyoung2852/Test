package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{
		"flower", "flow", "flight",
	}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {

	str1 := strs[0]

	for {
		var exist bool
		for _, v := range strs {
			exist = strings.HasPrefix(v, str1)
			if !exist {
				str1 = str1[0 : len(str1)-1]
			}
		}
		if exist == true || str1 == "" {
			break
		}

	}
	return str1

}
