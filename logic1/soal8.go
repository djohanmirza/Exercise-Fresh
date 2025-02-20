package logic1

import "fmt"

func Nomor8() {
	for i := 0; i < 10; i++ {
		fmt.Print(2*i+2, " ")
		if i == 4 {
			for i := 5; i > 0; i-- {
				fmt.Print(2*i, " ")
			}
			break
		}
	}

	fmt.Println(" ")

	for i := 0; i < 10; i++ {
		fmt.Print(2*i+2, " ")
		if i == 5 {
			for i := 5; i > 0; i-- {
				fmt.Print(2*i, " ")
			}
			break
		}
	}
}
