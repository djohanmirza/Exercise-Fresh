package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal6c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)

	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if i <= j {
				//result[i][j] = num
				if i%2 == 0 {
					result[i][j] = num
					result[n-1-i][j] = num
				} else {
					result[i][n-1-j] = num
					result[n-1-i][n-1-j] = num
				}
				num += 2
			}

		}
	}
	return result
}
