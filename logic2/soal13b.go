package logic2

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal13b(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)

	for row := 0; row < n; row++ {
		num := 1
		for col := 0; col < n; col++ {
			if row <= col && row+col <= n-1 {
				result[row][col] = num
				result[n-1-row][col] = num
			}
			num += 2
		}
	}
	return result
}
