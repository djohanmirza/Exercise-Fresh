package logic2

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal9b(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				result[i][j] = num
			} else if i+j == n-1 {
				result[j][i] = num
			}
		}
		num += 2
	}
	return result
}
