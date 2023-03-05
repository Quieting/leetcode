package leetcode

// findMedianSortedArrays 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
//
// 提示：
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10^6 <= nums1[i], nums2[i] <= 10^6
//
// 示例 1：
//
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：
//
// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
// PS: 边界判断，奇偶判断，写得太难看了
func findMedianSortedArrays(m, n []int) float64 {
	isTwo := (len(m)+len(n))&1 == 0 // 是否有两个中位数

	k := (len(m) + len(n)) / 2

	var ans []int
	for {
		if k == 0 {
			break
		}
		if len(m) == 0 {
			if isTwo {
				n = n[k-1:]
			} else {
				n = n[k:]
			}

			break
		}
		if len(n) == 0 {
			if isTwo {
				m = m[k-1:]
			} else {
				m = m[k:]
			}
			break
		}
		if k == 1 {
			if !isTwo {
				if m[0] < n[0] {
					m = m[1:]
				} else {
					n = n[1:]
				}
				k--
			}
			break
		}

		k1 := k / 2
		mi := min(len(m), k1)
		ni := min(len(n), k1)

		if m[mi-1] < n[ni-1] {
			m = m[mi:]
			k -= mi
		} else {
			n = n[ni:]
			k -= ni
		}
	}

	ans = twoSmaller(m, n)
	if isTwo {
		return float64(ans[0]+ans[1]) / 2
	}
	return float64(ans[0])
}

// twoSmaller 返回 m, n 中最小的两个元素
func twoSmaller(m, n []int) []int {
	if len(m) == 0 {
		if len(n) > 1 {
			return n[:2]
		} else {
			return []int{n[0], n[0]}
		}

	}
	if len(n) == 0 {
		if len(m) > 1 {
			return m[:2]
		} else {
			return []int{m[0], m[0]}
		}
	}

	var ans []int

	if m[0] < n[0] {
		ans = append(ans, m[0])
		m = m[1:]
	} else {
		ans = append(ans, n[0])
		n = n[1:]
	}

	if len(m) == 0 {
		ans = append(ans, n[0])
	} else if len(n) == 0 {
		ans = append(ans, m[0])
	} else if m[0] < n[0] {
		ans = append(ans, m[0])
	} else {
		ans = append(ans, n[0])
	}
	return ans
}
