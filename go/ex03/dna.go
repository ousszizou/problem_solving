// Go episode N06
// challenge N11 => (Complementary DNA) - 7Kyu

// By ALGORITHM ACADEMY


package main

import (
	"fmt"
	// "strings"
)

// DNAStrand function
func DNAStrand(dna string) string {

	var result string
	for _,c := range dna {
		// method 1
		// if string(c) == "A" {
		// 	result += "T"
		// } else if string(c) == "T" {
		// 	result += "A"
		// } else if string(c) == "C" {
		// 	result += "G"
		// } else {
		// 	result += "C"
		// }

		// method 2
		switch string(c) {
		case "A":
			result += "T"
		case "T":
			result += "A"
		case "C":
			result += "G"
		case "G":
			result += "C"
		}
	}
	return result
}

// Another method
// func DNAStrand(dna string) string {
//   // your code here
//   replacer := strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G")
//   return(replacer.Replace(dna))
// }

func main() {
	fmt.Println(DNAStrand("AAAA"))
	fmt.Println(DNAStrand("ATCG"))
}
