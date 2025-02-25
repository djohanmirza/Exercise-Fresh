package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal2c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)

	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			//if i <= j {
			//result[i][j] = num
			if i%2 == 0 {
				result[i][j+i] = num
			} else {
				result[i][n-j-1] = num
			}
			num += 2
		}

	}
	//}
	return result
}
