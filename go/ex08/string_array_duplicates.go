// Go episode N11
// challenge N16 = > (String array duplicates) - 6Kyu

// By ALGORITHM ACADEMY

package main

import (
	"fmt"
)

func rmDupConsecutiveLetters(word string) string {
	
	var str string

	for i := 0; i < (len(word) - 1); i++ {
		char := string(word[i])
		
		if char == string(word[i+1]) {
			continue
		} else {
			str += char
		}
	}
	str += word[len(word)-1:]

	return str
}

func dup(arr []string) []string {
	var arrWithoutDup []string

	for i := 0; i < len(arr); i++ {
		arrWithoutDup = append(arrWithoutDup, rmDupConsecutiveLetters(arr[i]))
	}

	return arrWithoutDup
}

func main() {
	
	words := []string{"ccooddddddewwwaaaaarrrrsssss", "piccaninny", "hubbubbubboo"}

	fmt.Println(dup(words))
	
}
