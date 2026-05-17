package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var counts [10]int
	for n > 0 {
		digit := n % 10
		counts[digit]++
		n /= 10
	}
	for i := 0; i <= 9; i++ {
		for counts[i] > 0 {
			z01.PrintRune(rune(i + '0'))
			counts[i]--
		}
	}
}
