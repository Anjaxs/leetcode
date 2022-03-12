package dynamic_planning

func tribonacci(n int) int {
	a := []int{0, 1, 1}
	if n <= 2 {
		return a[n]
	}
	for i := 2; i < n; i++ {
		a[0], a[1], a[2] = a[1], a[2], a[0]+a[1]+a[2]
	}
	return a[2]
}
