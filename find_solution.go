package leetcode

func findSolution(customFunction func(int, int) int, z int) [][]int {
	x, y := 1, 1
	nums := make([][]int, 0)

	// 寻找最大的 x
	for {
		r := customFunction(x, y)
		if r < z {
			x++
			continue
		}
		break
	}

	// 寻找满足条件的数对
	for x > 0 {
		r := customFunction(x, y)
		if r > z {
			x--
		}

		if r < z {
			y++
		}

		if r == z {
			nums = append(nums, []int{x, y})
			x--
			y++
		}
	}

	return nums
}

func add(x, y int) int {
	return x + y
}
