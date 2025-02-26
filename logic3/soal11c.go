package logic3

import go_slice_fresh "github.com/djohanmirza/Slice-Fresh"

func Soal11c(n int) (result [][]int) {
	result = go_slice_fresh.CreateSlice(n)
	mid := (n - 1) / 2
	num := 1
	for i := 0; i <= mid; i++ {
		for j := 0; j < mid; j++ {
			if i == j {
				result[i][j] = num
				result[n-1-i][j] = num
			}
		}
		num += 2
	}
	num = 1
	for i := 0; i <= mid; i++ {
		if i%2 == 0 {
			for j := n - 1; j >= mid; j-- {
				if i+j >= n-1 {
					result[i][j] = num
					result[n-1-i][j] = num
					num += 2
				}
			}
		} else {
			for j := n - 1; j >= mid; j-- {
				if i+j >= n-1 {
					result[i][(2*(n-1) - i - j)] = num
					result[n-1-i][(2*(n-1) - i - j)] = num
					num += 2
				}
			}
		}

	}

	return result
}
