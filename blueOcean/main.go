package main

import "fmt"

func BlueOcean(first []string, second []string) []string {
	result := []string{}

	for i := 0; i < len(first); i++ {
		found := false

		for j := 0; j < len(second); j++ {

			if first[i] == second[j] {
				found = true
			}

		}

		if found == false {
			result = append(result, first[i])
		}

	}

	return result
}

func main() {
	fmt.Println(BlueOcean([]string{"red", "blue", "yellow", "black", "grey"}, []string{"blue"}))
	fmt.Println(BlueOcean([]string{"red", "blue", "yellow", "black", "grey"}, []string{"blue", "yellow"}))
	fmt.Println(BlueOcean([]string{"red", "blue", "yellow", "black", "grey"}, []string{"blue", "yellow", "black"}))
}
