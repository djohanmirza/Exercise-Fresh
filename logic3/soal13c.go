package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal13c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	mid := (n - 1) / 2

	for i := 0; i <= mid; i++ {
		num := 1
		for j := 0; j <= mid; j++ {
			if i+j >= mid {
				if i%2 == 0 && j%2 == 0 || i%2 == 1 && j%2 == 1 {
					result[i][j] = num
					result[i][n-1-j] = num
					result[n-1-i][j] = num
					result[n-1-i][n-1-j] = num
				}
			}
			num += 2
		}
	}

	return result
}
