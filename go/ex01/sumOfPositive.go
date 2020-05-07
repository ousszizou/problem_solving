// Go episode N04
// challenge N08 => (Sum Of Positive) - 8Kyu

// By ALGORITHM

package main

import f "fmt"

func PositiveSum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			sum += numbers[i]
		}
	}
	return sum
}

func main() {

	numbers1 := []int{1,2,3,4,5,6,7,8,9,10}
	numbers2 := []int{-1,-2,-3,-4,-5}

	f.Println("sum = ",PositiveSum(numbers1))
	f.Println("sum = ",PositiveSum(numbers2))
}