package logic2

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Nomor2(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	for i := 0; i < n; i++ {
		num := 2
		for j := 0; j < n; j++ {
			result[i][j] = num
			num += 2
		}
	}
	go_slice_fresh.Print2D(result)
	return result
}
