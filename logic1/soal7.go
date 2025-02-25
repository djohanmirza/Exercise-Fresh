package logic1

func Nomor71(n int) []int {
	mid := n / 2
	num := 1
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = num
		if i < mid {
			num += 2
		} else if i > mid {
			num -= 2
		}

	}
	return result
}

func Nomor72(n int) []int {
	mid := n / 2
	num := 1
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = num
		if i < mid {
			num += 2
		} else if i >= mid {
			num -= 2
		}

	}
	return result
}
