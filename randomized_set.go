package leetcode

import "math/rand"

// 实现RandomizedSet 类：
//
// RandomizedSet() 初始化 RandomizedSet 对象
// bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
// bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
// int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
// 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
//
//
//
// 示例：
//
// 输入
// ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
// [[], [1], [2], [2], [], [1], [2], []]
// 输出
// [null, true, false, true, 2, true, false, 2]
//
// 解释
// RandomizedSet randomizedSet = new RandomizedSet();
// randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
// randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
// randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
// randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
// randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
// randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
// randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
//
//
// 提示：
//
// -2^31 <= val <= 2^31 - 1
// 最多调用 insert、remove 和 getRandom 函数 2 * 10^5 次
// 在调用 getRandom 方法时，数据结构中 至少存在一个 元素。
//
/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

type RandomizedSet struct {
	nums      []int
	numsIndex map[int]int // key: nums的元素值 val：nums元素值对应的下标
}

func Constructor() RandomizedSet {
	ins := RandomizedSet{
		nums:      []int{},
		numsIndex: make(map[int]int),
	}
	return ins
}

func (set *RandomizedSet) Insert(val int) bool {
	if _, ok := set.numsIndex[val]; ok {
		return false
	}
	set.nums = append(set.nums, val)
	set.numsIndex[val] = len(set.nums) - 1
	return true
}

func (set *RandomizedSet) Remove(val int) bool {
	index, ok := set.numsIndex[val]
	if !ok {
		return false
	}
	nLen := len(set.nums)
	set.nums[index] = set.nums[nLen-1]
	set.numsIndex[set.nums[index]] = index
	set.nums = set.nums[:nLen-1]
	delete(set.numsIndex, val)
	return true
}

func (set *RandomizedSet) GetRandom() int {
	return set.nums[rand.Intn(len(set.nums))]
}
