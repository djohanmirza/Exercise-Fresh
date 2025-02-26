package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal4c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)

	num := 1
	for i := 0; i < n; i++ {
		geser := 0
		for j := 0; j < n; j++ {
			if i+j == n-1 {
				if i%2 == 1 {
					result[i][j] = num //normal
				} else {
					// harusnya 0,8
					// harusnya yang diisi 6, 7, 8
					result[i][n-1-geser] = num
					geser++
				}
				num += 2
			}
		}

	}
	return result
}
