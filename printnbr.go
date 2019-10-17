package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {

	t := 1
	if n < 0 {
		t = -1
		z01.PrintNbr('-')
	}
	if n != 0 {
		f := (n / 10) * t
		if f != 0 {
			PrintNbr(f)
		}
		k := (n % 10 * t) + '0'
		z01.PrintNbr(k)
	} else {
		z01.PrintNbr('0')
	}

}
