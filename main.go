package main

import (
	"geneticAlgothm/core"
	"geneticAlgothm/funcbase"
	"geneticAlgothm/utils"
	"strconv"
)

// get the max value of func y=x1^3+x2^2-5*x1*x2
// x1 and x2 ranges in [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]

func main() {
	var firstAnswers core.ListAnswer
	var survivalAnswers core.ListAnswer
	//the first generation
	for i := 0; i < 6; i++ {
		an := core.Answer{}
		an.X1 = utils.RandomNum(16)
		an.X2 = utils.RandomNum(16)
		an.Str = utils.ConvertToBi(an.X1, an.X2, 4)
		an.Y = funcbase.GetFuncValue(an.X1, an.X2)
		firstAnswers.AList = append(firstAnswers.AList, an)
		firstAnswers.Generation = 1
	}
	firstAnswers.PrintListAnswer()
	//The survival of the fittest
	//The max value can survive,others are generated randomly,pick 4
	survivalAnswers = GetSursivalAnswers(firstAnswers)
	survivalAnswers.PrintListAnswer()
	//cross ope
	println("cross")
	crossLa := CrossoverOperation(survivalAnswers)
	crossLa.PrintListAnswer()
	variaLa := GeneticVariation(crossLa)
	newLa := Combine(variaLa)
	newLa.PrintListAnswer()
	println("circulation start !!!!!!")
	for i := 0; i < 100; i++ {
		newLa = GetSursivalAnswers(newLa)
		newLa = CrossoverOperation(newLa)
		newLa = GeneticVariation(newLa)
		newLa = Combine(newLa)
		newLa.PrintListAnswer()
	}
	//bnum := utils.GetIndexBinaryNot("10101001",4)
	//println(bnum)

}

func GetSursivalAnswers(la core.ListAnswer) core.ListAnswer {
	max := 0
	maxIndex := 0
	sum := 0
	percent := 0
	var survivalAnswers core.ListAnswer
	survivalAnswers.Generation = la.Generation
	for i, a := range la.AList {
		if max < a.Y {
			max = a.Y
			maxIndex = i
		}
		sum += a.Y
	}
	println("The max value is " + strconv.Itoa(max))
	survivalAnswers.AList = append(survivalAnswers.AList, la.AList[maxIndex])
	sum = sum - la.AList[maxIndex].Y
	la.AList = append(la.AList[0:maxIndex], la.AList[maxIndex+1:]...)
	for i, a := range la.AList {
		if i == len(la.AList)-1 {
			la.AList[i].RandomMin = percent
			la.AList[i].RandomMax = 99
			break
		}
		p := a.Y * 100 / sum
		la.AList[i].RandomMin = percent
		la.AList[i].RandomMax = percent + p
		percent = percent + p
	}
	breakFlag := 0
	for {
		rand := utils.RandomNum(100)
		//println("get random num : " + strconv.Itoa(rand))
		for i := 0; i < len(la.AList); i++ {
			if la.AList[i].RandomMin <= rand && rand < la.AList[i].RandomMax {
				//pick it
				survivalAnswers.AList = append(survivalAnswers.AList, la.AList[i])
				//origin list remove it
				la.AList = append(la.AList[0:i], la.AList[i+1:]...)
				breakFlag++
				break
			}
		}
		if breakFlag == 3 {
			break
		}
	}
	return survivalAnswers

}

func CrossoverOperation(la core.ListAnswer) core.ListAnswer {
	var firstList []core.Answer
	var secondList []core.Answer
	var newLa core.ListAnswer
	//random matching
	firstList = append(firstList, la.AList[len(la.AList)-1])
	matchNum := utils.RandomNum(3)
	switch matchNum {
	case 0:
		firstList = append(firstList, la.AList[matchNum])
		secondList = append(secondList, la.AList[1], la.AList[2])
	case 1:
		firstList = append(firstList, la.AList[matchNum])
		secondList = append(secondList, la.AList[0], la.AList[2])
	case 2:
		firstList = append(firstList, la.AList[matchNum])
		secondList = append(secondList, la.AList[0], la.AList[1])

	}
	//set the intersection position at random
	randomPosi := utils.RandomNum(8)
	fa, fb := utils.SwapString(firstList[0].Str, firstList[1].Str, randomPosi)
	//to have 4 children
	randomPosi_2 := utils.RandomNumExcept(8, randomPosi)
	fc, fd := utils.SwapString(firstList[0].Str, firstList[1].Str, randomPosi_2)

	randomPosi2 := utils.RandomNum(8)
	sa, sb := utils.SwapString(secondList[0].Str, secondList[1].Str, randomPosi2)
	//to have 4 children
	randomPosi2_2 := utils.RandomNumExcept(8, randomPosi)
	sc, sd := utils.SwapString(secondList[0].Str, secondList[1].Str, randomPosi2_2)

	newLa.AList = append(newLa.AList, core.Answer{Str: fa}, core.Answer{Str: fb}, core.Answer{Str: fc},
		core.Answer{Str: fd}, core.Answer{Str: sa}, core.Answer{Str: sb}, core.Answer{Str: sc}, core.Answer{Str: sd})
	newLa.Generation = la.Generation + 1
	return newLa
}

func GeneticVariation(la core.ListAnswer) core.ListAnswer {
	variaAnswer := utils.RandomNum(8)
	variaIndex := utils.RandomNum(8)
	println("before variation : " + la.AList[variaAnswer].Str)
	la.AList[variaAnswer].Str = utils.GetIndexBinaryNot(la.AList[variaAnswer].Str, variaIndex+1)
	println("after variation : " + la.AList[variaAnswer].Str)
	return la
}

func Combine(la core.ListAnswer) core.ListAnswer {
	for i, a := range la.AList {
		str1 := a.Str[0:4]
		str2 := a.Str[4:]
		la.AList[i].X1 = utils.Str2DEC(str1)
		la.AList[i].X2 = utils.Str2DEC(str2)
		la.AList[i].Y = funcbase.GetFuncValue(la.AList[i].X1, la.AList[i].X2)
	}
	return la
}
