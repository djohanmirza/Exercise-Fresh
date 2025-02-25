package logic1

func Nomor1(n int) []int {
	result := make([]int, n)
	num := 1
	for i := 0; i < n; i++ {
		result[i] = num
		num += 2
	}
	return result
}
