package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal3c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)

	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j <= n-1 {
				//result[i][j] = num
				if i%2 == 0 {
					result[i][j] = num
				} else {
					result[i][n-i-j-1] = num
				}
			}
			num += 3
		}

	}
	return result
}
