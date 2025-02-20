package logic2

import (
	go_slice_fresh "github.com/djohanmirza/Slice-Fresh"
)

func Nomor1(n int) (result [][]int) {
	//matrix := make([][]int, n)
	//for i := range matrix {
	//	matrix[i] = make([]int, n)
	//}
	result = go_slice_fresh.CreateSlice(n)
	for i := 0; i < n; i++ {
		num := 1
		for j := 0; j < n; j++ {
			result[i][j] = num
			num += 2
		}
	}
	go_slice_fresh.Print2D(result)
	return result
}
