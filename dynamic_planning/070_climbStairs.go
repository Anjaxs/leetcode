package dynamic_planning

func climbStairs(n int) int {
	a := []int{0, 1, 2}
	if n < 3 {
		return a[n]
	}
	for i := 3; i <= n; i++ {
		a[0], a[1], a[2] = a[1], a[2], a[1]+a[2]
	}
	return a[2]
}
