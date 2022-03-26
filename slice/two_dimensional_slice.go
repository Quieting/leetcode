package slice

// AddRow 在二维切片下标 index 行处添加一行， 如果 len(matrix) == 0 返回空数组
func AddRow(matrix [][]int, index int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	row := len(matrix) + 1
	column := len(matrix[0])
	ns := make([][]int, 0, row)
	for i := 0; i < row; i++ {
		arr := make([]int, column)
		for j := 0; j < len(arr); j++ {
			switch {
			case i == index:
				arr[j] = 0
			case i < index:
				arr[j] = matrix[i][j]
			case i > index:
				arr[j] = matrix[i-1][j]
			}
		}
		ns = append(ns, arr)
	}
	return ns
}

// AddColumn 在二维切片 index 列出添加一列， 如果 len(matrix) == 0 返回空数组
func AddColumn(matrix [][]int, index int) [][]int {
	if len(matrix) == 0 {
		return nil
	}

	row := len(matrix)
	column := len(matrix[0]) + 1
	ns := make([][]int, 0, row)

	for i := 0; i < row; i++ {
		arr := make([]int, column)
		for j := 0; j < len(arr); j++ {
			switch {
			case j == index:
				arr[j] = 0
			case j < index:
				arr[j] = matrix[i][j]
			case j > index:
				arr[j] = matrix[i][j-1]
			}
		}
		ns = append(ns, arr)
	}
	return ns
}

// PrefixSum 求取二维切片前缀和
func PrefixSum(s [][]int) [][]int {
	if len(s) == 0 {
		return nil
	}

	row := len(s)
	column := len(s[0])

	ns := make([][]int, row)
	for r := 0; r < row; r++ {
		ns[r] = make([]int, column)
	}

	// 计算前缀和
	for r := 0; r < row; r++ {
		for c := 0; c < column; c++ {
			// 不考虑边界计算二维数组前缀和公式：
			// ns[r][c] = ns[r][c-1] + ns[r-1][c] + s[r][c] - ns[r-1][c-1]
			ns[r][c] = s[r][c]
			if c-1 >= 0 {
				ns[r][c] += ns[r][c-1]
			}
			if r-1 >= 0 {
				ns[r][c] += ns[r-1][c]
			}
			if r-1 >= 0 && c-1 >= 0 {
				ns[r][c] -= ns[r-1][c-1]
			}
		}
	}

	return ns
}
