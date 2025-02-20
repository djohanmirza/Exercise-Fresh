package logic1

import "fmt"

func Nomor9() {
	for i := 0; i < 10; i++ {
		fmt.Print(i*3+3, " ")
		if i == 4 {
			for i := 5; i > 0; i-- {
				fmt.Print(i*3, " ")
			}
			break
		}
	}

	fmt.Println(" ")

	for i := 0; i < 10; i++ {
		fmt.Print(i*3+3, " ")
		if i == 5 {
			for i := 5; i > 0; i-- {
				fmt.Print(i*3, " ")
			}
			break
		}
	}
}
