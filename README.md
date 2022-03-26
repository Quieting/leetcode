# leetcode
leetcode题目详解</br>
每个文件对应leetcode一道题，文件命令格式为 p_number_*.go, number 为题序号，以题目标题结尾</br>
每个文件应包含如下内容</br>
1.题目描述</br>
2.示例（输入和输出）</br>
3.题解描述</br>
4.题解</br>
5.题标签</br>
例如：  
```
// optimalDivision  553. 最优除法
// 给定一组正整数，相邻的整数之间将会进行浮点除法操作。例如， [2,3,4] -> 2 / 3 / 4 。
// 你可以在任意位置添加任意数目的括号，来改变算数的优先级。你需要找出怎么添加括号，才能得到最大的结果，并且返回相应的字符串格式的表达式。
// 你的表达式不应该含有冗余的括号。

// 说明:
// 输入数组的长度在 [1, 10] 之间。
// 数组中每个元素的大小都在 [2, 1000] 之间。
// 每个测试用例只有一个最优除法解。

// 中等 

// 示例：
// 输入: [1000,100,10,2]
// 输出: "1000/(100/10/2)"
// 解释:
// 1000/(100/10/2) = 1000/((100/10)/2) = 200
// 但是，以下加粗的括号 "1000/((100/10)/2)" 是冗余的，
// 因为他们并不影响操作的优先级，所以你需要返回 "1000/(100/10/2)"。

// 动态规划
func optimalDivision(nums []int) string {

}

```

### 目录
* [6.字符串Z型变换](p_6_convert_z_string.go)
* [219.存在重复元素](p_219_contains_nearby_duplicate.go)
* [363.矩形区域不超过 K 的最大数值和](p_363_max_sum_submatrix.go)
* [413.等差数列划分 - 子数组](p_413_number_of_arithmetic_slices.go)
* [440.字典序的第K小数字](p_440_find_kth_number.go)
* [446.等差数列划分 II - 子序列](p_446_number_of_arithmetic_slices.go)
* [537.复数乘法](p_537_complex_number_multiply.go)
* [553.最优除法](p_553_optimal_division.go)
* [564.寻找最近的回文数](p_564_nearest_palindromic.go)
* [589.N 叉树的前序遍历](p_589_preorder.go)
* [661.图片平滑器](p_661_image_smoother.go)
* [682.棒球比赛](p_682_cal_points.go)
* [917.仅仅反转字符串](p_917_reverse_only_letters.go)
* [2038.如果相邻两个颜色均相同则删除当前颜色](p_2038_winner_of_game.go)
* [2044.统计按位或能得到最大值的子集数目](p_2044_count_max_or_subsets.go)
* [2049.统计最高分的节点数目](p_2049_count_highest_score_nodes.go)
