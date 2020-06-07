package core

import "strconv"

type Answer struct {
	X1        int
	X2        int
	Y         int
	Str       string
	RandomMin int
	RandomMax int
}
type ListAnswer struct {
	AList      []Answer
	Generation int
}

func (la ListAnswer) PrintListAnswer() {
	println("**************************************")
	println("Generation: " + strconv.Itoa(la.Generation))
	for _, a := range la.AList {
		println("X1=" + strconv.Itoa(a.X1) + ",X2=" + strconv.Itoa(a.X2) + ",Value=" + strconv.Itoa(a.Y))
		println("Binary string is " + a.Str)
		println("------------------------------------")
	}
	println("**************************************")
}
