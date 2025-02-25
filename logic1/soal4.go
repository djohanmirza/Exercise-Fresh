package logic1

func Nomor4(n int) []int {
	num := 19
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = num
		num -= 2
	}
	return result
}
