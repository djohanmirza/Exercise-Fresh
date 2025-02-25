package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal4c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)

	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j >= n-1 {
				if i%2 == 1 {
					result[i][j] = num
				} else {
					result[i][2*(n-1)-i-j] = num
				}
				num += 2
			}
		}

	}
	return result
}
