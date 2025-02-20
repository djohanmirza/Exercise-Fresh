package logic1

import "fmt"

func Nomor12(n int) {
	pattern := []int{1, 3, 5, 7}

	for i := 0; i < n; i++ {
		fmt.Print(pattern[i%4], " ")
	}
	fmt.Println()
}
