package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal3c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)

	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j {
				//result[i][j] = num
				if i == 0 {
					result[i][j-i] = num
					num += 3
				} else if i%2 == 0 {
					result[i][j-i] = num
					num += 2
					if j == n-1 {
						num += 1
					}
				} else {
					result[i][n-j-1] = num
					num += 3
					if j == n-1 {
						num -= 1
					}
				}
			}
		}

	}
	result[0][0] = 2
	return result
}
