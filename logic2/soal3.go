package logic2

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Nomor3(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 2
			} else {
				result[i][j] = start
				start += 2
			}
		}
	}
	go_slice_fresh.Print2D(result)
	return result

}
