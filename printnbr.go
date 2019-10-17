package piscine

import "github.com/01-edu/z01"

func PrintRune(n int) {

	t := 1
	if n < 0 {
		t = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		f := (n / 10) * t
		if f != 0 {
			PrintRune(f)
		}
		k := (n % 10 * t) + '0'
		z01.PrintRune(k)
	} else {
		z01.PrintRune('0')
	}

}
