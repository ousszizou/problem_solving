// Go episode N06
// challenge N12 => (Sum of Digits / Digital Root) - 6Kyu

// By ALGORITHM ACADEMY

package main

import (
	"fmt"
	"strconv"
)

// DigitalRoot function
func DigitalRoot(n int) int {
	str := strconv.Itoa(n)
	var result int = 0
	for _,d := range str {
		currentNumber,_ := strconv.Atoi(string(d))
		result += currentNumber
	}
	
	if result < 10 {
		return result
	}
	return DigitalRoot(result)
}

// Another method

// DigitalRoot function
// func DigitalRoot(n int) int {
//   return (n - 1) % 9 + 1
// }

func main()  {
	fmt.Println(DigitalRoot(16))
	fmt.Println(DigitalRoot(942))
	fmt.Println(DigitalRoot(132189))
}
