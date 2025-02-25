package logic1

func Nomor91(n int) []int {
	mid := n / 2
	num := 3
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = num
		if i < mid {
			num += 3
		} else if i > mid {
			num -= 3
		}
	}
	return result
}

func Nomor92(n int) []int {
	mid := n / 2
	num := 3
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = num
		if i < mid {
			num += 3
		} else if i >= mid {
			num -= 3
		}
	}
	return result
}
