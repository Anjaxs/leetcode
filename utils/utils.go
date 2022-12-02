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

// LoopCompare 循环遍历比较
// 先比较两个数的长度是否相等
// 再循环遍历每一个元素进行比较
func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	//与reflect.DeepEqual的结果保持一致：[]int{} != []int(nil)
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
