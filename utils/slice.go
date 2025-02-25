package utils

import "fmt"

func Print(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], "\t")
	}
}

func Print2D(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if slice[i][j] == 0 {
				fmt.Print(".\t")
			} else {
				fmt.Print(slice[i][j], "\t")
			}
		}
		fmt.Println()
	}
}
