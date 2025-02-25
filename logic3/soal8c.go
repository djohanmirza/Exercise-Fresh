package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal8c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	mid := (n - 1) / 2
	num := 1
	for k := 0; k <= mid; k++ {
		for b := 0; b < n; b++ {
			if k+b >= mid && b-k <= mid {
				if k%2 == 1 {
					result[b][k] = num
					result[b][n-1-k] = num
				} else {
					result[n-1-b][k] = num
					result[n-1-b][n-1-k] = num
				}
				//if b%2 == 1 {
				//	result[k][b] = num
				//	result[n-1-k][b] = num
				//} else {
				//	result[k][n-1-b] = num
				//	result[n-1-k][n-1-b] = num
				//}
				num += 2
			}
		}
	}
	return result
}
