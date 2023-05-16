package main

import "fmt"

func FindNeedle(haystack []string, needle string) int {
	result := 0

	for i, v := range haystack {
		if v == needle {
			result = i

			return i //burn the haystack
		}
	}

	return result
}

func main() {
	fmt.Println(FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue"))
}
