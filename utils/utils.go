package utils

func MaxInt(arr ...int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > m {
			m = arr[i]
		}
	}
	return m
}

func MinInt(arr ...int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < m {
			m = arr[i]
		}
	}
	return m
}
