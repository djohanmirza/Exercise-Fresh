package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal7c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	mid := (n - 1) / 2
	num := 1
	for row := 0; row <= mid; row++ {
		for col := 0; col < n; col++ {
			if row+col >= mid && col-row <= mid {
				if row%2 == 1 {
					result[row][col] = num
					result[n-1-row][col] = num
				} else {
					result[row][n-1-col] = num
					result[n-1-row][n-1-col] = num
				}
				num += 2
			}
		}
	}
	return result
}
