package logic2

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal10b(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)

	for i := 0; i < n; i++ {
		num := 1
		for j := 0; j < n; j++ {
			if i >= j {
				result[i][j] = num
				num += 2
			}
		}
	}
	return result
}
