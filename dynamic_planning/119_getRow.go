package leetcode

func getRow(rowIndex int) []int {
	prevRow := make([]int, rowIndex+1)
	row := make([]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = prevRow[j-1] + prevRow[j]
			}
		}
		copy(prevRow, row)
	}
	return row
}
