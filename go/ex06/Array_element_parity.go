// Go episode N08
// challenge N14 = > (Array element parity) - 7Kyu

// By ALGORITHM ACADEMY

package main

import "fmt"

func count(collection []int, num int) int {
	var counter int
	for i := 0; i < len(collection); i++ {
		if (num == collection[i]) {
			counter++
		}
	}
	return counter
}

func solve(arr []int) int {
	var result int;
	for i := 0; i < len(arr); i++ {
		if (count(arr, -arr[i]) > 0) {
			continue
		} else {
			result = arr[i]
		}
	}
	return result
}

func main() {
	arr1 := []int{1, -1, 2, -2, 3}
	arr2 := []int{-3, 1, 2, 3, -1, -4, -2}
	fmt.Println(solve(arr1))
	fmt.Println(solve(arr2))
}
