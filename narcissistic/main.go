package main

import (
	"fmt"
	"math"
	"strconv"
)

func Narcissistic(number int) bool {
	result := false

	strNumber := strconv.Itoa(number)
	lengtNumber := len(strNumber)

	sum := 0
	for i := 0; i < lengtNumber; i++ {
		numberI, _ := strconv.Atoi(string(strNumber[i]))
		numb := math.Pow(float64(numberI), float64(lengtNumber))

		sum = sum + int(numb)
	}

	// fmt.Println("SUM : ", sum)
	if number == sum {
		result = true
	}

	return result
}

func main() {
	// t := "roda"
	// fmt.Println(string(t[0]))

	fmt.Println(Narcissistic(1634))
	fmt.Println(Narcissistic(153))
	fmt.Println(Narcissistic(111))
}
