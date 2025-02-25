package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal5c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	//mid := (n - 1) / 2
	num := 1
	for i := 0; i <= n/2; i++ {
		for j := 0; j < n; j++ {
			if (i+j) >= (n-1) && ((i + j) >= n/2) {
				if i%2 != 0 {
					result[i][j] = num
					result[i][n-1-j] = num
					result[n-1-i][j] = num
					result[n-1-i][n-1-j] = num
				} else {
					result[i][(2*(n-1) - i - j)] = num
					result[i][n-1-(2*(n-1)-i-j)] = num
					result[n-1-i][(2*(n-1) - i - j)] = num
					result[n-1-i][n-1-(2*(n-1)-i-j)] = num
				}
				//if i%2 == 1 {
				//	result[i][j] = num
				//	result[i][n-1-j] = num
				//} else {
				//	result[i][j] = num
				//	result[i][n-1-j] = num
				//}
				num += 2
			}

		}
	}
	return result
}
