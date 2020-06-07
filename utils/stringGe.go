package utils

import "fmt"

func SwapString(str1 string, str2 string, index int) (string, string) {
	if index >= len(str1) || index >= len(str2) || index < 0 {
		return str1, str2
	}
	tmp1 := str1[0:index] + str2[index:]
	tmp2 := str2[0:index] + str1[index:]
	fmt.Printf("%s to %s index is %d ", str1, tmp1, index)
	fmt.Printf("%s to %s \n", str2, tmp2)
	return tmp1, tmp2
}
