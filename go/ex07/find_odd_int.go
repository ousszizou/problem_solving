// Go episode N10
// challenge N15 = > (Find the odd int) - 6Kyu

// By ALGORITHM ACADEMY

package main

import "fmt"

func findOdd(seq []int) int {

	var num int
	for i := 0; i < len(seq); i++ {
		num = num ^ seq[i]
	}
	return num
}

// XOR
// 0 ^ 0 = 0 
// 1 ^ 1 = 0
// 0 ^ 1 = 1
// 1 ^ 0 = 1

func main() {
	arr := []int{1, 1, 1} // 0 ^ 1 = 1
	arr1 := []int{1, 1, 2, 5, 3, 2, 5} // 0 ^ 3 = 3
	fmt.Println(findOdd(arr))
	fmt.Println(findOdd(arr1))
}
