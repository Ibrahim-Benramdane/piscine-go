package main

import (
	"github.com/01-edu/z01"
)
func generateCombitions(prefix []rune, start, n, k int){
	if len(prefix) == k {
		for _, r := range prefix {
			z01.PrintRune(r)
		}
		if prefix[0] != '9'-rune(k-1){
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}
	for i := start; i <=9; i++ {
		newPrefix := append(prefix,rune('0'+i))
		generateCombitions(newPrefix, i+1, n, k)
	}
}
func PrintCombN(n int) {
	
		if n <= 0 || n >= 10 {
			return
		}
		generateCombitions([]rune{}, 0, 9, n)
	z01.PrintRune('\n')
}
func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}
