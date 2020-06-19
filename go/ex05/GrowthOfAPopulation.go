// Go episode N07
// challenge N13 => (Growth of a Population) - 7Kyu

// By ALGORITHM ACADEMY

package main

import "fmt"

// NbYear function
func NbYear(p0 int, percent float64, aug int, p int) int {
	
	inhabitants := float64(p0) + (float64(p0) * (percent/100)) + float64(aug)

	var counter int
	for {
		if inhabitants < float64(p) {
			counter++
			inhabitants = float64(inhabitants) + (float64(inhabitants) * (percent/100)) + float64(aug)
		} else {
			counter++
			break
		}
	}
	return counter
}

// Another Solution
// func NbYear(p0 int, percent float64, aug int, p int) (n int) {
//   if p0 >= p {
//     return n
//   }
//   newP := p0 + aug + int(float64(p0)*percent/100)
//   n++
//   return n + NbYear(newP, percent, aug, p)
// }

func main()  {
	fmt.Println(NbYear(1500, 5, 100, 5000)) // 15
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000)) // 10
	fmt.Println(NbYear(1500000, 0.25, 1000, 2000000)) // 94
	fmt.Println(NbYear(1500000, 0.25, -1000, 2000000)) // 151
}
