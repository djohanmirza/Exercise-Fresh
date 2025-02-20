package logic2

import "github.com/djohanmirza/Slice-Fresh"

func Nomor6a(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 3
			} else {
				result[i][n-j-1] = start
				start += 2
			}
		}
	}
	go_slice_fresh.Print2D(result)
	return result

}
