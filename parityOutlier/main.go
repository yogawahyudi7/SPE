package main

import "fmt"

func ParityOutlier(number []int) interface{} {
	var result interface{}

	odd := []int{}
	even := []int{}

	if len(number) < 3 {
		result = "the minimum length of the array is 3"
		return result
	}

	for i := 0; i < len(number); i++ {
		if number[i]%2 == 0 {
			even = append(even, number[i])
		} else {
			odd = append(odd, number[i])
		}
	}

	if len(odd) > len(even) {
		result = even
	} else if len(odd) < len(even) {
		result = odd
	} else {
		result = "no outlier numbers"
	}

	// fmt.Println("ODD :", odd)
	// fmt.Println("EVEN :", even)

	return result
}

func main() {
	fmt.Println(ParityOutlier([]int{1, 2, 3}))
	fmt.Println(ParityOutlier([]int{1, 2, 3, 4, 5}))
	fmt.Println(ParityOutlier([]int{1, 2, 3, 4, 5, 6, 8}))
	fmt.Println(ParityOutlier([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(ParityOutlier([]int{1, 2}))
}
