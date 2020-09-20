// Go episode N12
// challenge N17 = > (String ends with?) - 7Kyu

// By ALGORITHM ACADEMY

package main

import (
	"fmt"
	"strings"
)

func solution1(str, ending string) bool {
  return strings.HasSuffix(str, ending)
}

func solution2(str, ending string) bool {
  return len(str) >= len(ending) && str[len(str) - len(ending):] == ending
}

// str = abc     ending = c
// len(str) = 3  len(ending) = 1

// len(str) - len(ending) = 2
//  str[2:] => c


// str = abc     ending = bc
// len(str) = 3  len(ending) = 2

// len(str) - len(ending) = 1
// str[1:] => bc


func main() {
	fmt.Println(solution2("abc", "c"))          // true
	fmt.Println(solution2(" ", ""))             // true
	fmt.Println(solution2("    ", "    "))      // true
	fmt.Println(solution2("", "t"))             // false
	fmt.Println(solution2("$a = $b + 1", "+1")) // false
}
