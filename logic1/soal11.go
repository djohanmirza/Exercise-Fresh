package logic1

import "fmt"

func Nomor11() {
	n := 10
	num := 1

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Print(num, " ")
			num *= 2
		} else {
			fmt.Print("buzz ")
			if 1 == i {
				num *= 3
			}
		}
	}
}
