package utils

import (
	"math"
	"strconv"
)

func ConvertToBi(a int, b int, length int) string {
	bia := ConvertToBinary(a, length)
	bib := ConvertToBinary(b, length)
	return bia + bib
}

func ConvertToBinary(n int, length int) string {
	var b string
	switch {
	case n == 0:
		for i := 0; i < length; i++ {
			b += "0"
		}
	case n > 0:
		for ; n > 0; n /= 2 {
			b = strconv.Itoa(n%2) + b
		}
		j := length - len(b)
		for i := 0; i < j; i++ {
			b = "0" + b
		}
	}
	return b
}

func Str2DEC(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}

func GetIndexBinaryNot(binary string, index int) string {
	length := len(binary)
	if index <= 0 || index > length {
		return ""
	}
	biru := []rune(binary)
	snum := Str2DEC(binary)
	differenceNum := 0
	if length != index {
		differenceNum = int(math.Pow(2, float64(length-index)))
	} else {
		differenceNum = 1
	}
	if string(biru[index-1]) == "0" {
		snum = snum + differenceNum
	} else if string(biru[index-1]) == "1" {
		snum = snum - differenceNum
	} else {
		return ""
	}
	newBi := ConvertToBinary(snum, length)
	return newBi
}
