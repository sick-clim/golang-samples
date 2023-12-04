package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func bogo(numbers []int) []int {
	// deprecated
	//rand.Seed(time.Now().UnixNano())
	for {
		fmt.Print(".")
		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
		if sort.IntsAreSorted(numbers) {
			break
		}
	}

	return numbers
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	bogo(numbers)
	fmt.Println(numbers)
}
