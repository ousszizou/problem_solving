// Go episode N04
// challenge N09 => (Total Amout Of Points) - 8Kyu

// By ALGORITHM

package main

import (
	f "fmt"
	"strings"
)

func points(games []string) int {
	total := 0
	for i := 0; i < len(games); i++ {
		x := strings.Split(games[i], ":")[0]
		y := strings.Split(games[i], ":")[1]

		if x > y {
			total += 3
		} else if x < y {
			total += 0
		} else {
			total++
		}
	}
	return total
}

func main() {

	matches1 := []string{"1:0","2:0","3:0","4:0","2:1","3:1","4:1","3:2","4:2","4:3"}

	matches2 := []string{"1:1","2:2","3:3","4:4","2:2","3:3","4:4","3:3","4:4","4:4"}

	f.Println("total = ",points(matches1))
	f.Println("total = ",points(matches2))
}